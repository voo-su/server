package v1

import (
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/repository/repo"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
)

type ContactRequest struct {
	ContactRequestUseCase *usecase.ContactRequestUseCase
	ContactUseCase        *usecase.ContactUseCase
	MessageSendUseCase    usecase.MessageSendUseCase
	ContactRepo           *repo.Contact
}

func (ca *ContactRequest) ApplyUnreadNum(ctx *core.Context) error {
	return ctx.Success(map[string]any{
		"unread_num": ca.ContactRequestUseCase.GetApplyUnreadNum(ctx.Ctx(), ctx.UserId()),
	})
}

func (ca *ContactRequest) Create(ctx *core.Context) error {
	params := &v1Pb.ContactRequestCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if ca.ContactRepo.IsFriend(ctx.Ctx(), uid, int(params.FriendId), false) {
		return ctx.Success(nil)
	}

	if err := ca.ContactRequestUseCase.Create(ctx.Ctx(), &usecase.ContactApplyCreateOpt{
		UserId:   ctx.UserId(),
		FriendId: int(params.FriendId),
	}); err != nil {
		return ctx.ErrorBusiness(err)
	}

	return ctx.Success(&v1Pb.ContactRequestCreateResponse{})
}

func (ca *ContactRequest) Accept(ctx *core.Context) error {
	params := &v1Pb.ContactRequestAcceptRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	_, err := ca.ContactRequestUseCase.Accept(ctx.Ctx(), &usecase.ContactApplyAcceptOpt{
		ApplyId: int(params.ApplyId),
		UserId:  uid,
	})
	if err != nil {
		return ctx.ErrorBusiness(err)
	}
	//err = ca.MessageUseCase.SendSystemText(ctx.Ctx(), applyInfo.UserId, &v1Pb.TextMessageRequest{
	//	Content: "Теперь можете начать общаться",
	//	Receiver: &v1Pb.MessageReceiver{
	//		DialogType: domain.ChatPrivateMode,
	//		ReceiverId: int32(applyInfo.FriendId),
	//	},
	//})
	//if err != nil {
	//	fmt.Println("ошибка", err.Error())
	//}
	return ctx.Success(&v1Pb.ContactRequestAcceptResponse{})
}

func (ca *ContactRequest) Decline(ctx *core.Context) error {
	params := &v1Pb.ContactRequestDeclineRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := ca.ContactRequestUseCase.Decline(ctx.Ctx(), &usecase.ContactApplyDeclineOpt{
		UserId:  ctx.UserId(),
		ApplyId: int(params.ApplyId),
	}); err != nil {
		return ctx.ErrorBusiness(err)
	}

	return ctx.Success(&v1Pb.ContactRequestDeclineResponse{})
}

func (ca *ContactRequest) List(ctx *core.Context) error {
	list, err := ca.ContactRequestUseCase.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.Error(err.Error())
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

	ca.ContactRequestUseCase.ClearApplyUnreadNum(ctx.Ctx(), ctx.UserId())

	return ctx.Success(&v1Pb.ContactRequestListResponse{Items: items})
}
