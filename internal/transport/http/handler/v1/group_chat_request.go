package v1

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type GroupChatRequest struct {
	GroupRequestCache       *cache.GroupChatRequestCache
	GroupChatRequestUseCase *usecase.GroupChatRequestUseCase
	GroupChatMemberUseCase  *usecase.GroupChatMemberUseCase
	GroupChatUseCase        *usecase.GroupChatUseCase
	Redis                   *redis.Client
}

func (g *GroupChatRequest) Create(ctx *core.Context) error {
	params := &v1Pb.GroupChatRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindByWhere(ctx.Ctx(), "group_id = ? AND status = ?", params.GroupId, constant.GroupChatRequestStatusWait)
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	uid := ctx.UserId()

	if apply == nil {
		err = g.GroupChatRequestUseCase.GroupChatRequestRepo.Create(ctx.Ctx(), &model.GroupChatRequest{
			GroupId: int(params.GroupId),
			UserId:  uid,
			Status:  constant.GroupChatRequestStatusWait,
		})
	} else {
		data := map[string]interface{}{
			"updated_at": timeutil.DateTime(),
		}

		_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", apply.Id)
	}

	if err != nil {
		return ctx.Error(err.Error())
	}

	find, err := g.GroupChatMemberUseCase.MemberRepo.FindByWhere(ctx.Ctx(), "group_id = ? AND leader = ?", params.GroupId, 2)
	if err == nil && find != nil {
		g.GroupRequestCache.Incr(ctx.Ctx(), find.UserId)
	}

	g.Redis.Publish(ctx.Ctx(), constant.ImTopicChat, jsonutil.Encode(map[string]interface{}{
		"event": constant.SubEventGroupChatRequest,
		"data": jsonutil.Encode(map[string]interface{}{
			"group_id": params.GroupId,
			"user_id":  ctx.UserId(),
		}),
	}))

	return ctx.Success(nil)
}

func (g *GroupChatRequest) Agree(ctx *core.Context) error {
	uid := ctx.UserId()

	params := &v1Pb.GroupChatRequestAgreeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.ErrorBusiness("Информация о заявке не найдена")
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden("Нет доступа")
	}

	if apply.Status != constant.GroupChatRequestStatusWait {
		return ctx.ErrorBusiness("Информация о заявке уже обработана другим пользователем")
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), apply.GroupId, apply.UserId, false) {
		err = g.GroupChatUseCase.Invite(ctx.Ctx(), &usecase.GroupInviteOpt{
			UserId:    uid,
			GroupId:   apply.GroupId,
			MemberIds: []int{apply.UserId},
		})

		if err != nil {
			return ctx.ErrorBusiness(err.Error())
		}
	}

	data := map[string]interface{}{
		"status":     constant.GroupChatRequestStatusPass,
		"updated_at": timeutil.DateTime(),
	}

	_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (g *GroupChatRequest) Decline(ctx *core.Context) error {
	params := &v1Pb.GroupChatRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err.Error())
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.ErrorBusiness("Запись заявки не найдена")
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden("Нет доступа")
	}

	if apply.Status != constant.GroupChatRequestStatusWait {
		return ctx.ErrorBusiness("Информация о заявке уже обработана другим пользователем")
	}

	data := map[string]interface{}{
		"status":     constant.GroupChatRequestStatusRefuse,
		"updated_at": timeutil.DateTime(),
	}

	_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.GroupChatRequestDeclineResponse{})
}

func (g *GroupChatRequest) List(ctx *core.Context) error {
	params := &v1Pb.GroupRequestListRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), ctx.UserId()) {
		return ctx.Forbidden("Недостаточно прав доступа")
	}

	list, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.List(ctx.Ctx(), []int{int(params.GroupId)})
	if err != nil {
		return ctx.ErrorBusiness("Ошибка создания группового чата. Пожалуйста, попробуйте еще раз")
	}

	items := make([]*v1Pb.GroupChatRequestListResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.GroupChatRequestListResponse_Item{
			Id:        int32(item.Id),
			UserId:    int32(item.UserId),
			GroupId:   int32(item.GroupId),
			Avatar:    item.Avatar,
			Username:  item.Username,
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	return ctx.Success(&v1Pb.GroupChatRequestListResponse{Items: items})
}

func (g *GroupChatRequest) All(ctx *core.Context) error {
	uid := ctx.UserId()
	all, err := g.GroupChatMemberUseCase.MemberRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
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

	resp := &v1Pb.GroupChatRequestAllResponse{Items: make([]*v1Pb.GroupChatRequestAllResponse_Item, 0)}
	if len(groupIds) == 0 {
		return ctx.Success(resp)
	}

	list, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.List(ctx.Ctx(), groupIds)
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать групповой чат. Пожалуйста, попробуйте позже!")
	}

	groups, err := g.GroupChatUseCase.GroupChatRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
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
		resp.Items = append(resp.Items, &v1Pb.GroupChatRequestAllResponse_Item{
			Id:        int32(item.Id),
			UserId:    int32(item.UserId),
			GroupName: groupMap[item.GroupId].Name,
			GroupId:   int32(item.GroupId),
			Avatar:    item.Avatar,
			//Nickname:  item.Nickname,
			CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	g.GroupRequestCache.Del(ctx.Ctx(), ctx.UserId())

	return ctx.Success(resp)
}

func (g *GroupChatRequest) RequestUnreadNum(ctx *core.Context) error {
	return ctx.Success(map[string]any{
		"unread_num": g.GroupRequestCache.Get(ctx.Ctx(), ctx.UserId()),
	})
}
