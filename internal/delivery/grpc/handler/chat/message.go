package chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/jsonutil"
)

func (c *Chat) GetHistory(ctx context.Context, in *chatPb.GetHistoryRequest) (*chatPb.GetHistoryResponse, error) {
	uid := grpcutil.UserId(ctx)
	receiver := in.Receiver
	if receiver.ChatType == constant.ChatGroupMode {
		if err := c.MessageUseCase.IsAccess(ctx, &entity.MessageAccess{
			ChatType:   int(receiver.ChatType),
			ReceiverId: int(receiver.ReceiverId),
			UserId:     uid,
		}); err != nil {
			items := make([]*chatPb.MessageItem, 0)

			items = append(items, &chatPb.MessageItem{
				//Id: strutil.NewMsgId(),
				Receiver: &chatPb.Receiver{
					ChatType:   receiver.ChatType,
					ReceiverId: receiver.ReceiverId,
				},
				MsgType: constant.ChatMsgSysText,
				Content: c.Locale.Localize("insufficient_permissions_to_view_messages"),
			})

			return &chatPb.GetHistoryResponse{
				Limit:     in.Limit,
				MessageId: 0,
				Items:     items,
			}, nil
		}
	}

	messages, err := c.MessageUseCase.GetHistory(ctx, &entity.QueryGetHistoryOpt{
		ChatType:   int(receiver.ChatType),
		ReceiverId: int(receiver.ReceiverId),
		UserId:     uid,
		MessageId:  int(in.MessageId),
		Limit:      int(in.Limit),
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*chatPb.MessageItem, 0)
	for _, item := range messages {
		messageItem := &chatPb.MessageItem{
			Id: int64(item.Id),
			Receiver: &chatPb.Receiver{
				ChatType:   int32(item.ChatType),
				ReceiverId: int64(item.ReceiverId),
			},
			MsgType:   int32(item.MsgType),
			UserId:    int64(item.UserId),
			Content:   item.Content,
			IsRead:    item.IsRead == 1,
			CreatedAt: item.CreatedAt,
		}

		if item.MsgType == constant.ChatMsgTypeImage {
			var file entity.MessageExtraImage
			if err := jsonutil.Decode(item.Extra0, &file); err != nil {
				fmt.Println(err)
			}

			messageItem.Media = &chatPb.MessageMedia{
				Media: &chatPb.MessageMedia_MessageMediaPhoto{
					MessageMediaPhoto: &chatPb.MessageMediaPhoto{
						File: file.Url,
					},
				},
			}
		}

		if item.MsgType == constant.ChatMsgTypeVideo {
			var file entity.MessageExtraVideo
			if err := jsonutil.Decode(item.Extra0, &file); err != nil {
				fmt.Println(err)
			}

			messageItem.Media = &chatPb.MessageMedia{
				Media: &chatPb.MessageMedia_MessageMediaDocument{
					MessageMediaDocument: &chatPb.MessageMediaDocument{
						File:     file.Url,
						MimeType: "video/" + file.Suffix,
					},
				},
			}
		}

		if item.MsgType == constant.ChatMsgTypeAudio {
			var file entity.MessageExtraAudio
			if err := jsonutil.Decode(item.Extra0, &file); err != nil {
				fmt.Println(err)
			}

			messageItem.Media = &chatPb.MessageMedia{
				Media: &chatPb.MessageMedia_MessageMediaDocument{
					MessageMediaDocument: &chatPb.MessageMediaDocument{
						File:     file.Url,
						MimeType: "audio/" + file.Suffix,
					},
				},
			}
		}

		if item.QuoteId != "" {
			type ReplyData struct {
				UserId   int    `json:"user_id,omitempty"`
				Username string `json:"username,omitempty"`
				MsgType  int    `json:"msg_type,omitempty"`
				Content  string `json:"content,omitempty"`
				MsgId    string `json:"msg_id,omitempty"`
			}
			type Reply struct {
				Reply ReplyData `json:"reply,omitempty"`
			}

			var reply Reply
			if err := jsonutil.Decode(item.Extra0, &reply); err != nil {
				fmt.Println(err)
			}

			mId, err := c.MessageUseCase.GetMessageByMsgId(ctx, item.QuoteId)
			if err != nil {
				fmt.Println(err)
			}

			messageItem.Reply = &chatPb.MessageReply{
				Id:       int64(mId),
				MsgType:  int32(reply.Reply.MsgType),
				UserId:   int64(reply.Reply.UserId),
				Username: reply.Reply.MsgId,
				Content:  reply.Reply.MsgId,
			}
		}

		items = append(items, messageItem)
	}

	rid := 0
	if length := len(messages); length > 0 {
		rid = messages[length-1].Sequence
	}

	return &chatPb.GetHistoryResponse{
		Limit:     in.Limit,
		MessageId: int64(rid),
		Items:     items,
	}, nil
}

func (c *Chat) SendMessage(ctx context.Context, in *chatPb.SendMessageRequest) (*chatPb.SendMessageResponse, error) {
	uid := grpcutil.UserId(ctx)
	if err := c.MessageUseCase.SendText(ctx, uid, &entity.SendText{
		Receiver: entity.MessageReceiver{
			ChatType:   in.Receiver.ChatType,
			ReceiverId: int32(in.Receiver.ReceiverId),
		},
		Content:      in.Message,
		ReplyToMsgId: in.ReplyToMsgId,
	}); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &chatPb.SendMessageResponse{
		Success: true,
	}, nil
}

func (c *Chat) SendPhoto(ctx context.Context, in *chatPb.SendPhotoRequest) (*chatPb.SendPhotoResponse, error) {
	// TODO
	return &chatPb.SendPhotoResponse{}, nil
}

func (c *Chat) ViewMessages(ctx context.Context, in *chatPb.ViewMessagesRequest) (*chatPb.ViewMessagesResponse, error) {
	uid := grpcutil.UserId(ctx)

	// TODO

	c.ChatUseCase.UnreadCacheRepo.Reset(ctx, int(in.Receiver.ChatType), int(in.Receiver.ReceiverId), uid)

	return &chatPb.ViewMessagesResponse{}, nil
}

func (c *Chat) DeleteMessages(ctx context.Context, in *chatPb.DeleteMessagesRequest) (*chatPb.DeleteMessagesResponse, error) {
	uid := grpcutil.UserId(ctx)
	if err := c.ChatUseCase.DeleteRecordList(ctx, &usecase.RemoveRecordListOpt{
		UserId:     uid,
		ChatType:   int(in.Receiver.ChatType),
		ReceiverId: int(in.Receiver.ReceiverId),
		MsgIds:     in.MessageIds,
	}); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &chatPb.DeleteMessagesResponse{
		Success: true,
	}, nil
}
