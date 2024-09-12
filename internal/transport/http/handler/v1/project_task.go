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
		TaskType:    1,
		Title:       params.Title,
		Description: params.Description,
		CreatedBy:   ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать, попробуйте позже" + err.Error())
	}

	return ctx.Success(&api_v1.ProjectTaskCreateResponse{Id: taskId})
}

func (p *ProjectTask) Types(ctx *core.Context) error {
	params := &api_v1.ProjectTaskTypeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := p.ProjectService.TypeTasks(ctx.Ctx(), params.ProjectId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}
	items := make([]*api_v1.ProjectTaskTypeResponse_Item, 0)
	for _, item := range data {
		items = append(items, &api_v1.ProjectTaskTypeResponse_Item{
			Id:    int64(item.Id),
			Title: item.Title,
		})
	}

	return ctx.Success(api_v1.ProjectTaskTypeResponse{
		Items: items,
	})
}

func (p *ProjectTask) Tasks(ctx *core.Context) error {
	params := &api_v1.ProjectTaskRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := p.ProjectService.Tasks(ctx.Ctx(), params.ProjectId, int(params.TypeId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*api_v1.ProjectTaskResponse_Item, 0)
	for _, item := range data {
		items = append(items, &api_v1.ProjectTaskResponse_Item{
			Id:    int64(item.Id),
			Title: item.Title,
		})
	}

	return ctx.Success(api_v1.ProjectTaskResponse{
		Items: items,
	})
}
