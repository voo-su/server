package v1

import (
	"github.com/google/uuid"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type ProjectTask struct {
	Locale         locale.ILocale
	ProjectUseCase *usecase.ProjectUseCase
}

func (p *ProjectTask) Create(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectTaskCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	projectId, err := uuid.Parse(params.ProjectId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	taskId, err := p.ProjectUseCase.CreateTask(ctx.Ctx(), &usecase.ProjectTaskOpt{
		ProjectId:   projectId,
		TypeId:      int(params.TypeId),
		Title:       params.Title,
		Description: params.Description,
		CreatedBy:   ctx.UserId(),
	})
	if err != nil {
		return ctx.Error(p.Locale.Localize("creation_failed_try_later"))
	}

	return ctx.Success(&v1Pb.ProjectTaskCreateResponse{
		Id: taskId.String(),
	})
}

func (p *ProjectTask) Tasks(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectTaskRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	projectId, err := uuid.Parse(params.ProjectId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	data, err := p.ProjectUseCase.TypeTasks(ctx.Ctx(), projectId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	categories := make([]*v1Pb.ProjectTaskResponse_Categories, 0)
	for _, item := range data {

		tasks, _err := p.ProjectUseCase.Tasks(ctx.Ctx(), projectId, item.Id)
		if _err != nil {
			return ctx.Error(_err.Error())
		}

		taskItems := make([]*v1Pb.ProjectTaskResponse_Tasks, 0)
		for _, taskItem := range tasks {
			taskItems = append(taskItems, &v1Pb.ProjectTaskResponse_Tasks{
				Id:    taskItem.Id.String(),
				Title: taskItem.Title,
			})
		}

		categories = append(categories, &v1Pb.ProjectTaskResponse_Categories{
			Id:    item.Id,
			Title: item.Title,
			Tasks: taskItems,
		})
	}

	return ctx.Success(v1Pb.ProjectTaskResponse{
		Categories: categories,
	})
}

func (p *ProjectTask) TaskDetail(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectTaskDetailRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	task, err := p.ProjectUseCase.TaskDetail(ctx.Ctx(), taskId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(v1Pb.ProjectTaskDetailResponse{
		Id:          task.Id.String(),
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   timeutil.FormatDatetime(task.CreatedAt),
		Assigner: &v1Pb.ProjectTaskDetailResponse_Member{
			Id: task.AssignerId,
			//	Avatar: "",
			Username: task.AssignerUsername,
			Name:     task.AssignerName,
			Surname:  task.AssignerSurname,
		},
		Executor: &v1Pb.ProjectTaskDetailResponse_Member{
			Id: task.ExecutorId,
			//	Avatar: "",
			Username: task.ExecutorUsername,
			Name:     task.ExecutorName,
			Surname:  task.ExecutorSurname,
		},
	})
}

func (p *ProjectTask) Executor(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectExecutorRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	if !p.ProjectUseCase.IsMemberProjectByTask(ctx.Ctx(), taskId, uid) {
		return ctx.Error(p.Locale.Localize("not_project_member_cannot_invite"))
	}

	if err := p.ProjectUseCase.TaskExecutor(
		ctx.Ctx(),
		taskId,
		params.MemberId,
	); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(v1Pb.ProjectExecutorResponse{})
}

func (p *ProjectTask) TaskMove(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectTaskMoveRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	projectId, err := uuid.Parse(params.ProjectId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	if err := p.ProjectUseCase.TaskMove(
		ctx.Ctx(),
		projectId,
		taskId,
		params.FromId,
		params.ToId,
	); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(v1Pb.ProjectTaskMoveResponse{})
}

func (p *ProjectTask) TaskTypeName(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectTaskTypeNameRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	if err := p.ProjectUseCase.TaskTypeName(
		ctx.Ctx(),
		taskId,
		params.Name,
	); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(v1Pb.ProjectTaskTypeNameResponse{})
}

func (p *ProjectTask) TaskCoexecutorInvite(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectCoexecutorInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	mids := sliceutil.Unique(sliceutil.ParseIds(params.MemberIds))
	if len(mids) == 0 {
		return ctx.Error(p.Locale.Localize("invited_list_cannot_be_empty"))
	}

	uid := ctx.UserId()
	if !p.ProjectUseCase.IsMemberProjectByTask(ctx.Ctx(), taskId, uid) {
		return ctx.Error(p.Locale.Localize("not_project_member_cannot_invite"))
	}

	if err := p.ProjectUseCase.InviteCoexecutor(ctx.Ctx(), taskId, mids, uid); err != nil {
		return ctx.Error(p.Locale.Localize("failed_to_invite") + ": " + err.Error())
	}

	return ctx.Success(v1Pb.ProjectCoexecutorInviteResponse{})
}

func (p *ProjectTask) TaskCoexecutors(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectCoexecutorsRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	list, _ := p.ProjectUseCase.GetCoexecutors(ctx.Ctx(), taskId)

	items := make([]*v1Pb.ProjectCoexecutorsResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.ProjectCoexecutorsResponse_Item{
			Id: item.Id,
			//	Avatar: "",
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return ctx.Success(v1Pb.ProjectCoexecutorsResponse{
		Items: items,
	})
}

func (p *ProjectTask) TaskWatcherInvite(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectWatcherInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	mids := sliceutil.Unique(sliceutil.ParseIds(params.MemberIds))
	if len(mids) == 0 {
		return ctx.Error(p.Locale.Localize("invited_list_cannot_be_empty"))
	}

	uid := ctx.UserId()
	if !p.ProjectUseCase.IsMemberProjectByTask(ctx.Ctx(), taskId, uid) {
		return ctx.Error(p.Locale.Localize("not_project_member_cannot_invite"))
	}

	if err := p.ProjectUseCase.InviteWatcher(ctx.Ctx(), taskId, mids, uid); err != nil {
		return ctx.Error(p.Locale.Localize("failed_to_invite") + ": " + err.Error())
	}

	return ctx.Success(v1Pb.ProjectWatcherInviteResponse{})
}

func (p *ProjectTask) TaskWatchers(ctx *ginutil.Context) error {
	params := &v1Pb.ProjectWatchersRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	taskId, err := uuid.Parse(params.TaskId)
	if err != nil {
		return ctx.Error(p.Locale.Localize("network_error"))
	}

	list, _ := p.ProjectUseCase.GetWatchers(ctx.Ctx(), taskId)

	items := make([]*v1Pb.ProjectWatchersResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.ProjectWatchersResponse_Item{
			Id: item.Id,
			//	Avatar: "",
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return ctx.Success(v1Pb.ProjectWatchersResponse{
		Items: items,
	})
}
