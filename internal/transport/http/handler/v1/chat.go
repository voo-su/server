package v1

import (
	"fmt"
	"strconv"
	"strings"
	api_v1 "voo.su/api/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type Chat struct {
	ChatService      *service.ChatService
	RedisLock        *cache.RedisLock
	ClientStorage    *cache.ClientStorage
	MessageStorage   *cache.MessageStorage
	ContactService   *service.ContactService
	UnreadStorage    *cache.UnreadStorage
	ContactRemark    *cache.ContactRemark
	GroupChatService *service.GroupChatService
	AuthService      *service.AuthService
	ContactRepo      *repo.Contact
	UserRepo         *repo.User
	GroupChatRepo    *repo.GroupChat
}

func (c *Chat) Create(ctx *core.Context) error {
	var (
		params = &api_v1.ChatCreateRequest{}
		uid    = ctx.UserId()
		agent  = strings.TrimSpace(ctx.Context.GetHeader("user-agent"))
	)
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if agent != "" {
		agent = encrypt.Md5(agent)
	}

	if params.DialogType == entity.ChatPrivateMode && int(params.ReceiverId) == ctx.UserId() {
		return ctx.ErrorBusiness("Ошибка создания")
	}

	key := fmt.Sprintf("dialog:list:%d-%d-%d-%s", uid, params.ReceiverId, params.DialogType, agent)
	if !c.RedisLock.Lock(ctx.Ctx(), key, 10) {
		return ctx.ErrorBusiness("Ошибка создания")
	}

	if c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType: int(params.DialogType),
		UserId:     uid,
		ReceiverId: int(params.ReceiverId),
	}) != nil {
		return ctx.ErrorBusiness("Недостаточно прав")
	}

	result, err := c.ChatService.Create(ctx.Ctx(), &service.CreateChatOpt{
		UserId:     uid,
		DialogType: int(params.DialogType),
		ReceiverId: int(params.ReceiverId),
	})
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	item := &api_v1.ChatItem{
		Id:         int32(result.Id),
		DialogType: int32(result.DialogType),
		ReceiverId: int32(result.ReceiverId),
		IsBot:      int32(result.IsBot),
		UpdatedAt:  timeutil.DateTime(),
	}
	if item.DialogType == entity.ChatPrivateMode {
		item.UnreadNum = int32(c.UnreadStorage.Get(ctx.Ctx(), 1, int(params.ReceiverId), uid))
		//item.Remark = c.contactService.Dao().GetFriendRemark(ctx.Ctx(), uid, int(params.ReceiverId))
		if user, err := c.UserRepo.FindById(ctx.Ctx(), result.ReceiverId); err == nil {
			item.Username = user.Username
			item.Name = user.Name
			item.Surname = user.Surname
			item.Avatar = user.Avatar
		}
	} else if result.DialogType == entity.ChatGroupMode {
		if group, err := c.GroupChatRepo.FindById(ctx.Ctx(), int(params.ReceiverId)); err == nil {
			item.Name = group.Name
		}
	}

	if msg, err := c.MessageStorage.Get(ctx.Ctx(), result.DialogType, uid, result.ReceiverId); err == nil {
		item.MsgText = msg.Content
		item.UpdatedAt = msg.Datetime
	}

	return ctx.Success(&api_v1.ChatCreateResponse{
		Id:         item.Id,
		DialogType: item.DialogType,
		ReceiverId: item.ReceiverId,
		IsTop:      item.IsTop,
		IsDisturb:  item.IsDisturb,
		IsOnline:   item.IsOnline,
		IsBot:      item.IsBot,
		Username:   item.Username,
		Name:       item.Name,
		Surname:    item.Surname,
		Avatar:     item.Avatar,
		//RemarkName: item.Remark,
		UnreadNum: item.UnreadNum,
		MsgText:   item.MsgText,
		UpdatedAt: item.UpdatedAt,
	})
}

func (c *Chat) List(ctx *core.Context) error {
	uid := ctx.UserId()
	unReads := c.UnreadStorage.All(ctx.Ctx(), uid)
	if len(unReads) > 0 {
		c.ChatService.BatchAddList(ctx.Ctx(), uid, unReads)
	}

	data, err := c.ChatService.List(ctx.Ctx(), uid)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}
	friends := make([]int, 0)
	for _, item := range data {
		if item.DialogType == 1 {
			friends = append(friends, item.ReceiverId)
		}
	}

	//remarks, _ := c.contactService.Dao().Remarks(ctx.Ctx(), uid, friends)
	items := make([]*api_v1.ChatItem, 0)
	for _, item := range data {
		value := &api_v1.ChatItem{
			Id:         int32(item.Id),
			DialogType: int32(item.DialogType),
			ReceiverId: int32(item.ReceiverId),
			IsTop:      int32(item.IsTop),
			IsDisturb:  int32(item.IsDisturb),
			IsBot:      int32(item.IsBot),
			Avatar:     item.UserAvatar,
			MsgText:    "",
			UpdatedAt:  timeutil.FormatDatetime(item.UpdatedAt),
		}

		if num, ok := unReads[fmt.Sprintf("%d_%d", item.DialogType, item.ReceiverId)]; ok {
			value.UnreadNum = int32(num)
		}

		if item.DialogType == 1 {
			value.Username = item.Username
			value.Avatar = item.UserAvatar
			//if item.IsBot == 1 {
			//    bot, err := d.BotRepo.GetByUserId(ctx.Ctx(), 1)
			//    if err != nil {
			//
			//    }
			//    value.Name = bot.Name
			//} else {
			value.Name = item.Name
			//}
			value.Surname = item.Surname
			//value.Remark = remarks[item.ReceiverId]
			value.IsOnline = int32(strutil.BoolToInt(c.ClientStorage.IsOnline(ctx.Ctx(), entity.ImChannelChat, strconv.Itoa(int(value.ReceiverId)))))
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		if msg, err := c.MessageStorage.Get(ctx.Ctx(), item.DialogType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return ctx.Success(&api_v1.ChatListResponse{Items: items})
}

func (c *Chat) Delete(ctx *core.Context) error {
	params := &api_v1.ChatDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatService.Delete(ctx.Ctx(), ctx.UserId(), int(params.ListId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.ChatDeleteResponse{})
}

func (c *Chat) ClearUnreadMessage(ctx *core.Context) error {
	params := &api_v1.ChatClearUnreadNumRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	c.UnreadStorage.Reset(ctx.Ctx(), int(params.DialogType), int(params.ReceiverId), ctx.UserId())

	return ctx.Success(&api_v1.ChatClearUnreadNumResponse{})
}

func (c *Chat) Top(ctx *core.Context) error {
	params := &api_v1.ChatTopRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatService.Top(ctx.Ctx(), &service.ChatTopOpt{
		UserId: ctx.UserId(),
		Id:     int(params.ListId),
		Type:   int(params.Type),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.ChatTopResponse{})
}

func (c *Chat) Disturb(ctx *core.Context) error {
	params := &api_v1.ChatDisturbRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatService.Disturb(ctx.Ctx(), &service.ChatDisturbOpt{
		UserId:     ctx.UserId(),
		DialogType: int(params.DialogType),
		ReceiverId: int(params.ReceiverId),
		IsDisturb:  int(params.IsDisturb),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.ChatDisturbResponse{})
}
