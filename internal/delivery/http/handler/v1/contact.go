package v1

import (
	"errors"
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

type Contact struct {
	Locale         locale.ILocale
	ContactUseCase *usecase.ContactUseCase
	ChatUseCase    *usecase.ChatUseCase
	UserUseCase    *usecase.UserUseCase
	MessageUseCase usecase.IMessageUseCase
}

func (c *Contact) List(ctx *ginutil.Context) error {
	list, err := c.ContactUseCase.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.Error(err.Error())
	}

	items := make([]*v1Pb.ContactListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &v1Pb.ContactListResponse_Item{
			Id:       int32(item.Id),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
			Gender:   int32(item.Gender),
			About:    item.About,
			FolderId: int32(item.FolderId),
		})
	}

	return ctx.Success(&v1Pb.ContactListResponse{Items: items})
}

func (c *Contact) Get(ctx *ginutil.Context) error {
	params := &v1Pb.ContactDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	user, err := c.UserUseCase.UserRepo.FindById(ctx.Ctx(), int(params.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Error(c.Locale.Localize("user_not_found"))
		}

		return ctx.Error(err.Error())
	}

	data := v1Pb.ContactDetailResponse{
		Id:           int32(user.Id),
		Username:     user.Username,
		Avatar:       user.Avatar,
		Name:         user.Name,
		Surname:      user.Surname,
		Gender:       int32(user.Gender),
		About:        user.About,
		FriendStatus: 0,
		IsBot:        int32(user.IsBot),
	}
	if uid != user.Id {
		data.FriendStatus = 1
		contact, err := c.ContactUseCase.ContactRepo.FindByWhere(ctx.Ctx(), "user_id = ? AND friend_id = ?", uid, user.Id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err == nil && contact.Status == 1 {
			if c.ContactUseCase.ContactRepo.IsFriend(ctx.Ctx(), uid, user.Id, false) {
				data.FriendStatus = 2
				data.FolderId = int32(contact.FolderId)
			}
		}
	}

	return ctx.Success(&data)
}

func (c *Contact) Delete(ctx *ginutil.Context) error {
	params := &v1Pb.ContactDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := c.ContactUseCase.Delete(ctx.Ctx(), uid, int(params.FriendId)); err != nil {
		return ctx.Error(err.Error())
	}

	_ = c.MessageUseCase.SendSystemText(ctx.Ctx(), uid, &entity.TextMessageRequest{
		Content: c.Locale.Localize("contact_deleted"),
		Receiver: &entity.MessageReceiver{
			ChatType:   constant.ChatPrivateMode,
			ReceiverId: params.FriendId,
		},
	})

	sid := c.ChatUseCase.ChatRepo.FindBySessionId(uid, int(params.FriendId), constant.ChatPrivateMode)
	if err := c.ChatUseCase.Delete(ctx.Ctx(), ctx.UserId(), sid); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ContactDeleteResponse{})
}
