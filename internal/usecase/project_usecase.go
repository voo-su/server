package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/locale"
)

type ProjectUseCase struct {
	Locale                    locale.ILocale
	Source                    *infrastructure.Source
	ProjectRepo               *postgresRepo.ProjectRepository
	ProjectMemberRepo         *postgresRepo.ProjectMemberRepository
	ProjectTaskTypeRepo       *postgresRepo.ProjectTaskTypeRepository
	ProjectTaskRepo           *postgresRepo.ProjectTaskRepository
	ProjectTaskCommentRepo    *postgresRepo.ProjectTaskCommentRepository
	UserRepo                  *postgresRepo.UserRepository
	ProjectContactCache       *redisRepo.ProjectContactCacheRepository
	ProjectTaskCoexecutorRepo *postgresRepo.ProjectTaskCoexecutorRepository
	ProjectTaskWatcherRepo    *postgresRepo.ProjectTaskWatcherRepository
	RedisLockCacheRepo        *redisRepo.RedisLockCacheRepository
}

func NewProjectUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	projectRepo *postgresRepo.ProjectRepository,
	projectMemberRepo *postgresRepo.ProjectMemberRepository,
	projectTaskTypeRepo *postgresRepo.ProjectTaskTypeRepository,
	projectTaskRepo *postgresRepo.ProjectTaskRepository,
	projectTaskCommentRepo *postgresRepo.ProjectTaskCommentRepository,
	userRepo *postgresRepo.UserRepository,
	projectContactCache *redisRepo.ProjectContactCacheRepository,
	projectTaskCoexecutorRepo *postgresRepo.ProjectTaskCoexecutorRepository,
	projectTaskWatcherRepo *postgresRepo.ProjectTaskWatcherRepository,
	redisLockCacheRepo *redisRepo.RedisLockCacheRepository,
) *ProjectUseCase {
	return &ProjectUseCase{
		Locale:                    locale,
		Source:                    source,
		ProjectRepo:               projectRepo,
		ProjectMemberRepo:         projectMemberRepo,
		ProjectTaskTypeRepo:       projectTaskTypeRepo,
		ProjectTaskRepo:           projectTaskRepo,
		ProjectTaskCommentRepo:    projectTaskCommentRepo,
		UserRepo:                  userRepo,
		ProjectContactCache:       projectContactCache,
		ProjectTaskCoexecutorRepo: projectTaskCoexecutorRepo,
		ProjectTaskWatcherRepo:    projectTaskWatcherRepo,
		RedisLockCacheRepo:        redisLockCacheRepo,
	}
}

type ProjectOpt struct {
	UserId int
	Title  string
}

func (p *ProjectUseCase) CreateProject(ctx context.Context, opt *ProjectOpt) (uuid.UUID, error) {
	var (
		err       error
		members   []*postgresModel.ProjectMember
		taskTypes []*postgresModel.ProjectTaskType
	)

	project := &postgresModel.Project{
		Name:      opt.Title,
		CreatedBy: opt.UserId,
		CreatedAt: time.Now(),
	}

	types := [4]string{"Новые", "Выполняются", "Тест", "Сделаны"}

	err = p.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(project).Error; err != nil {
			return err
		}

		members = append(members, &postgresModel.ProjectMember{
			ProjectId: project.Id,
			UserId:    opt.UserId,
			CreatedBy: opt.UserId,
		})

		if err = tx.Create(members).Error; err != nil {
			return err
		}

		for _, title := range types {
			taskTypes = append(taskTypes, &postgresModel.ProjectTaskType{
				ProjectId: project.Id,
				Title:     title,
				CreatedBy: opt.UserId,
			})
		}
		if err = tx.Create(taskTypes).Error; err != nil {
			return err
		}

		return nil
	})

	return project.Id, nil
}

func (p *ProjectUseCase) Projects(userId int) ([]*entity.ProjectItem, error) {
	tx := p.Source.Postgres().
		Table("project_members").
		Select("p.id AS id, p.name AS name").
		Joins("LEFT JOIN projects p ON p.id = project_members.project_id").
		Where("project_members.user_id = ?", userId).
		Order("project_members.created_at desc")

	items := make([]*entity.ProjectItem, 0)
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	length := len(items)
	if length == 0 {
		return items, nil
	}

	ids := make([]uuid.UUID, 0, length)
	for i := range items {
		ids = append(ids, items[i].Id)
	}

	return items, nil
}

func (p *ProjectUseCase) Detail(ctx context.Context, uid int, projectId uuid.UUID) (*entity.ProjectDetailItem, error) {
	exist, err := p.ProjectMemberRepo.QueryExist(ctx, "project_id = ? AND user_id = ?", projectId, uid)
	if err != nil {
		return nil, errors.New(p.Locale.Localize("not_project_member"))
	}
	if !exist {
		return nil, errors.New(p.Locale.Localize("not_project_member"))
	}

	project, err := p.ProjectRepo.FindByWhere(ctx, "id = ?", projectId)
	if err != nil {
		return nil, err
	}

	return &entity.ProjectDetailItem{
		Id:   project.Id,
		Name: project.Name,
	}, nil
}

func (p *ProjectUseCase) GetMembers(ctx context.Context, projectId uuid.UUID) []*entity.ProjectMemberItem {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
	}
	tx := p.Source.Postgres().WithContext(ctx).Table("project_members").
		Joins("LEFT JOIN users ON users.id = project_members.user_id").
		Where("project_members.project_id = ?", projectId)
	//.Order("project_members.leader desc")

	var items []*entity.ProjectMemberItem
	tx.Unscoped().Select(fields).Scan(&items)

	return items
}

func (p *ProjectUseCase) IsMember(ctx context.Context, gid uuid.UUID, uid int, cache bool) bool {
	if cache && p.ProjectContactCache.IsGroupRelation(ctx, uid, gid) == nil {
		return true
	}

	exist, err := p.ProjectMemberRepo.QueryExist(ctx, "project_id = ? AND user_id = ?", gid, uid)
	if err != nil {
		return false
	}
	if exist {
		p.ProjectContactCache.SetGroupRelation(ctx, uid, gid)
	}

	return exist
}

type ProjectInviteOpt struct {
	ProjectId uuid.UUID
	UserId    int
	MemberIds []int
}

func (p *ProjectUseCase) Invite(ctx context.Context, opt *ProjectInviteOpt) error {
	var (
		err        error
		addMembers []*postgresModel.ProjectMember
		db         = p.Source.Postgres().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectMemberRepo.GetMemberIds(ctx, opt.ProjectId) {
		m[value] = struct{}{}
	}

	for _, value := range opt.MemberIds {
		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &postgresModel.ProjectMember{
				ProjectId: opt.ProjectId,
				UserId:    value,
				CreatedBy: opt.UserId,
			})
		}
	}
	if len(addMembers) == 0 {
		return errors.New(p.Locale.Localize("all_invited_contacts_are_project_members"))
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&postgresModel.ProjectMember{}, "project_id = ? AND user_id in ?", opt.ProjectId, opt.MemberIds)
		if err = tx.Create(&addMembers).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
