package v1

import (
	"fmt"
	"strconv"
	"strings"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type Chat struct {
	Locale                 locale.ILocale
	ChatUseCase            *usecase.ChatUseCase
	ContactUseCase         *usecase.ContactUseCase
	GroupChatUseCase       *usecase.GroupChatUseCase
	AuthUseCase            *usecase.AuthUseCase
	UserUseCase            *usecase.UserUseCase
	RedisLockRepo          *redisRepo.RedisLockCacheRepository
	ClientCacheRepo        *redisRepo.ClientCacheRepository
	MessageCacheRepo       *redisRepo.MessageCacheRepository
	UnreadCacheRepo        *redisRepo.UnreadCacheRepository
	ContactRemarkCacheRepo *redisRepo.ContactRemarkCacheRepository
}

func (c *Chat) Create(ctx *ginutil.Context) error {
	var (
		params = &v1Pb.ChatCreateRequest{}
		uid    = ctx.UserId()
		agent  = strings.TrimSpace(ctx.Context.GetHeader("user-agent"))
	)
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if agent != "" {
		agent = encrypt.Md5(agent)
	}

	if params.DialogType == constant.ChatPrivateMode && int(params.ReceiverId) == ctx.UserId() {
		return ctx.Error(c.Locale.Localize("creation_error"))
	}

	key := fmt.Sprintf("dialog:list:%d-%d-%d-%s", uid, params.ReceiverId, params.DialogType, agent)
	if !c.RedisLockRepo.Lock(ctx.Ctx(), key, 10) {
		return ctx.Error(c.Locale.Localize("creation_error"))
	}

	if c.AuthUseCase.IsAuth(ctx.Ctx(), &usecase.AuthOption{
		DialogType: int(params.DialogType),
		UserId:     uid,
		ReceiverId: int(params.ReceiverId),
	}) != nil {
		return ctx.Error(c.Locale.Localize("insufficient_permissions"))
	}

	result, err := c.ChatUseCase.Create(ctx.Ctx(), &usecase.CreateChatOpt{
		UserId:     uid,
		DialogType: int(params.DialogType),
		ReceiverId: int(params.ReceiverId),
	})
	if err != nil {
		return ctx.Error(err.Error())
	}

	item := &v1Pb.ChatItem{
		Id:         int32(result.Id),
		DialogType: int32(result.DialogType),
		ReceiverId: int32(result.ReceiverId),
		IsBot:      int32(result.IsBot),
		UpdatedAt:  timeutil.DateTime(),
	}
	if item.DialogType == constant.ChatPrivateMode {
		item.UnreadNum = int32(c.UnreadCacheRepo.Get(ctx.Ctx(), 1, int(params.ReceiverId), uid))
		item.Remark = c.ContactUseCase.ContactRepo.GetFriendRemark(ctx.Ctx(), uid, int(params.ReceiverId))
		if user, err := c.UserUseCase.UserRepo.FindById(ctx.Ctx(), result.ReceiverId); err == nil {
			item.Username = user.Username
			item.Name = user.Name
			item.Surname = user.Surname
			item.Avatar = user.Avatar
		}
	} else if result.DialogType == constant.ChatGroupMode {
		if group, err := c.GroupChatUseCase.GroupChatRepo.FindById(ctx.Ctx(), int(params.ReceiverId)); err == nil {
			item.Name = group.Name
		}
	}

	if msg, err := c.MessageCacheRepo.Get(ctx.Ctx(), result.DialogType, uid, result.ReceiverId); err == nil {
		item.MsgText = msg.Content
		item.UpdatedAt = msg.Datetime
	}

	return ctx.Success(&v1Pb.ChatCreateResponse{
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
		RemarkName: item.Remark,
		UnreadNum:  item.UnreadNum,
		MsgText:    item.MsgText,
		UpdatedAt:  item.UpdatedAt,
	})
}

func (c *Chat) List(ctx *ginutil.Context) error {
	uid := ctx.UserId()
	unReads := c.UnreadCacheRepo.All(ctx.Ctx(), uid)
	if len(unReads) > 0 {
		c.ChatUseCase.BatchAddList(ctx.Ctx(), uid, unReads)
	}

	data, err := c.ChatUseCase.List(ctx.Ctx(), uid)
	if err != nil {
		return ctx.Error(err.Error())
	}
	friends := make([]int, 0)
	for _, item := range data {
		if item.DialogType == 1 {
			friends = append(friends, item.ReceiverId)
		}
	}

	remarks, _ := c.ContactUseCase.ContactRepo.Remarks(ctx.Ctx(), uid, friends)
	items := make([]*v1Pb.ChatItem, 0)
	for _, item := range data {
		value := &v1Pb.ChatItem{
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
			value.Remark = remarks[item.ReceiverId]
			value.IsOnline = int32(strutil.BoolToInt(c.ClientCacheRepo.IsOnline(ctx.Ctx(), constant.ImChannelChat, strconv.Itoa(int(value.ReceiverId)))))
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		if msg, err := c.MessageCacheRepo.Get(ctx.Ctx(), item.DialogType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return ctx.Success(&v1Pb.ChatListResponse{Items: items})
}

func (c *Chat) Delete(ctx *ginutil.Context) error {
	params := &v1Pb.ChatDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatUseCase.Delete(ctx.Ctx(), ctx.UserId(), int(params.ListId)); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ChatDeleteResponse{})
}

func (c *Chat) ClearUnreadMessage(ctx *ginutil.Context) error {
	params := &v1Pb.ChatClearUnreadNumRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	c.UnreadCacheRepo.Reset(ctx.Ctx(), int(params.DialogType), int(params.ReceiverId), ctx.UserId())

	return ctx.Success(&v1Pb.ChatClearUnreadNumResponse{})
}

func (c *Chat) Top(ctx *ginutil.Context) error {
	params := &v1Pb.ChatTopRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatUseCase.Top(ctx.Ctx(), &usecase.ChatTopOpt{
		UserId: ctx.UserId(),
		Id:     int(params.ListId),
		Type:   int(params.Type),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ChatTopResponse{})
}

func (c *Chat) Disturb(ctx *ginutil.Context) error {
	params := &v1Pb.ChatDisturbRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ChatUseCase.Disturb(ctx.Ctx(), &usecase.ChatDisturbOpt{
		UserId:     ctx.UserId(),
		DialogType: int(params.DialogType),
		ReceiverId: int(params.ReceiverId),
		IsDisturb:  int(params.IsDisturb),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ChatDisturbResponse{})
}
