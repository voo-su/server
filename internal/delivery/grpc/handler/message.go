package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	messagePb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/strutil"
)

type Message struct {
	messagePb.UnimplementedMessageServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	MessageUseCase usecase.IMessageUseCase
}

func NewMessageHandler(
	conf *config.Config,
	locale locale.ILocale,
	messageUseCase usecase.IMessageUseCase,
) *Message {
	return &Message{
		Conf:           conf,
		Locale:         locale,
		MessageUseCase: messageUseCase,
	}
}

func (m *Message) GetHistory(ctx context.Context, in *messagePb.GetHistoryRequest) (*messagePb.GetHistoryResponse, error) {
	uid := grpcutil.UserId(ctx)

	if in.ChatType == constant.ChatGroupMode {
		if err := m.MessageUseCase.IsAccess(ctx, &entity.MessageAccess{
			ChatType:   int(in.ChatType),
			UserId:     uid,
			ReceiverId: int(in.ReceiverId),
		}); err != nil {
			items := make([]*messagePb.MessageItem, 0)

			items = append(items, &messagePb.MessageItem{
				Id:         strutil.NewMsgId(),
				ChatType:   int32(in.ChatType),
				MsgType:    constant.ChatMsgSysText,
				ReceiverId: in.ReceiverId,
				Content:    m.Locale.Localize("insufficient_permissions_to_view_messages"),
			})

			return &messagePb.GetHistoryResponse{
				Limit:    in.Limit,
				RecordId: 0,
				Items:    items,
			}, nil
		}
	}

	records, err := m.MessageUseCase.GetHistory(ctx, &entity.QueryGetHistoryOpt{
		ChatType:   int(in.ChatType),
		UserId:     uid,
		ReceiverId: int(in.ReceiverId),
		RecordId:   int(in.RecordId),
		Limit:      int(in.Limit),
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*messagePb.MessageItem, 0)
	for _, item := range records {
		items = append(items, &messagePb.MessageItem{
			Id:         item.MsgId,
			ChatType:   int32(item.ChatType),
			MsgType:    int32(item.MsgType),
			ReceiverId: int64(item.ReceiverId),
			UserId:     int64(item.UserId),
			Content:    item.Content,
			IsRead:     item.IsRead == 1,
			CreatedAt:  item.CreatedAt,
		})
	}

	rid := 0
	if length := len(records); length > 0 {
		rid = records[length-1].Sequence
	}

	return &messagePb.GetHistoryResponse{
		Limit:    in.Limit,
		RecordId: int64(rid),
		Items:    items,
	}, nil
}

func (m *Message) Send(ctx context.Context, in *messagePb.SendMessageRequest) (*messagePb.SendMessageResponse, error) {
	uid := grpcutil.UserId(ctx)

	if err := m.MessageUseCase.SendText(ctx, uid, &entity.SendText{
		Receiver: entity.MessageReceiver{
			ChatType:   int32(in.ChatType),
			ReceiverId: int32(in.ReceiverId),
		},
		Content: in.Message,
		//QuoteId: ,
	}); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &messagePb.SendMessageResponse{}, nil
}
