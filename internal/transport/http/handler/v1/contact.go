package v1

import (
	"errors"
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type Contact struct {
	ContactService     *service.ContactService
	ClientStorage      *cache.ClientStorage
	ChatService        *service.ChatService
	MessageSendService service.MessageSendService
	ContactRepo        *repo.Contact
	UserRepo           *repo.User
	DialogRepo         *repo.Dialog
}

func (c *Contact) List(ctx *core.Context) error {
	list, err := c.ContactService.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
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
			Remark:   item.Remark,
		})
	}

	return ctx.Success(&v1Pb.ContactListResponse{Items: items})
}

func (c *Contact) Get(ctx *core.Context) error {
	params := &v1Pb.ContactDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	user, err := c.UserRepo.FindById(ctx.Ctx(), int(params.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.ErrorBusiness("Пользователь не существует")
		}

		return ctx.ErrorBusiness(err.Error())
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
		contact, err := c.ContactRepo.FindByWhere(ctx.Ctx(), "user_id = ? and friend_id = ?", uid, user.Id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err == nil && contact.Status == 1 {
			if c.ContactRepo.IsFriend(ctx.Ctx(), uid, user.Id, false) {
				data.FriendStatus = 2
				data.FolderId = int32(contact.FolderId)
				data.Remark = contact.Remark
			}
		}
	}

	return ctx.Success(&data)
}

func (c *Contact) Delete(ctx *core.Context) error {
	params := &v1Pb.ContactDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := c.ContactService.Delete(ctx.Ctx(), uid, int(params.FriendId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	_ = c.MessageSendService.SendSystemText(ctx.Ctx(), uid, &v1Pb.TextMessageRequest{
		Content: "Контакт удален",
		Receiver: &v1Pb.MessageReceiver{
			DialogType: entity.ChatPrivateMode,
			ReceiverId: params.FriendId,
		},
	})

	sid := c.DialogRepo.FindBySessionId(uid, int(params.FriendId), entity.ChatPrivateMode)
	if err := c.ChatService.Delete(ctx.Ctx(), ctx.UserId(), sid); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.ContactDeleteResponse{})
}

func (c *Contact) EditRemark(ctx *core.Context) error {
	params := &v1Pb.ContactRemarkEditRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ContactService.UpdateRemark(ctx.Ctx(), ctx.UserId(), int(params.FriendId), params.Remark); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.ContactRemarkEditResponse{})
}
