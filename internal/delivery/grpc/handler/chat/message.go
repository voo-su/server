package chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/strutil"
)

func (c *Chat) GetHistory(ctx context.Context, in *chatPb.GetHistoryRequest) (*chatPb.GetHistoryResponse, error) {
	uid := grpcutil.UserId(ctx)
	receiver := in.Receiver
	fmt.Println(receiver, uid)
	if receiver.ChatType == constant.ChatGroupMode {
		if err := c.MessageUseCase.IsAccess(ctx, &entity.MessageAccess{
			ChatType:   int(receiver.ChatType),
			ReceiverId: int(receiver.ReceiverId),
			UserId:     uid,
		}); err != nil {
			items := make([]*chatPb.MessageItem, 0)

			items = append(items, &chatPb.MessageItem{
				Id: strutil.NewMsgId(),
				Receiver: &chatPb.Receiver{
					ChatType:   receiver.ChatType,
					ReceiverId: receiver.ReceiverId,
				},
				MsgType: constant.ChatMsgSysText,
				Content: c.Locale.Localize("insufficient_permissions_to_view_messages"),
			})

			return &chatPb.GetHistoryResponse{
				Limit:    in.Limit,
				RecordId: 0,
				Items:    items,
			}, nil
		}
	}

	records, err := c.MessageUseCase.GetHistory(ctx, &entity.QueryGetHistoryOpt{
		ChatType:   int(receiver.ChatType),
		ReceiverId: int(receiver.ReceiverId),
		UserId:     uid,
		RecordId:   int(in.RecordId),
		Limit:      int(in.Limit),
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*chatPb.MessageItem, 0)
	for _, item := range records {
		items = append(items, &chatPb.MessageItem{
			Id: item.MsgId,
			Receiver: &chatPb.Receiver{
				ChatType:   int32(item.ChatType),
				ReceiverId: int64(item.ReceiverId),
			},
			MsgType:   int32(item.MsgType),
			UserId:    int64(item.UserId),
			Content:   item.Content,
			IsRead:    item.IsRead == 1,
			CreatedAt: item.CreatedAt,
		})
	}

	rid := 0
	if length := len(records); length > 0 {
		rid = records[length-1].Sequence
	}

	return &chatPb.GetHistoryResponse{
		Limit:    in.Limit,
		RecordId: int64(rid),
		Items:    items,
	}, nil
}

func (c *Chat) SendMessage(ctx context.Context, in *chatPb.SendMessageRequest) (*chatPb.SendMessageResponse, error) {
	uid := grpcutil.UserId(ctx)

	if err := c.MessageUseCase.SendText(ctx, uid, &entity.SendText{
		Receiver: entity.MessageReceiver{
			ChatType:   int32(in.Receiver.ChatType),
			ReceiverId: int32(in.Receiver.ReceiverId),
		},
		Content: in.Message,
		//QuoteId: ,
	}); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &chatPb.SendMessageResponse{}, nil
}

//func (c *Chat) SendPhoto(ctx context.Context, in *chatPb.SendPhotoRequest) (*chatPb.SendPhotoResponse, error) {
//	// TODO
//	return &chatPb.DeleteMessagesResponse{}, nil
//}

func (c *Chat) ViewMessages(ctx context.Context, in *chatPb.ViewMessagesRequest) (*chatPb.ViewMessagesResponse, error) {
	uid := grpcutil.UserId(ctx)

	// TODO

	c.ChatUseCase.UnreadCacheRepo.Reset(ctx, int(in.Receiver.ChatType), int(in.Receiver.ReceiverId), uid)

	return &chatPb.ViewMessagesResponse{}, nil
}

func (c *Chat) DeleteMessages(ctx context.Context, in *chatPb.DeleteMessagesRequest) (*chatPb.DeleteMessagesResponse, error) {
	uid := grpcutil.UserId(ctx)
	fmt.Println(uid)
	fmt.Println(in.Receiver.ChatType)
	fmt.Println(in.Receiver.ReceiverId)
	fmt.Println(in.MessageIds)
	fmt.Println(in.Revoke)
	// TODO
	return &chatPb.DeleteMessagesResponse{}, nil
}
