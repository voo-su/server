package chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	chatPb "voo.su/api/grpc/pb"
	commonPb "voo.su/api/grpc/pb/common"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/strutil"
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

			var _id string
			if item.FileId != nil {
				_id = item.FileId.String()
			}

			messageItem.Media = &chatPb.MessageMedia{
				Media: &chatPb.MessageMedia_MessageMediaPhoto{
					MessageMediaPhoto: &chatPb.MessageMediaPhoto{
						Id:   _id,
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
			var _id string
			if item.FileId != nil {
				_id = item.FileId.String()
			}
			messageItem.Media = &chatPb.MessageMedia{
				Media: &chatPb.MessageMedia_MessageMediaDocument{
					MessageMediaDocument: &chatPb.MessageMediaDocument{
						Id:       _id,
						MimeType: "video/" + file.Suffix,
						Size:     int32(file.Size),
						File:     file.Url,
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
						Attribute: &chatPb.DocumentAttribute{
							Attributes: &chatPb.DocumentAttribute_Audio{
								Audio: &commonPb.DocumentAttributeAudio{
									Duration:  int32(file.Duration),
									Title:     file.Title,
									Performer: file.Performer,
									Voice:     file.Voice,
								},
							},
						},
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
				Username: reply.Reply.Username,
				Content:  reply.Reply.Content,
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

func (c *Chat) SendMedia(ctx context.Context, in *chatPb.SendMediaRequest) (*chatPb.SendMediaResponse, error) {
	uid := grpcutil.UserId(ctx)

	receiver := entity.MessageReceiver{
		ChatType:   in.Receiver.ChatType,
		ReceiverId: int32(in.Receiver.ReceiverId),
	}

	switch media := in.Media.Media.(type) {
	case *chatPb.InputMedia_Photo:
		file := media.Photo.GetFile()
		fileExt := strutil.ExtractFileExtension(file.GetName())

		filePath, err := c.UploadUseCase.AssembleFileParts(ctx, uid, file.GetId(), file.GetParts(), file.GetName(), fileExt)
		if err != nil {
			log.Printf("ошибка сборки файла: %v", err)
			return nil, status.Error(codes.Unknown, "ошибка сборки файла")
		}

		if err := c.MessageUseCase.SendImage(ctx, uid, &entity.SendImage{
			Receiver: receiver,
			Url:      c.UploadUseCase.Minio.PublicUrl(c.Conf.Minio.GetBucket(), filePath.FilePath),
			//Width:   params.Width,
			//Height:  params.Height,
			ReplyToMsgId: in.ReplyToMsgId,
			Content:      in.Message,
			FileId:       &filePath.FileId,
		}); err != nil {
			log.Println(err)
			return nil, status.Error(codes.Unknown, "не удалось")
		}

	case *chatPb.InputMedia_Document:
		file := media.Document.GetFile()
		mimeType := media.Document.GetMimeType()
		attributes := media.Document.GetAttributes()
		fileExt := strutil.ExtractFileExtension(file.GetName())

		filePath, err := c.UploadUseCase.AssembleFileParts(ctx, uid, file.GetId(), file.GetParts(), file.GetName(), fileExt)
		if err != nil {
			log.Printf("ошибка сборки файла: %v", err)
			return nil, status.Error(codes.Unknown, "ошибка сборки файла")
		}

		switch {
		case strings.HasPrefix(mimeType, "video/"):
			if err := c.MessageUseCase.SendVideo(ctx, uid, &entity.SendVideo{
				Receiver: receiver,
				Url:      c.UploadUseCase.Minio.PublicUrl(c.Conf.Minio.GetBucket(), filePath.FilePath),
				Duration: attributes.GetVideo().GetDuration(),
				Size:     int32(filePath.Size),
				//Cover:    params.Cover,
				Content: in.Message,
				FileId:  &filePath.FileId,
			}); err != nil {
				log.Println(err)
				return nil, status.Error(codes.Unknown, "не удалось")
			}
		case strings.HasPrefix(mimeType, "audio/"):
			if err := c.MessageUseCase.SendAudio(ctx, uid, &entity.SendAudio{
				Receiver:  receiver,
				Url:       c.UploadUseCase.Minio.PublicUrl(c.Conf.Minio.GetBucket(), filePath.FilePath),
				Size:      int32(filePath.Size),
				Content:   in.Message,
				FileId:    &filePath.FileId,
				Voice:     attributes.GetAudio().GetVoice(),
				Title:     attributes.GetAudio().GetTitle(),
				Performer: attributes.GetAudio().GetPerformer(),
			}); err != nil {
				log.Println(err)
				return nil, status.Error(codes.Unknown, "не удалось")
			}
		default:
			if err := c.MessageUseCase.SendBotFile(ctx, uid, &entity.SendBotFile{
				Receiver:     receiver,
				Drive:        1,
				OriginalName: attributes.GetFilename().GetFileName(),
				FileExt:      fileExt,
				FileSize:     int(filePath.Size),
				FilePath:     c.UploadUseCase.Minio.PublicUrl(c.Conf.Minio.GetBucket(), filePath.FilePath),
				Content:      in.Message,
				FileId:       &filePath.FileId,
			}); err != nil {
				log.Println(err)
				return nil, status.Error(codes.Unknown, "не удалось")
			}
		}
	default:
		return nil, status.Error(codes.Unknown, c.Locale.Localize("general_error"))
	}

	return &chatPb.SendMediaResponse{
		Success: true,
	}, nil
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
