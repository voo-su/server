// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package v1

import (
	"fmt"
	v1Pb "voo.su/api/http/pb/v1"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
)

type Project struct {
	Locale             locale.ILocale
	ProjectUseCase     *usecase.ProjectUseCase
	RedisLockCacheRepo *redisRepo.RedisLockCacheRepository
}

func (p *Project) Create(ctx *core.Context) error {
	params := &v1Pb.ProjectCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	projectId, err := p.ProjectUseCase.CreateProject(ctx.Ctx(), &usecase.ProjectOpt{
		UserId: ctx.UserId(),
		Title:  params.Title,
	})
	if err != nil {
		return ctx.ErrorBusiness(p.Locale.Localize("creation_failed_try_later") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectCreateResponse{Id: projectId})
}

func (p *Project) Projects(ctx *core.Context) error {
	data, err := p.ProjectUseCase.Projects(ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
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

func (p *Project) Members(ctx *core.Context) error {
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

func (p *Project) Invite(ctx *core.Context) error {
	params := &v1Pb.ProjectInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	key := fmt.Sprintf("project-join:%d", params.ProjectId)
	if !p.RedisLockCacheRepo.Lock(ctx.Ctx(), key, 20) {
		return ctx.ErrorBusiness(p.Locale.Localize("network_error"))
	}

	defer p.RedisLockCacheRepo.UnLock(ctx.Ctx(), key)

	project, err := p.ProjectUseCase.ProjectRepo.FindById(ctx.Ctx(), int(params.ProjectId))
	if err != nil {
		return ctx.ErrorBusiness(p.Locale.Localize("network_error"))
	}

	if project == nil {
		return ctx.ErrorBusiness(p.Locale.Localize("project_dissolved"))
	}

	uids := sliceutil.Unique(sliceutil.ParseIds(params.Ids))
	if len(uids) == 0 {
		return ctx.ErrorBusiness(p.Locale.Localize("invited_list_cannot_be_empty"))
	}

	uid := ctx.UserId()
	if !p.ProjectUseCase.IsMember(ctx.Ctx(), int(params.ProjectId), uid, true) {
		return ctx.ErrorBusiness(p.Locale.Localize("not_project_member_cannot_invite"))
	}

	if err := p.ProjectUseCase.Invite(ctx.Ctx(), &usecase.ProjectInviteOpt{
		ProjectId: int(params.ProjectId),
		UserId:    uid,
		MemberIds: uids,
	}); err != nil {
		return ctx.ErrorBusiness(p.Locale.Localize("failed_to_send_invitations") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectInviteResponse{})
}
