package v1

import (
	"github.com/google/uuid"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type ProjectTaskComment struct {
	Locale         locale.ILocale
	ProjectUseCase *usecase.ProjectUseCase
}

func (p *ProjectTaskComment) Create(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectCommentCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	commentId, err := p.ProjectUseCase.CreateComment(ctx.Ctx(), &usecase.ProjectCommentOpt{
		TaskId:    taskId,
		Comment:   params.Comment,
		CreatedBy: ctx.UserId(),
	})
	if err != nil {
		return ctx.Error(p.Locale.Localize("creation_failed_try_later") + ": " + err.Error())
	}

	return ctx.Success(&v1Pb.ProjectCommentCreateResponse{Id: commentId})
}

func (p *ProjectTaskComment) Comments(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectCommentRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	data, err := p.ProjectUseCase.Comments(ctx.Ctx(), taskId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	items := make([]*v1Pb.ProjectCommentResponse_Item, 0)
	for _, item := range data {
		user, err := p.ProjectUseCase.UserRepo.FindById(ctx.Ctx(), item.CreatedBy)
		if err != nil {
			return ctx.Error(err.Error())
		}

		items = append(items, &v1Pb.ProjectCommentResponse_Item{
			Id:      item.Id,
			TaskId:  item.TaskId.String(),
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
