package v1

import (
	"errors"
	"gorm.io/gorm"
	"voo.su/api/pb/v1"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type Contact struct {
	ContactService     *service.ContactService
	ClientStorage      *cache.ClientStorage
	DialogService      *service.DialogService
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

	items := make([]*api_v1.ContactListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &api_v1.ContactListResponse_Item{
			Id:       int32(item.Id),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
			Gender:   int32(item.Gender),
			About:    item.About,
			//Remark:   item.Remark,
			GroupId: int32(item.GroupId),
		})
	}

	return ctx.Success(&api_v1.ContactListResponse{Items: items})
}

func (c *Contact) Get(ctx *core.Context) error {
	params := &api_v1.ContactDetailRequest{}
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

	data := api_v1.ContactDetailResponse{
		Id:       int32(user.Id),
		Username: user.Username,
		Avatar:   user.Avatar,
		Name:     user.Name,
		Surname:  user.Surname,
		Gender:   int32(user.Gender),
		About:    user.About,
		//FriendApply:  0,
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
				data.GroupId = int32(contact.FolderId)
				//data.Remark = contact.Remark
			}
		}
	}

	return ctx.Success(&data)
}
