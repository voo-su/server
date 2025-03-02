package chat

import (
	"encoding/json"
	"fmt"
	_nats "github.com/nats-io/nats.go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	chatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/logger"
)

func (c *Chat) GetUpdates(req *chatPb.UpdatesRequest, stream chatPb.ChatService_GetUpdatesServer) error {
	ctx := stream.Context()

	select {
	case <-ctx.Done():
		log.Println("Запрос был отменен или истекло время")
		return ctx.Err()
	default:
	}

	claims, _, err := grpcutil.GrpcToken(ctx, c.Locale, constant.GuardGrpcAuth, c.Conf.App.Jwt.Secret)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	c.mu.Lock()
	if _, exists := c.userChannels[int64(uid)]; !exists {
		c.userChannels[int64(uid)] = make(chan *chatPb.Update)
	}

	userChannel := c.userChannels[int64(uid)]
	c.mu.Unlock()

	onMessage := func(receiverId int64, _message entity.Message) {
		c.mu.Lock()
		receiverChannel, isReceiverConnected := c.userChannels[receiverId]
		c.mu.Unlock()

		if isReceiverConnected {
			update := &chatPb.Update{
				Update: &chatPb.Update_NewMessage{
					NewMessage: &chatPb.UpdateNewMessage{
						Message: &chatPb.MessageItem{
							Id: _message.Id,
							Receiver: &chatPb.Receiver{
								ChatType:   int32(_message.ChatType),
								ReceiverId: int64(_message.ReceiverId),
							},
							MsgType:   int32(_message.MsgType),
							UserId:    int64(_message.UserId),
							Content:   _message.Content,
							IsRead:    _message.IsRead,
							CreatedAt: _message.CreatedAt,
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

	sub, err := c.Nats.Subscribe(constant.ImTopicChat, func(msg *_nats.Msg) {
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

			c.mu.Lock()
			receiverChannel, isReceiverConnected := c.userChannels[receiverId]
			c.mu.Unlock()

			if isReceiverConnected {
				update := &chatPb.Update{
					Update: &chatPb.Update_UserTyping{
						UserTyping: &chatPb.UpdateUserTyping{
							Receiver: &chatPb.Receiver{
								ChatType:   1,
								ReceiverId: receiverId,
							},
							UserId:   int64(_in.SenderId),
							IsTyping: true,
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

	c.mu.Lock()
	delete(c.userChannels, int64(uid))
	c.mu.Unlock()

	return nil
}
