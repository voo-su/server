package event

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"voo.su/internal/config"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	group_chat2 "voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/internal/transport/ws/event/chat"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/jsonutil"
)

type ChatEvent struct {
	Redis               *redis.Client
	Config              *config.Config
	RoomStorage         *cache.RoomStorage
	GroupChatMemberRepo *group_chat2.GroupChatMember

	GroupChatMemberService *service.GroupChatMemberService
	Handler                *chat.Handler
}

func (c *ChatEvent) OnOpen(client socket.IClient) {
	ctx := context.TODO()
	ids := c.GroupChatMemberRepo.GetUserGroupIds(ctx, client.Uid())
	rooms := make([]*cache.RoomOption, 0, len(ids))
	for _, id := range ids {
		rooms = append(rooms, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: entity.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      c.Config.ServerId(),
			Cid:      client.Cid(),
		})
	}
	if err := c.RoomStorage.BatchAdd(ctx, rooms); err != nil {
		log.Println("Ошибка при вступлении в групповой чат", err.Error())
	}

	c.Redis.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventContactStatus,
		"data": jsonutil.Encode(map[string]any{
			"user_id": client.Uid(),
			"status":  1,
		}),
	}))
}

func (c *ChatEvent) OnMessage(client socket.IClient, message []byte) {
	val, err := sonic.Get(message, "event")
	if err == nil {
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
			RoomType: entity.RoomImGroup,
			Number:   strconv.Itoa(id),
			Sid:      c.Config.ServerId(),
			Cid:      client.Cid(),
		})
	}
	if err := c.RoomStorage.BatchDel(ctx, rooms); err != nil {
		log.Println("Ошибка при выходе из группового чата", err.Error())
	}

	c.Redis.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventContactStatus,
		"data": jsonutil.Encode(map[string]any{
			"user_id": client.Uid(),
			"status":  0,
		}),
	}))
}
