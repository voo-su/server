package v1

import (
	"fmt"
	"voo.su/api/pb/v1"
	"voo.su/internal/repository/cache"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/sliceutil"
)

type Project struct {
	ProjectService *service.ProjectService
	RedisLock      *cache.RedisLock
}

func (p *Project) Create(ctx *core.Context) error {
	params := &api_v1.ProjectCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	projectId, err := p.ProjectService.CreateProject(ctx.Ctx(), &service.ProjectOpt{
		UserId: ctx.UserId(),
		Title:  params.Title,
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать, попробуйте позже: " + err.Error())
	}

	return ctx.Success(&api_v1.ProjectCreateResponse{Id: projectId})
}

func (p *Project) Projects(ctx *core.Context) error {
	data, err := p.ProjectService.Projects(ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*api_v1.ProjectListResponse_Item, 0)
	for _, item := range data {
		items = append(items, &api_v1.ProjectListResponse_Item{
			Id:    int64(item.Id),
			Title: item.Name,
		})
	}

	return ctx.Success(api_v1.ProjectListResponse{Items: items})
}

func (p *Project) Members(ctx *core.Context) error {
	params := &api_v1.ProjectMembersRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	list := p.ProjectService.GetMembers(ctx.Ctx(), params.ProjectId)

	items := make([]*api_v1.ProjectMembersResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.ProjectMembersResponse_Item{
			Id:       item.Id,
			Username: item.Username,
		})
	}

	return ctx.Success(api_v1.ProjectMembersResponse{Items: items})
}

func (p *Project) Invite(ctx *core.Context) error {
	params := &api_v1.ProjectInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	key := fmt.Sprintf("project-join:%d", params.ProjectId)
	if !p.RedisLock.Lock(ctx.Ctx(), key, 20) {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	defer p.RedisLock.UnLock(ctx.Ctx(), key)
	project, err := p.ProjectService.ProjectRepo.FindById(ctx.Ctx(), int(params.ProjectId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	if project == nil {
		return ctx.ErrorBusiness("Проект была расформирована")
	}

	uid := ctx.UserId()
	uids := sliceutil.Unique(sliceutil.ParseIds(params.Ids))
	if len(uids) == 0 {
		return ctx.ErrorBusiness("Список приглашенных не может быть пустым")
	}

	if !p.ProjectService.IsMember(ctx.Ctx(), int(params.ProjectId), uid, true) {
		return ctx.ErrorBusiness("Вы не являетесь участником проекта и не имеете права приглашать")
	}

	if err := p.ProjectService.Invite(ctx.Ctx(), &service.ProjectInviteOpt{
		ProjectId: int(params.ProjectId),
		UserId:    uid,
		MemberIds: uids,
	}); err != nil {
		return ctx.ErrorBusiness("Не удалось пригласить: " + err.Error())
	}

	return ctx.Success(&api_v1.ProjectInviteResponse{})
}
