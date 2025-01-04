package v1

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type GroupChatRequest struct {
	Locale                  locale.ILocale
	GroupRequestCacheRepo   *redisRepo.GroupChatRequestCacheRepository
	GroupChatRequestUseCase *usecase.GroupChatRequestUseCase
	GroupChatMemberUseCase  *usecase.GroupChatMemberUseCase
	GroupChatUseCase        *usecase.GroupChatUseCase
	Redis                   *redis.Client
}

func (g *GroupChatRequest) Create(ctx *ginutil.Context) error {
	params := &v1Pb.GroupChatRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindByWhere(ctx.Ctx(), "group_id = ? AND status = ?", params.GroupId, constant.GroupChatRequestStatusWait)
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err)
	}

	uid := ctx.UserId()

	if apply == nil {
		err = g.GroupChatRequestUseCase.GroupChatRequestRepo.Create(ctx.Ctx(), &postgresModel.GroupChatRequest{
			GroupId: int(params.GroupId),
			UserId:  uid,
			Status:  constant.GroupChatRequestStatusWait,
		})
	} else {
		_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), map[string]interface{}{
			"updated_at": timeutil.DateTime(),
		}, "id = ?", apply.Id)
	}

	if err != nil {
		return ctx.Error(err)
	}

	find, err := g.GroupChatMemberUseCase.MemberRepo.FindByWhere(ctx.Ctx(), "group_id = ? AND leader = ?", params.GroupId, 2)
	if err == nil && find != nil {
		g.GroupRequestCacheRepo.Incr(ctx.Ctx(), find.UserId)
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

func (g *GroupChatRequest) Agree(ctx *ginutil.Context) error {
	uid := ctx.UserId()

	params := &v1Pb.GroupChatRequestAgreeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err)
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.Error(g.Locale.Localize("data_not_found"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden(g.Locale.Localize("access_rights_missing"))
	}

	if apply.Status != constant.GroupChatRequestStatusWait {
		return ctx.Error(g.Locale.Localize("request_already_processed"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), apply.GroupId, apply.UserId, false) {
		err = g.GroupChatUseCase.Invite(ctx.Ctx(), &usecase.GroupInviteOpt{
			UserId:    uid,
			GroupId:   apply.GroupId,
			MemberIds: []int{apply.UserId},
		})

		if err != nil {
			return ctx.Error(err.Error())
		}
	}

	_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), map[string]interface{}{
		"status":     constant.GroupChatRequestStatusPass,
		"updated_at": timeutil.DateTime(),
	}, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Success(nil)
}

func (g *GroupChatRequest) Decline(ctx *ginutil.Context) error {
	params := &v1Pb.GroupChatRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	apply, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.FindById(ctx.Ctx(), int(params.ApplyId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return ctx.Error(err)
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.Error(g.Locale.Localize("data_not_found"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), apply.GroupId, uid) {
		return ctx.Forbidden(g.Locale.Localize("access_rights_missing"))
	}

	if apply.Status != constant.GroupChatRequestStatusWait {
		return ctx.Error(g.Locale.Localize("request_already_processed"))
	}

	_, err = g.GroupChatRequestUseCase.GroupChatRequestRepo.UpdateWhere(ctx.Ctx(), map[string]interface{}{
		"status":     constant.GroupChatRequestStatusRefuse,
		"updated_at": timeutil.DateTime(),
	}, "id = ?", params.ApplyId)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Success(&v1Pb.GroupChatRequestDeclineResponse{})
}

func (g *GroupChatRequest) List(ctx *ginutil.Context) error {
	params := &v1Pb.GroupRequestListRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), ctx.UserId()) {
		return ctx.Forbidden(g.Locale.Localize("access_rights_missing"))
	}

	list, err := g.GroupChatRequestUseCase.GroupChatRequestRepo.List(ctx.Ctx(), []int{int(params.GroupId)})
	if err != nil {
		return ctx.Error(g.Locale.Localize("chat_creation_error"))
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

func (g *GroupChatRequest) All(ctx *ginutil.Context) error {
	uid := ctx.UserId()
	all, err := g.GroupChatMemberUseCase.MemberRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
		db.Select("group_id").
			Where("user_id = ?", uid).
			Where("leader = ?", 2).
			Where("is_quit = ?", 0)
	})

	if err != nil {
		return ctx.Error(g.Locale.Localize("network_error"))
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
		return ctx.Error(g.Locale.Localize("chat_creation_error"))
	}

	groups, err := g.GroupChatUseCase.GroupChatRepo.FindAll(ctx.Ctx(), func(db *gorm.DB) {
		db.Select("id, group_name").
			Where("id IN ?", groupIds)
	})
	if err != nil {
		return err
	}

	groupMap := sliceutil.ToMap(groups, func(t *postgresModel.GroupChat) int {
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

	g.GroupRequestCacheRepo.Del(ctx.Ctx(), ctx.UserId())

	return ctx.Success(resp)
}

func (g *GroupChatRequest) RequestUnreadNum(ctx *ginutil.Context) error {
	return ctx.Success(map[string]any{
		"unread_num": g.GroupRequestCacheRepo.Get(ctx.Ctx(), ctx.UserId()),
	})
}
