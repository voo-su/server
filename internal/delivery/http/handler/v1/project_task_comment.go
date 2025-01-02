// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package v1

import (
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type ProjectTaskComment struct {
	Locale         locale.ILocale
	ProjectUseCase *usecase.ProjectUseCase
}

func (p *ProjectTaskComment) Create(ctx *core.Context) error {
	params := &v1Pb.ProjectCommentCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	commentId, err := p.ProjectUseCase.CreateComment(ctx.Ctx(), &usecase.ProjectCommentOpt{
		TaskId:    params.TaskId,
		Comment:   params.Comment,
		CreatedBy: ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness(p.Locale.Localize("creation_failed_try_later") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectCreateResponse{Id: commentId})
}

func (p *ProjectTaskComment) Comments(ctx *core.Context) error {
	params := &v1Pb.ProjectCommentRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := p.ProjectUseCase.Comments(ctx.Ctx(), params.TaskId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*v1Pb.ProjectCommentResponse_Item, 0)
	for _, item := range data {
		user, err := p.ProjectUseCase.UserRepo.FindById(ctx.Ctx(), item.CreatedBy)
		if err != nil {
			return ctx.ErrorBusiness(err.Error())
		}

		items = append(items, &v1Pb.ProjectCommentResponse_Item{
			Id:      item.Id,
			TaskId:  item.TaskId,
			Comment: item.Comment,
			User: &v1Pb.ProjectCommentResponse_User{
				Id:       int64(user.Id),
				Avatar:   user.Avatar,
				Username: user.Username,
				Name:     user.Name,
				Surname:  user.Surname,
			},
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	return ctx.Success(v1Pb.ProjectCommentResponse{Items: items})
}
