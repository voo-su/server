package v1

import (
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

type ContactRequest struct {
	Locale                locale.ILocale
	ContactRequestUseCase *usecase.ContactRequestUseCase
	ContactUseCase        *usecase.ContactUseCase
	MessageUseCase        usecase.IMessageUseCase
}

func (c *ContactRequest) ApplyUnreadNum(ctx *ginutil.Context) error {
	return ctx.Success(&v1Pb.ContactApplyUnreadNumResponse{
		UnreadNum: int64(c.ContactRequestUseCase.GetApplyUnreadNum(ctx.Ctx(), ctx.UserId())),
	})
}

func (c *ContactRequest) Create(ctx *ginutil.Context) error {
	params := &v1Pb.ContactRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if c.ContactUseCase.ContactRepo.IsFriend(ctx.Ctx(), uid, int(params.FriendId), false) {
		return ctx.Success(nil)
	}

	if err := c.ContactRequestUseCase.Create(ctx.Ctx(), &usecase.ContactApplyCreateOpt{
		UserId:   ctx.UserId(),
		FriendId: int(params.FriendId),
	}); err != nil {
		return ctx.Error(err)
	}

	return ctx.Success(&v1Pb.ContactRequestCreateResponse{})
}

func (c *ContactRequest) Accept(ctx *ginutil.Context) error {
	params := &v1Pb.ContactRequestAcceptRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	_, err := c.ContactRequestUseCase.Accept(ctx.Ctx(), &usecase.ContactApplyAcceptOpt{
		ApplyId: int(params.ApplyId),
		UserId:  uid,
	})
	if err != nil {
		return ctx.Error(err)
	}

	//if err := c.MessageUseCase.SendSystemText(ctx.Ctx(), applyInfo.UserId, &entity.TextMessageRequest{
	//	Content: c.Locale.Localize("can_start_communicating"),
	//	Receiver: &entity.MessageReceiver{
	//		ChatType: constant.ChatPrivateMode,
	//		ReceiverId: int32(applyInfo.FriendId),
	//	},
	//}); err != nil {
	//	log.Println(err)
	//}

	return ctx.Success(&v1Pb.ContactRequestAcceptResponse{})
}

func (c *ContactRequest) Decline(ctx *ginutil.Context) error {
	params := &v1Pb.ContactRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ContactRequestUseCase.Decline(ctx.Ctx(), &usecase.ContactApplyDeclineOpt{
		UserId:  ctx.UserId(),
		ApplyId: int(params.ApplyId),
	}); err != nil {
		return ctx.Error(err)
	}

	return ctx.Success(&v1Pb.ContactRequestDeclineResponse{})
}

func (c *ContactRequest) List(ctx *ginutil.Context) error {
	list, err := c.ContactRequestUseCase.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.Error(err)
	}

	items := make([]*v1Pb.ContactRequestListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &v1Pb.ContactRequestListResponse_Item{
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

	c.ContactRequestUseCase.ClearApplyUnreadNum(ctx.Ctx(), ctx.UserId())

	return ctx.Success(&v1Pb.ContactRequestListResponse{Items: items})
}
