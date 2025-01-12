package v1

import (
	"fmt"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
)

type Project struct {
	Locale         locale.ILocale
	ProjectUseCase *usecase.ProjectUseCase
	ContactUseCase *usecase.ContactUseCase
}

func (p *Project) Create(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	projectId, err := p.ProjectUseCase.CreateProject(ctx.Ctx(), &usecase.ProjectOpt{
		UserId: uid,
		Title:  params.Title,
	})
	if err != nil {
		return ctx.Error(p.Locale.Localize("creation_failed_try_later") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectCreateResponse{Id: projectId})
}

func (p *Project) Projects(ctx *ginutil.Context) error {
	data, err := p.ProjectUseCase.Projects(ctx.UserId())
	if err != nil {
		return ctx.Error(err.Error())
	}

	items := make([]*v1Pb.ProjectListResponse_Item, 0)
	for _, item := range data {
		items = append(items, &v1Pb.ProjectListResponse_Item{
			Id:    int64(item.Id),
			Title: item.Name,
		})
	}

	return ctx.Success(v1Pb.ProjectListResponse{Items: items})
}

func (p *Project) Members(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectMembersRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	list := p.ProjectUseCase.GetMembers(ctx.Ctx(), params.ProjectId)

	items := make([]*v1Pb.ProjectMembersResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.ProjectMembersResponse_Item{
			Id:       item.Id,
			Username: item.Username,
		})
	}

	return ctx.Success(v1Pb.ProjectMembersResponse{Items: items})
}

func (p *Project) Invite(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	key := fmt.Sprintf("project-join:%d", params.ProjectId)
	if !p.ProjectUseCase.RedisLockCacheRepo.Lock(ctx.Ctx(), key, 20) {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	defer p.ProjectUseCase.RedisLockCacheRepo.UnLock(ctx.Ctx(), key)

	project, err := p.ProjectUseCase.ProjectRepo.FindById(ctx.Ctx(), int(params.ProjectId))
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	if project == nil {
		return ctx.Error(p.Locale.Localize("project_dissolved"))
	}

	uIds := sliceutil.Unique(sliceutil.ParseIds(params.Ids))
	if len(uIds) == 0 {
		return ctx.Error(p.Locale.Localize("invited_list_cannot_be_empty"))
	}

	uid := ctx.UserId()
	if !p.ProjectUseCase.IsMember(ctx.Ctx(), int(params.ProjectId), uid, true) {
		return ctx.Error(p.Locale.Localize("not_project_member_cannot_invite"))
	}

	if err := p.ProjectUseCase.Invite(ctx.Ctx(), &usecase.ProjectInviteOpt{
		ProjectId: int(params.ProjectId),
		UserId:    uid,
		MemberIds: uIds,
	}); err != nil {
		return ctx.Error(p.Locale.Localize("failed_to_send_invitations") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectInviteResponse{})
}
