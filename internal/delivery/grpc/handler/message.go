package handler

import (
	"context"
	"encoding/json"
	"fmt"
	_nats "github.com/nats-io/nats.go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"sync"
	messagePb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/logger"
	"voo.su/pkg/nats"
	"voo.su/pkg/strutil"
)

type Message struct {
	messagePb.UnimplementedMessageServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase

	mu           sync.Mutex
	userChannels map[int64]chan *messagePb.Update

	Nats nats.INatsClient
}

func NewMessageHandler(
	conf *config.Config,
	locale locale.ILocale,
	chatUseCase *usecase.ChatUseCase,
	messageUseCase usecase.IMessageUseCase,

	nats nats.INatsClient,
) *Message {
	return &Message{
		Conf:           conf,
		Locale:         locale,
		ChatUseCase:    chatUseCase,
		MessageUseCase: messageUseCase,

		userChannels: make(map[int64]chan *messagePb.Update),

		Nats: nats,
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

func (m *Message) ViewMessages(ctx context.Context, in *messagePb.ViewMessagesRequest) (*messagePb.ViewMessagesResponse, error) {
	uid := grpcutil.UserId(ctx)

	// TODO

	m.ChatUseCase.UnreadCacheRepo.Reset(ctx, int(in.ChatType), int(in.ReceiverId), uid)

	return &messagePb.ViewMessagesResponse{}, nil
}

func (m *Message) GetUpdates(req *messagePb.UpdatesRequest, stream messagePb.MessageService_GetUpdatesServer) error {
	ctx := stream.Context()

	select {
	case <-ctx.Done():
		log.Println("Запрос был отменен или истекло время")
		return ctx.Err()
	default:
	}

	claims, _, err := grpcutil.GrpcToken(ctx, m.Locale, constant.GuardGrpcAuth, m.Conf.App.Jwt.Secret)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	m.mu.Lock()
	if _, exists := m.userChannels[int64(uid)]; !exists {
		m.userChannels[int64(uid)] = make(chan *messagePb.Update)
	}

	userChannel := m.userChannels[int64(uid)]
	m.mu.Unlock()

	onMessage := func(receiverId int64, _message entity.Message) {
		m.mu.Lock()
		receiverChannel, isReceiverConnected := m.userChannels[receiverId]
		m.mu.Unlock()

		if isReceiverConnected {
			update := &messagePb.Update{
				Update: &messagePb.Update_NewMessage{
					NewMessage: &messagePb.UpdateNewMessage{
						Message: &messagePb.MessageItem{
							Id:         _message.Id,
							ChatType:   int32(_message.ChatType),
							MsgType:    int32(_message.MsgType),
							ReceiverId: int64(_message.ReceiverId),
							UserId:     int64(_message.UserId),
							Content:    _message.Content,
							IsRead:     _message.IsRead,
							CreatedAt:  _message.CreatedAt,
						},
					},
				},
			}

			select {
			case receiverChannel <- update:
				log.Printf("Обновление отправлено пользователю %d", receiverId)
			default:
				log.Printf("Пользователь %d подключен, но не принимает обновления", receiverId)
			}
		}
	}

	sub, err := m.Nats.Subscribe(constant.ImTopicChat, func(msg *_nats.Msg) {
		var in entity.SubscribeContent
		if err := json.Unmarshal(msg.Data, &in); err != nil {
			log.Printf("Ошибка при разборе JSON: %s", err)
			return
		}

		switch in.Event {
		case constant.SubEventImMessage:
			var _in entity.ConsumeMessage
			if err := json.Unmarshal([]byte(in.Data), &_in); err != nil {
				logger.Errorf("Ошибка декодирования JSON: %s", err.Error())
				return
			}
			fmt.Println(_in)
			for _, rid := range _in.UserIds {

				onMessage(int64(rid), _in.Message)
			}

		case constant.SubEventImMessageKeyboard:
			var _in entity.ConsumeChatKeyboard
			if err := json.Unmarshal([]byte(in.Data), &_in); err != nil {
				logger.Errorf("Ошибка декодирования JSON: %s", err.Error())
				return
			}

			receiverId := int64(_in.ReceiverId)

			m.mu.Lock()
			receiverChannel, isReceiverConnected := m.userChannels[receiverId]
			m.mu.Unlock()

			if isReceiverConnected {
				update := &messagePb.Update{
					Update: &messagePb.Update_UserTyping{
						UserTyping: &messagePb.UpdateUserTyping{
							ChatType:   1,
							ReceiverId: receiverId,
							UserId:     int64(_in.SenderId),
							IsTyping:   true,
						},
					},
				}

				select {
				case receiverChannel <- update:
					log.Printf("Обновление отправлено пользователю %d", receiverId)
				default:
					log.Printf("Пользователь %d подключен, но не принимает обновления", receiverId)
				}
			}
		}

	})
	if err != nil {
		log.Println(err)
	}

	defer sub.Unsubscribe()

	for update := range userChannel {
		if err := stream.Send(update); err != nil {
			log.Printf("Ошибка отправки обновления: %v", err)
			break
		}
	}

	m.mu.Lock()
	delete(m.userChannels, int64(uid))
	m.mu.Unlock()

	return nil
}
