package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
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
		TaskId:      params.TaskId,
		CommentText: params.Comment,
		CreatedBy:   ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать, попробуйте позже" + err.Error())
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
		items = append(items, &api_v1.ProjectCommentResponse_Item{
			Id:      int64(item.Id),
			Comment: item.Text,
		})
	}

	return ctx.Success(api_v1.ProjectCommentResponse{Items: items})
}
