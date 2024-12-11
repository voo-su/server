package v1

import (
	"fmt"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/logger"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type GroupChat struct {
	GroupChatUseCase       *usecase.GroupChatUseCase
	GroupChatMemberUseCase *usecase.GroupChatMemberUseCase
	ChatUseCase            *usecase.ChatUseCase
	ContactUseCase         *usecase.ContactUseCase
	MessageUseCase         usecase.IMessageUseCase
	UserUseCase            *usecase.UserUseCase
	RedisLockCacheRepo     *redisRepo.RedisLockCacheRepository
}

func (g *GroupChat) Create(ctx *core.Context) error {
	params := &v1Pb.GroupChatCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	gid, err := g.GroupChatUseCase.Create(ctx.Ctx(), &usecase.GroupCreateOpt{
		UserId:    ctx.UserId(),
		Name:      params.Name,
		Avatar:    params.Avatar,
		MemberIds: sliceutil.ParseIds(params.GetIds()),
	})
	if err != nil {
		return ctx.ErrorBusiness("Не удалось создать групповой чат, попробуйте позже" + err.Error())
	}

	return ctx.Success(&v1Pb.GroupChatCreateResponse{GroupId: int32(gid)})
}

func (g *GroupChat) Invite(ctx *core.Context) error {
	params := &v1Pb.GroupChatInviteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	key := fmt.Sprintf("group-join:%d", params.GroupId)
	if !g.RedisLockCacheRepo.Lock(ctx.Ctx(), key, 20) {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	defer g.RedisLockCacheRepo.UnLock(ctx.Ctx(), key)
	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
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

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), int(params.GroupId), uid, true) {
		return ctx.ErrorBusiness("Вы не являетесь участником группы и не имеете права приглашать друзей")
	}

	if err := g.GroupChatUseCase.Invite(ctx.Ctx(), &usecase.GroupInviteOpt{
		UserId:    uid,
		GroupId:   int(params.GroupId),
		MemberIds: uids,
	}); err != nil {
		return ctx.ErrorBusiness("Не удалось пригласить друзей в групповой чат" + err.Error())
	}

	return ctx.Success(&v1Pb.GroupChatInviteResponse{})
}

func (g *GroupChat) SignOut(ctx *core.Context) error {
	params := &v1Pb.GroupChatSecedeRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := g.GroupChatUseCase.Secede(ctx.Ctx(), int(params.GroupId), uid); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	sid := g.ChatUseCase.ChatRepo.FindBySessionId(uid, int(params.GroupId), constant.ChatGroupMode)
	_ = g.ChatUseCase.Delete(ctx.Ctx(), ctx.UserId(), sid)

	return ctx.Success(nil)
}

func (g *GroupChat) Setting(ctx *core.Context) error {
	params := &v1Pb.GroupSettingRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, повторите попытку позже")
	}

	if group != nil && group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}

	uid := ctx.UserId()
	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав для выполнения этой операции")
	}

	if err := g.GroupChatUseCase.Update(ctx.Ctx(), &usecase.GroupUpdateOpt{
		GroupId:     int(params.GroupId),
		Name:        params.GroupName,
		Avatar:      params.Avatar,
		Description: params.Description,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	//_ = g.MessageUseCase.SendSystemText(ctx.Ctx(), uid, &v1Pb.TextMessageRequest{
	//	Content: "Владелец группы или администратор изменили информацию о группе",
	//	Receiver: &v1Pb.MessageReceiver{
	//		DialogType: constant.ChatPrivateMode,
	//		ReceiverId: params.GroupId,
	//	},
	//})

	return ctx.Success(&v1Pb.GroupChatSettingResponse{})
}

func (g *GroupChat) Get(ctx *core.Context) error {
	params := &v1Pb.GroupChatDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	groupInfo, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if groupInfo.Id == 0 {
		return ctx.ErrorBusiness("Данные не существуют")
	}

	resp := &v1Pb.GroupChatDetailResponse{
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

	if g.ChatUseCase.ChatRepo.IsDisturb(uid, groupInfo.Id, 2) {
		resp.IsDisturb = 1
	}

	return ctx.Success(resp)
}

func (g *GroupChat) GetInviteFriends(ctx *core.Context) error {
	params := &v1Pb.GroupChatGetInviteFriendsRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	items, err := g.ContactUseCase.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if params.GroupId <= 0 {
		return ctx.Success(items)
	}

	mids := g.GroupChatUseCase.MemberRepo.GetMemberIds(ctx.Ctx(), int(params.GroupId))
	if len(mids) == 0 {
		return ctx.Success(items)
	}

	data := make([]*entity.ContactListItem, 0)
	for i := 0; i < len(items); i++ {
		if !sliceutil.Include(items[i].Id, mids) {
			data = append(data, items[i])
		}
	}

	return ctx.Success(data)
}

func (g *GroupChat) GroupList(ctx *core.Context) error {
	items, err := g.GroupChatUseCase.List(ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	resp := &v1Pb.GroupChatListResponse{
		Items: make([]*v1Pb.GroupChatListResponse_Item, 0, len(items)),
	}
	for _, item := range items {
		resp.Items = append(resp.Items, &v1Pb.GroupChatListResponse_Item{
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

func (g *GroupChat) Members(ctx *core.Context) error {
	params := &v1Pb.GroupChatMemberListRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети, попробуйте позже")
	}

	if group != nil && group.IsDismiss == 1 {
		return ctx.Success([]any{})
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), int(params.GroupId), ctx.UserId(), false) {
		return ctx.ErrorBusiness("Не являетесь членом группы и не имеете права просматривать список участников")
	}

	list := g.GroupChatMemberUseCase.MemberRepo.GetMembers(ctx.Ctx(), int(params.GroupId))

	items := make([]*v1Pb.GroupChatMemberListResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.GroupChatMemberListResponse_Item{
			UserId:   int32(item.UserId),
			Username: item.Username,
			Avatar:   item.Avatar,
			Gender:   int32(item.Gender),
			Leader:   int32(item.Leader),
			IsMute:   int32(item.IsMute),
			//Remark:   item.UserCard,
		})
	}

	return ctx.Success(&v1Pb.GroupChatMemberListResponse{Items: items})
}

func (g *GroupChat) RemoveMembers(ctx *core.Context) error {
	params := &v1Pb.GroupChatRemoveMemberRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав для выполнения этой операции")
	}

	err := g.GroupChatUseCase.RemoveMember(ctx.Ctx(), &usecase.GroupRemoveMembersOpt{
		UserId:    uid,
		GroupId:   int(params.GroupId),
		MemberIds: sliceutil.ParseIds(params.MembersIds),
	})

	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.GroupChatRemoveMemberResponse{})
}

func (g *GroupChat) AssignAdmin(ctx *core.Context) error {
	params := &v1Pb.GroupChatAssignAdminRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !g.GroupChatMemberUseCase.MemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("Отсутствуют права доступа")
	}

	leader := 0
	if params.Mode == 1 {
		leader = 1
	}

	if err := g.GroupChatMemberUseCase.SetLeaderStatus(ctx.Ctx(), int(params.GroupId), int(params.UserId), leader); err != nil {
		logger.Errorf("Не удалось установить информацию администратора:%s", err.Error())
		return ctx.ErrorBusiness("Не удалось установить информацию администратора!")
	}

	return ctx.Success(nil)
}

func (g *GroupChat) Dismiss(ctx *core.Context) error {
	params := &v1Pb.GroupChatDismissRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !g.GroupChatMemberUseCase.MemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав на расформирование группы")
	}
	if err := g.GroupChatUseCase.Dismiss(ctx.Ctx(), int(params.GroupId), ctx.UserId()); err != nil {
		return ctx.ErrorBusiness("Не удалось расформировать группу")
	}

	_ = g.MessageUseCase.SendSystemText(ctx.Ctx(), uid, &v1Pb.TextMessageRequest{
		Content: "Группа была расформирована владельцем группы",
		Receiver: &v1Pb.MessageReceiver{
			DialogType: constant.ChatGroupMode,
			ReceiverId: params.GroupId,
		},
	})

	return ctx.Success(nil)
}

func (g *GroupChat) Mute(ctx *core.Context) error {
	params := &v1Pb.GroupChatMuteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети. Пожалуйста, попробуйте еще раз")
	}

	if group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}
	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав доступа")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_mute"] = 1
	} else {
		data["is_mute"] = 0
	}

	affected, err := g.GroupChatUseCase.GroupChatRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.GroupId)
	if err != nil {
		return ctx.Error("Серверная ошибка. Пожалуйста, попробуйте еще раз")
	}
	if affected == 0 {
		return ctx.Success(v1Pb.GroupChatMuteResponse{})
	}

	user, err := g.UserUseCase.UserRepo.FindById(ctx.Ctx(), uid)
	if err != nil {
		return err
	}

	var extra any
	var msgType int
	if params.Mode == 1 {
		msgType = constant.ChatMsgSysGroupMuted
		extra = entity.DialogRecordExtraGroupMuted{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}
	} else {
		msgType = constant.ChatMsgSysGroupCancelMuted
		extra = entity.DialogRecordExtraGroupCancelMuted{
			OwnerId:   user.Id,
			OwnerName: user.Username,
		}
	}

	if err := g.MessageUseCase.SendSysOther(ctx.Ctx(), &postgresModel.Message{
		MsgType:    msgType,
		DialogType: constant.DialogRecordDialogTypeGroup,
		UserId:     uid,
		ReceiverId: int(params.GroupId),
		Extra:      jsonutil.Encode(extra),
	}); err != nil {
		fmt.Println(err)
	}

	return ctx.Success(v1Pb.GroupChatMuteResponse{})
}

func (g *GroupChat) Overt(ctx *core.Context) error {
	params := &v1Pb.GroupChatOvertRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness("Ошибка сети. Пожалуйста, попробуйте еще раз")
	}

	if group.IsDismiss == 1 {
		return ctx.ErrorBusiness("Эта группа была расформирована")
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMaster(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав доступа")
	}

	data := make(map[string]any)
	if params.Mode == 1 {
		data["is_overt"] = 1
	} else {
		data["is_overt"] = 0
	}

	_, err = g.GroupChatUseCase.GroupChatRepo.UpdateWhere(ctx.Ctx(), data, "id = ?", params.GroupId)
	if err != nil {
		return ctx.Error("Серверная ошибка. Пожалуйста, попробуйте еще раз")
	}

	return ctx.Success(v1Pb.GroupChatOvertResponse{})
}
