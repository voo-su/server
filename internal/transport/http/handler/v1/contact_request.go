package v1

import (
	"voo.su/api/pb/v1"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type ContactRequest struct {
	ContactRequestService *service.ContactRequestService
	ContactService        *service.ContactService
	MessageSendService    service.MessageSendService
	ContactRepo           *repo.Contact
}

func (ca *ContactRequest) ApplyUnreadNum(ctx *core.Context) error {
	return ctx.Success(map[string]any{
		"unread_num": ca.ContactRequestService.GetApplyUnreadNum(ctx.Ctx(), ctx.UserId()),
	})
}

func (ca *ContactRequest) Create(ctx *core.Context) error {
	params := &api_v1.ContactRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if ca.ContactRepo.IsFriend(ctx.Ctx(), uid, int(params.FriendId), false) {
		return ctx.Success(nil)
	}

	if err := ca.ContactRequestService.Create(ctx.Ctx(), &service.ContactApplyCreateOpt{
		UserId:   ctx.UserId(),
		FriendId: int(params.FriendId),
	}); err != nil {
		return ctx.ErrorBusiness(err)
	}

	return ctx.Success(&api_v1.ContactRequestCreateResponse{})
}

func (ca *ContactRequest) Accept(ctx *core.Context) error {
	params := &api_v1.ContactRequestAcceptRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	_, err := ca.ContactRequestService.Accept(ctx.Ctx(), &service.ContactApplyAcceptOpt{
		ApplyId: int(params.ApplyId),
		UserId:  uid,
	})
	if err != nil {
		return ctx.ErrorBusiness(err)
	}
	//err = ca.MessageService.SendSystemText(ctx.Ctx(), applyInfo.UserId, &api_v1.TextMessageRequest{
	//	Content: "Теперь можете начать общаться",
	//	Receiver: &api_v1.MessageReceiver{
	//		DialogType: entity.ChatPrivateMode,
	//		ReceiverId: int32(applyInfo.FriendId),
	//	},
	//})
	//if err != nil {
	//	fmt.Println("ошибка", err.Error())
	//}
	return ctx.Success(&api_v1.ContactRequestAcceptResponse{})
}

func (ca *ContactRequest) Decline(ctx *core.Context) error {
	params := &api_v1.ContactRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := ca.ContactRequestService.Decline(ctx.Ctx(), &service.ContactApplyDeclineOpt{
		UserId:  ctx.UserId(),
		ApplyId: int(params.ApplyId),
	}); err != nil {
		return ctx.ErrorBusiness(err)
	}

	return ctx.Success(&api_v1.ContactRequestDeclineResponse{})
}

func (ca *ContactRequest) List(ctx *core.Context) error {
	list, err := ca.ContactRequestService.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.Error(err.Error())
	}

	items := make([]*api_v1.ContactRequestListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &api_v1.ContactRequestListResponse_Item{
			Id:       int32(item.Id),
			UserId:   int32(item.UserId),
			FriendId: int32(item.FriendId),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
			//CreatedAt: timeutil.FormatDatetime(item.CreatedAt),
		})
	}

	ca.ContactRequestService.ClearApplyUnreadNum(ctx.Ctx(), ctx.UserId())

	return ctx.Success(&api_v1.ContactRequestListResponse{Items: items})
}
