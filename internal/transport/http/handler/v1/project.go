package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type Project struct {
	ProjectService *service.ProjectService
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
	data, err := p.ProjectService.Projects(ctx.Ctx())
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
