package v1

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"voo.su/api/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type GroupChatRequest struct {
	GroupRequestStorage     *cache.GroupChatRequestStorage
	GroupChatRepo           *repo.GroupChat
	GroupChatRequestRepo    *repo.GroupChatRequest
	GroupMemberRepo         *repo.GroupChatMember
	GroupChatRequestService *service.GroupChatRequestService
	GroupChatMemberService  *service.GroupChatMemberService
	GroupChatService        *service.GroupChatService
	Redis                   *redis.Client
}

func (g *GroupChatRequest) Create(ctx *core.Context) error {
	params := &api_v1.GroupChatRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestRepo.FindByWhere(ctx.Ctx(), "group_id = ? and status = ?", params.GroupId, model.GroupChatRequestStatusWait)
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	uid := ctx.UserId()

	if apply == nil {
		err = g.GroupChatRequestRepo.Create(ctx.Ctx(), &model.GroupChatRequest{
			GroupId: int(params.GroupId),
			UserId:  uid,
			Status:  model.GroupChatRequestStatusWait,
			Remark:  params.Remark,
		})
	} else {
		data := map[string]interface{}{
			"remark":     params.Remark,
			"updated_at": timeutil.DateTime(),
		}

		_, err = g.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", apply.Id)
	}

	if err != nil {
		return ctx.Error(err.Error())
	}

	find, err := g.GroupMemberRepo.FindByWhere(ctx.Ctx(), "group_id = ? and leader = ?", params.GroupId, 2)
	if err == nil && find != nil {
		g.GroupRequestStorage.Incr(ctx.Ctx(), find.UserId)
	}

	g.Redis.Publish(ctx.Ctx(), entity.ImTopicChat, jsonutil.Encode(map[string]interface{}{
		"event": entity.SubEventGroupChatRequest,
		"data": jsonutil.Encode(map[string]interface{}{
			"group_id": params.GroupId,
			"user_id":  ctx.UserId(),
		}),
	}))

	return ctx.Success(nil)
}

func (g *GroupChatRequest) Agree(ctx *core.Context) error {
	uid := ctx.UserId()

	params := &api_v1.GroupChatRequestAgreeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.ErrorBusiness("Информация о заявке не найдена")
	}

	if !g.GroupMemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden("Нет доступа")
	}

	if apply.Status != model.GroupChatRequestStatusWait {
		return ctx.ErrorBusiness("Информация о заявке уже обработана другим пользователем")
	}

	if !g.GroupMemberRepo.IsMember(ctx.Ctx(), apply.GroupId, apply.UserId, false) {
		err = g.GroupChatService.Invite(ctx.Ctx(), &service.GroupInviteOpt{
			UserId:    uid,
			GroupId:   apply.GroupId,
			MemberIds: []int{apply.UserId},
		})

		if err != nil {
			return ctx.ErrorBusiness(err.Error())
		}
	}

	data := map[string]interface{}{
		"status":     model.GroupChatRequestStatusPass,
		"updated_at": timeutil.DateTime(),
	}

	_, err = g.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (g *GroupChatRequest) Decline(ctx *core.Context) error {
	params := &api_v1.GroupChatRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	apply, err := g.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.ErrorBusiness("Запись заявки не найдена")
	}

	if !g.GroupMemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden("Нет доступа")
	}

	if apply.Status != model.GroupChatRequestStatusWait {
		return ctx.ErrorBusiness("Информация о заявке уже обработана другим пользователем")
	}

	data := map[string]interface{}{
		"status":     model.GroupChatRequestStatusRefuse,
		"reason":     params.Remark,
		"updated_at": timeutil.DateTime(),
	}

	_, err = g.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&api_v1.GroupChatRequestDeclineResponse{})
}

func (g *GroupChatRequest) List(ctx *core.Context) error {
	params := &api_v1.GroupRequestListRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !g.GroupMemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), ctx.UserId()) {
		return ctx.Forbidden("Недостаточно прав доступа")
	}

	list, err := g.GroupChatRequestRepo.List(ctx.Ctx(), []int{int(params.GroupId)})
	if err != nil {
		return ctx.ErrorBusiness("Ошибка создания группового чата. Пожалуйста, попробуйте еще раз")
	}

	items := make([]*api_v1.GroupChatRequestListResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.GroupChatRequestListResponse_Item{
			Id:      int32(item.Id),
			UserId:  int32(item.UserId),
			GroupId: int32(item.GroupId),
			//Remark:    item.Remark,
			Avatar:    item.Avatar,
			Username:  item.Username,
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	return ctx.Success(&api_v1.GroupChatRequestListResponse{Items: items})
}

func (g *GroupChatRequest) All(ctx *core.Context) error {
	uid := ctx.UserId()
	all, err := g.GroupMemberRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
		db.Select("group_id")
		db.Where("user_id = ?", uid)
		db.Where("leader = ?", 2)
		db.Where("is_quit = ?", 0)
	})

	if err != nil {
		return ctx.ErrorBusiness("Системная ошибка. Пожалуйста, попробуйте позже!")
	}

	groupIds := make([]int, 0, len(all))
	for _, m := range all {
		groupIds = append(groupIds, m.GroupId)
	}

	resp := &api_v1.GroupChatRequestAllResponse{Items: make([]*api_v1.GroupChatRequestAllResponse_Item, 0)}
	if len(groupIds) == 0 {
		return ctx.Success(resp)
	}

	list, err := g.GroupChatRequestRepo.List(ctx.Ctx(), groupIds)
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать групповой чат. Пожалуйста, попробуйте позже!")
	}

	groups, err := g.GroupChatRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
		db.Select("id,group_name")
		db.Where("id in ?", groupIds)
	})
	if err != nil {
		return err
	}

	groupMap := sliceutil.ToMap(groups, func(t *model.GroupChat) int {
		return t.Id
	})

	for _, item := range list {
		resp.Items = append(resp.Items, &api_v1.GroupChatRequestAllResponse_Item{
			Id:        int32(item.Id),
			UserId:    int32(item.UserId),
			GroupName: groupMap[item.GroupId].Name,
			GroupId:   int32(item.GroupId),
			//Remark:    item.Remark,
			Avatar: item.Avatar,
			//Nickname:  item.Nickname,
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	g.GroupRequestStorage.Del(ctx.Ctx(), ctx.UserId())

	return ctx.Success(resp)
}

func (g *GroupChatRequest) RequestUnreadNum(ctx *core.Context) error {
	return ctx.Success(map[string]any{
		"unread_num": g.GroupRequestStorage.Get(ctx.Ctx(), ctx.UserId()),
	})
}
