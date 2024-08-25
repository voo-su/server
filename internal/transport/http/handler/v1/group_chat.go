package v1

import (
	"fmt"
	"voo.su/api/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/logger"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type GroupChat struct {
	Repo                   *repo.Source
	UserRepo               *repo.User
	GroupChatRepo          *repo.GroupChat
	GroupChatMemberRepo    *repo.GroupChatMember
	DialogRepo             *repo.Dialog
	GroupChatService       *service.GroupChatService
	GroupChatMemberService *service.GroupChatMemberService
	DialogService          *service.DialogService
	ContactService         *service.ContactService
	MessageSendService     service.MessageSendService
	RedisLock              *cache.RedisLock
}

func (c *GroupChat) Create(ctx *core.Context) error {
	params := &api_v1.GroupChatCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	gid, err := c.GroupChatService.Create(ctx.Ctx(), &service.GroupCreateOpt{
		UserId:    ctx.UserId(),
		Name:      params.Name,
		Avatar:    params.Avatar,
		MemberIds: sliceutil.ParseIds(params.GetIds()),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать групповой чат, попробуйте позже" + err.Error())
	}

	return ctx.Success(&api_v1.GroupChatCreateResponse{GroupId: int32(gid)})
}

func (c *GroupChat) Invite(ctx *core.Context) error {
	params := &api_v1.GroupChatInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	key := fmt.Sprintf("group-join:%d", params.GroupId)
	if !c.RedisLock.Lock(ctx.Ctx(), key, 20) {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	defer c.RedisLock.UnLock(ctx.Ctx(), key)
	group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	if group != nil && group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}

	uid := ctx.UserId()
	uids := sliceutil.Unique(sliceutil.ParseIds(params.Ids))
	if len(uids) == 0 {
		return ctx.ErrorBusiness("Список приглашенных друзей не может быть пустым")
	}

	if !c.GroupChatMemberRepo.IsMember(ctx.Ctx(), int(params.GroupId), uid, true) {
		return ctx.ErrorBusiness("Вы не являетесь участником группы и не имеете права приглашать друзей")
	}

	if err := c.GroupChatService.Invite(ctx.Ctx(), &service.GroupInviteOpt{
		UserId:    uid,
		GroupId:   int(params.GroupId),
		MemberIds: uids,
	}); err != nil {
		return ctx.ErrorBusiness("Не удалось пригласить друзей в групповой чат" + err.Error())
	}

	return ctx.Success(&api_v1.GroupChatInviteResponse{})
}

func (c *GroupChat) SignOut(ctx *core.Context) error {
	params := &api_v1.GroupChatSecedeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := c.GroupChatService.Secede(ctx.Ctx(), int(params.GroupId), uid); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	sid := c.DialogRepo.FindBySessionId(uid, int(params.GroupId), entity.ChatGroupMode)
	_ = c.DialogService.Delete(ctx.Ctx(), ctx.UserId(), sid)

	return ctx.Success(nil)
}

func (c *GroupChat) Setting(ctx *core.Context) error {
	params := &api_v1.GroupSettingRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	if group != nil && group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}

	uid := ctx.UserId()
	if !c.GroupChatMemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав для выполнения этой операции")
	}

	if err := c.GroupChatService.Update(ctx.Ctx(), &service.GroupUpdateOpt{
		GroupId:     int(params.GroupId),
		Name:        params.GroupName,
		Avatar:      params.Avatar,
		Description: params.Description,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	//_ = c.messageService.SendSystemText(ctx.Ctx(), uid, &api_v1.TextMessageRequest{
	//	Content: "Владелец группы или администратор изменили информацию о группе",
	//	Receiver: &api_v1.MessageReceiver{
	//		DialogType: entity.ChatPrivateMode,
	//		ReceiverId: params.GroupId,
	//	},
	//})

	return ctx.Success(&api_v1.GroupChatSettingResponse{})
}

func (c *GroupChat) RemoveMembers(ctx *core.Context) error {
	params := &api_v1.GroupChatRemoveMemberRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !c.GroupChatMemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав для выполнения этой операции")
	}

	err := c.GroupChatService.RemoveMember(ctx.Ctx(), &service.GroupRemoveMembersOpt{
		UserId:    uid,
		GroupId:   int(params.GroupId),
		MemberIds: sliceutil.ParseIds(params.MembersIds),
	})

	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.GroupChatRemoveMemberResponse{})
}

func (c *GroupChat) Get(ctx *core.Context) error {
	params := &api_v1.GroupChatDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	groupInfo, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if groupInfo.Id == 0 {
		return ctx.ErrorBusiness("Данные не существуют")
	}

	resp := &api_v1.GroupChatDetailResponse{
		GroupId:     int32(groupInfo.Id),
		GroupName:   groupInfo.Name,
		Description: groupInfo.Description,
		Avatar:      groupInfo.Avatar,
		CreatedAt:   timeutil.FormatDatetime(groupInfo.CreatedAt),
		IsManager:   uid == groupInfo.CreatorId,
		IsDisturb:   0,
		IsMute:      int32(groupInfo.IsMute),
		IsOvert:     int32(groupInfo.IsOvert),
		//VisitCard: c.GroupMemberRepo.GetMemberRemark(ctx.Ctx(), int(params.GroupId), uid),
	}

	if c.DialogRepo.IsDisturb(uid, groupInfo.Id, 2) {
		resp.IsDisturb = 1
	}

	return ctx.Success(resp)
}

//func (c *GroupChat) UpdateMemberRemark(ctx *core.Context) error {
//	params := &api_v1.GroupChatRemarkUpdateRequest{}
//	if err := ctx.Context.ShouldBind(params); err != nil {
//		return ctx.InvalidParams(err)
//	}
//
//	_, err := c.GroupChatMemberRepo.UpdateWhere(ctx.Ctx(), map[string]any{
//		"user_card": params.VisitCard,
//	}, "group_id = ? and user_id = ?", params.GroupId, ctx.UserId())
//	if err != nil {
//		return ctx.ErrorBusiness("Не удалось изменить заметку в группе")
//	}
//
//	return ctx.Success(nil)
//}

func (c *GroupChat) GetInviteFriends(ctx *core.Context) error {
	params := &api_v1.GroupChatGetInviteFriendsRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	items, err := c.ContactService.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if params.GroupId <= 0 {
		return ctx.Success(items)
	}

	mids := c.GroupChatMemberRepo.GetMemberIds(ctx.Ctx(), int(params.GroupId))
	if len(mids) == 0 {
		return ctx.Success(items)
	}

	data := make([]*model.ContactListItem, 0)
	for i := 0; i < len(items); i++ {
		if !sliceutil.Include(items[i].Id, mids) {
			data = append(data, items[i])
		}
	}

	return ctx.Success(data)
}

func (c *GroupChat) GroupList(ctx *core.Context) error {
	items, err := c.GroupChatService.List(ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	resp := &api_v1.GroupChatListResponse{
		Items: make([]*api_v1.GroupChatListResponse_Item, 0, len(items)),
	}
	for _, item := range items {
		resp.Items = append(resp.Items, &api_v1.GroupChatListResponse_Item{
			Id:          int32(item.Id),
			GroupName:   item.GroupName,
			Avatar:      item.Avatar,
			Description: item.Description,
			Leader:      int32(item.Leader),
			IsDisturb:   int32(item.IsDisturb),
			CreatorId:   int32(item.CreatorId),
		})
	}

	return ctx.Success(resp)
}

func (c *GroupChat) Members(ctx *core.Context) error {
	params := &api_v1.GroupChatMemberListRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, попробуйте позже")
	}

	if group != nil && group.IsDismiss == 1 {
		return ctx.Success([]any{})
	}

	if !c.GroupChatMemberRepo.IsMember(ctx.Ctx(), int(params.GroupId), ctx.UserId(), false) {
		return ctx.ErrorBusiness("Не являетесь членом группы и не имеете права просматривать список участников")
	}

	list := c.GroupChatMemberRepo.GetMembers(ctx.Ctx(), int(params.GroupId))

	items := make([]*api_v1.GroupChatMemberListResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.GroupChatMemberListResponse_Item{
			UserId:   int32(item.UserId),
			Username: item.Username,
			Avatar:   item.Avatar,
			Gender:   int32(item.Gender),
			Leader:   int32(item.Leader),
			IsMute:   int32(item.IsMute),
			//Remark:   item.UserCard,
		})
	}

	return ctx.Success(&api_v1.GroupChatMemberListResponse{Items: items})
}

func (c *GroupChat) AssignAdmin(ctx *core.Context) error {
	params := &api_v1.GroupChatAssignAdminRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !c.GroupChatMemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("Отсутствуют права доступа")
	}

	leader := 0
	if params.Mode == 1 {
		leader = 1
	}

	err := c.GroupChatMemberService.SetLeaderStatus(ctx.Ctx(), int(params.GroupId), int(params.UserId), leader)
	if err != nil {
		logger.Errorf("Не удалось установить информацию администратора:%s", err.Error())
		return ctx.ErrorBusiness("Не удалось установить информацию администратора!")
	}

	return ctx.Success(nil)
}

func (c *GroupChat) Dismiss(ctx *core.Context) error {
	params := &api_v1.GroupChatDismissRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !c.GroupChatMemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав на расформирование группы")
	}
	if err := c.GroupChatService.Dismiss(ctx.Ctx(), int(params.GroupId), ctx.UserId()); err != nil {
		return ctx.ErrorBusiness("Не удалось расформировать группу")
	}

	_ = c.MessageSendService.SendSystemText(ctx.Ctx(), uid, &api_v1.TextMessageRequest{
		Content: "Группа была расформирована владельцем группы",
		Receiver: &api_v1.MessageReceiver{
			DialogType: entity.ChatGroupMode,
			ReceiverId: params.GroupId,
		},
	})

	return ctx.Success(nil)
}

func (c *GroupChat) Mute(ctx *core.Context) error {
	params := &api_v1.GroupChatMuteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети. Пожалуйста, попробуйте еще раз")
	}

	if group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}
	if !c.GroupChatMemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав доступа")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_mute"] = 1
	} else {
		data["is_mute"] = 0
	}

	affected, err := c.GroupChatRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.GroupId)
	if err != nil {
		return ctx.Error("Серверная ошибка. Пожалуйста, попробуйте еще раз")
	}
	if affected == 0 {
		return ctx.Success(api_v1.GroupChatMuteResponse{})
	}

	user, err := c.UserRepo.FindById(ctx.Ctx(), uid)
	if err != nil {
		return err
	}

	var extra any
	var msgType int
	if params.Mode == 1 {
		msgType = entity.ChatMsgSysGroupMuted
		extra = model.DialogRecordExtraGroupMuted{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}
	} else {
		msgType = entity.ChatMsgSysGroupCancelMuted
		extra = model.DialogRecordExtraGroupCancelMuted{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}
	}

	_ = c.MessageSendService.SendSysOther(ctx.Ctx(), &model.Message{
		MsgType:    msgType,
		DialogType: model.DialogRecordDialogTypeGroup,
		UserId:     uid,
		ReceiverId: int(params.GroupId),
		Extra:      jsonutil.Encode(extra),
	})

	return ctx.Success(api_v1.GroupChatMuteResponse{})
}

func (c *GroupChat) Overt(ctx *core.Context) error {
	params := &api_v1.GroupChatOvertRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети. Пожалуйста, попробуйте еще раз")
	}

	if group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}

	if !c.GroupChatMemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав доступа")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_overt"] = 1
	} else {
		data["is_overt"] = 0
	}

	_, err = c.GroupChatRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.GroupId)
	if err != nil {
		return ctx.Error("Серверная ошибка. Пожалуйста, попробуйте еще раз")
	}

	return ctx.Success(api_v1.GroupChatOvertResponse{})
}
