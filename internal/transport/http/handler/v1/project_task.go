package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/timeutil"
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
		TypeId:      int(params.TypeId),
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

func (p *ProjectTask) TaskDetail(ctx *core.Context) error {
	params := &api_v1.ProjectTaskDetailRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	task, err := p.ProjectService.TaskDetail(ctx.Ctx(), params.TaskId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(api_v1.ProjectTaskDetailResponse{
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   timeutil.FormatDatetime(task.CreatedAt),
	})
}

func (p *ProjectTask) TaskMove(ctx *core.Context) error {
	params := &api_v1.ProjectTaskMoveRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.ProjectService.TaskMove(
		ctx.Ctx(),
		params.ProjectId,
		params.TaskId,
		params.FromId,
		params.ToId,
	); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(api_v1.ProjectTaskMoveResponse{})
}

func (p *ProjectTask) TaskTypeName(ctx *core.Context) error {
	params := &api_v1.ProjectTaskTypeNameRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.ProjectService.TaskTypeName(
		ctx.Ctx(),
		params.TaskId,
		params.Name,
	); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(api_v1.ProjectTaskTypeNameResponse{})
}

func (p *ProjectTask) TaskCoexecutors(ctx *core.Context) error {
	params := &api_v1.ProjectCoexecutorsRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	list := p.ProjectService.GetCoexecutors(ctx.Ctx(), params.TaskId)

	items := make([]*api_v1.ProjectCoexecutorsResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.ProjectCoexecutorsResponse_Item{
			Id:       item.Id,
			Username: item.Username,
		})
	}

	return ctx.Success(api_v1.ProjectCoexecutorsResponse{
		Items: items,
	})
}

func (p *ProjectTask) TaskWatchers(ctx *core.Context) error {
	params := &api_v1.ProjectWatchersRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	list := p.ProjectService.GetWatchers(ctx.Ctx(), params.TaskId)

	items := make([]*api_v1.ProjectWatchersResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.ProjectWatchersResponse_Item{
			Id:       item.Id,
			Username: item.Username,
		})
	}

	return ctx.Success(api_v1.ProjectWatchersResponse{
		Items: items,
	})
}
