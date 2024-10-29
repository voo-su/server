package v1

import (
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/timeutil"
)

type GroupChatAds struct {
	GroupAdsService     *service.GroupChatAdsService
	GroupMemberService  *service.GroupChatMemberService
	MessageSendService  service.MessageSendService
	GroupChatMemberRepo *repo.GroupChatMember
	GroupChatAdsRepo    *repo.GroupChatAds
}

func (g *GroupChatAds) CreateAndUpdate(ctx *core.Context) error {
	params := &v1Pb.GroupChatAdsEditRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if !g.GroupChatMemberRepo.IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("У вас нет прав для выполнения этой операции")
	}

	var (
		msg string
		err error
	)
	if params.AdsId == 0 {
		err = g.GroupAdsService.Create(ctx.Ctx(), &service.GroupChatAdsEditOpt{
			UserId:    uid,
			GroupId:   int(params.GroupId),
			AdsId:     int(params.AdsId),
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     int(params.IsTop),
			IsConfirm: int(params.IsConfirm),
		})
		msg = "Успешно добавлено объявление в группу"
	} else {
		err = g.GroupAdsService.Update(ctx.Ctx(), &service.GroupChatAdsEditOpt{
			GroupId:   int(params.GroupId),
			AdsId:     int(params.AdsId),
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     int(params.IsTop),
			IsConfirm: int(params.IsConfirm),
		})
		msg = "Успешно обновлено объявление группы"
	}
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	_ = g.MessageSendService.SendSysOther(ctx.Ctx(), &model.Message{
		DialogType: model.DialogRecordDialogTypeGroup,
		MsgType:    entity.ChatMsgSysGroupAds,
		UserId:     uid,
		ReceiverId: int(params.GroupId),
		Extra: jsonutil.Encode(model.DialogRecordExtraGroupAds{
			OwnerId:   uid,
			OwnerName: "magomedcoder",
			Title:     params.Title,
			Content:   params.Content,
		}),
	})

	return ctx.Success(nil, msg)
}

func (g *GroupChatAds) Delete(ctx *core.Context) error {
	params := &v1Pb.GroupChatAdsDeleteRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := g.GroupAdsService.Delete(ctx.Ctx(), int(params.GroupId), int(params.AdsId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}
	return ctx.Success(nil, "Успешно удалено объявление группы")
}

func (g *GroupChatAds) List(ctx *core.Context) error {
	params := &v1Pb.GroupChatAdsListRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !g.GroupChatMemberRepo.IsMember(ctx.Ctx(), int(params.GroupId), ctx.UserId(), true) {
		return ctx.ErrorBusiness("У вас нет прав на получение данных")
	}

	all, err := g.GroupChatAdsRepo.GetListAll(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*v1Pb.GroupChatAdsListResponse_Item, 0)
	for i := 0; i < len(all); i++ {
		items = append(items, &v1Pb.GroupChatAdsListResponse_Item{
			Id:           int32(all[i].Id),
			Title:        all[i].Title,
			Content:      all[i].Content,
			IsTop:        int32(all[i].IsTop),
			IsConfirm:    int32(all[i].IsConfirm),
			ConfirmUsers: all[i].ConfirmUsers,
			Avatar:       all[i].Avatar,
			CreatorId:    int32(all[i].CreatorId),
			CreatedAt:    timeutil.FormatDatetime(all[i].CreatedAt),
			UpdatedAt:    timeutil.FormatDatetime(all[i].UpdatedAt),
		})
	}

	return ctx.Success(&v1Pb.GroupChatAdsListResponse{
		Items: items,
	})
}
