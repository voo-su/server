package event

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/transport/ws/event/chat"
	"voo.su/internal/usecase"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/socket"
)

type ChatEvent struct {
	Redis               *redis.Client
	Conf                *config.Config
	RoomStorage         *cache.RoomStorage
	GroupChatMemberRepo *repo.GroupChatMember

	GroupChatMemberUseCase *usecase.GroupChatMemberUseCase
	Handler                *chat.Handler
}

func (c *ChatEvent) OnOpen(client socket.IClient) {
	ctx := context.TODO()
	ids := c.GroupChatMemberRepo.GetUserGroupIds(ctx, client.Uid())
	rooms := make([]*cache.RoomOption, 0, len(ids))
	for _, id := range ids {
		rooms = append(rooms, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: constant.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      c.Conf.ServerId(),
			Cid:      client.Cid(),
		})
	}
	if err := c.RoomStorage.BatchAdd(ctx, rooms); err != nil {
		log.Println("Ошибка при вступлении в групповой чат", err.Error())
	}

	c.Redis.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventContactStatus,
		"data": jsonutil.Encode(map[string]any{
			"user_id": client.Uid(),
			"status":  1,
		}),
	}))
}

func (c *ChatEvent) OnMessage(client socket.IClient, message []byte) {
	val, err := sonic.Get(message, "event")
	if err != nil {
		return
	}
	event, _ := val.String()
	if event != "" {
		c.Handler.Call(context.TODO(), client, event, message)
	}
}

func (c *ChatEvent) OnClose(client socket.IClient, code int, text string) {
	log.Println("Закрытие клиента: ", client.Uid(), client.Cid(), client.Channel().Name(), code, text)
	ctx := context.TODO()
	ids := c.GroupChatMemberRepo.GetUserGroupIds(ctx, client.Uid())
	rooms := make([]*cache.RoomOption, 0, len(ids))
	for _, id := range ids {
		rooms = append(rooms, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: constant.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      c.Conf.ServerId(),
			Cid:      client.Cid(),
		})
	}
	if err := c.RoomStorage.BatchDel(ctx, rooms); err != nil {
		log.Println("Ошибка при выходе из группового чата", err.Error())
	}

	c.Redis.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventContactStatus,
		"data": jsonutil.Encode(map[string]any{
			"user_id": client.Uid(),
			"status":  0,
		}),
	}))
}
