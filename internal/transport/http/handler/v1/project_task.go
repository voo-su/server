package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type ProjectTask struct {
	ProjectService *service.ProjectService
}

func (p *ProjectTask) Create(ctx *core.Context) error {
	params := &api_v1.ProjectTaskCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := p.ProjectService.CreateTask(ctx.Ctx(), &service.ProjectTaskOpt{
		ProjectId:   params.ProjectId,
		TypeId:      1,
		Title:       params.Title,
		Description: params.Description,
		CreatedBy:   ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать, попробуйте позже" + err.Error())
	}

	return ctx.Success(&api_v1.ProjectTaskCreateResponse{
		Id: taskId,
	})
}

func (p *ProjectTask) Tasks(ctx *core.Context) error {
	params := &api_v1.ProjectTaskRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := p.ProjectService.TypeTasks(ctx.Ctx(), params.ProjectId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	categories := make([]*api_v1.ProjectTaskResponse_Categories, 0)
	for _, item := range data {

		tasks, _err := p.ProjectService.Tasks(ctx.Ctx(), params.ProjectId, item.Id)
		if _err != nil {
			return ctx.ErrorBusiness(_err.Error())
		}

		taskItems := make([]*api_v1.ProjectTaskResponse_Tasks, 0)
		for _, taskItem := range tasks {
			taskItems = append(taskItems, &api_v1.ProjectTaskResponse_Tasks{
				Id:    taskItem.Id,
				Title: taskItem.Title,
			})
		}

		categories = append(categories, &api_v1.ProjectTaskResponse_Categories{
			Id:    item.Id,
			Title: item.Title,
			Tasks: taskItems,
		})
	}

	return ctx.Success(api_v1.ProjectTaskResponse{
		Categories: categories,
	})
}

func (p *ProjectTask) TaskMove(ctx *core.Context) error {
	params := &api_v1.ProjectTaskMoveRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	err := p.ProjectService.TaskMove(ctx.Ctx(), params.ProjectId, params.TaskId, params.FromId, params.ToId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(api_v1.ProjectTaskMoveResponse{})
}
