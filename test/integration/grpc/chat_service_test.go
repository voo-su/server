package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
	chatPb "voo.su/api/grpc/pb"
)

func TestChatListService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := grpc.NewClient("127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryInterceptor),
	)
	if err != nil {
		log.Println("Не удалось подключиться к gRPC серверу")
	}

	defer conn.Close()

	client := chatPb.NewChatServiceClient(conn)
	res, err := client.GetChats(ctx, &chatPb.GetChatsRequest{})
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса List: %v", err)
		return
	}

	if len(res.Items) == 0 {
		t.Error("Список чатов пуст, ожидались элементы")
	}

	for _, item := range res.Items {
		if item.Id == 0 {
			t.Error("Обнаружен чат с Id равным 0, ожидался валидный идентификатор")
		}
		if item.Receiver.ChatType == 0 {
			t.Error("Обнаружен чат с ChatType равным 0, ожидался валидный тип чата")
		}
		if item.Name == "" && item.Username == "" {
			t.Error("Обнаружен чат с пустыми Name и Username, ожидались заполненные значения хотя бы для одного поля")
		}
		if item.UpdatedAt == "" {
			t.Error("Обнаружен чат с пустым UpdatedAt, ожидалась корректная дата обновления")
		}
	}
}
