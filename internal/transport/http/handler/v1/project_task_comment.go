package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/timeutil"
)

type ProjectTaskComment struct {
	ProjectService *service.ProjectService
}

func (p *ProjectTaskComment) Create(ctx *core.Context) error {
	params := &api_v1.ProjectCommentCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	commentId, err := p.ProjectService.CreateComment(ctx.Ctx(), &service.ProjectCommentOpt{
		TaskId:    params.TaskId,
		Comment:   params.Comment,
		CreatedBy: ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать, попробуйте позже: " + err.Error())
	}

	return ctx.Success(&api_v1.ProjectCreateResponse{Id: commentId})
}

func (p *ProjectTaskComment) Comments(ctx *core.Context) error {
	params := &api_v1.ProjectCommentRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := p.ProjectService.Comments(ctx.Ctx(), params.TaskId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*api_v1.ProjectCommentResponse_Item, 0)
	for _, item := range data {
		user, err := p.ProjectService.UserRepo.FindById(ctx.Ctx(), item.CreatedBy)
		if err != nil {
			return ctx.ErrorBusiness(err.Error())
		}

		items = append(items, &api_v1.ProjectCommentResponse_Item{
			Id:      item.Id,
			TaskId:  item.TaskId,
			Comment: item.Comment,
			User: &api_v1.ProjectCommentResponse_User{
				Id:       int64(user.Id),
				Avatar:   user.Avatar,
				Username: user.Username,
				Name:     user.Name,
				Surname:  user.Surname,
			},
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	return ctx.Success(api_v1.ProjectCommentResponse{Items: items})
}
