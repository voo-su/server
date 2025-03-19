package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
	contactPb "voo.su/api/grpc/pb"
)

func TestContactListService(t *testing.T) {
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

	client := contactPb.NewContactServiceClient(conn)

	res, err := client.GetContacts(ctx, &contactPb.GetContactsRequest{})
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса List: %v", err)
		return
	}

	if len(res.Items) == 0 {
		t.Error("Список контактов пуст, ожидались элементы")
	}

	for _, item := range res.Items {
		if item.Id == 0 {
			t.Error("Обнаружен контакт с Id равным 0, ожидался валидный идентификатор")
		}
		if item.Username == "" {
			t.Error("Обнаружен контакт с пустым Username, ожидалось заполненное значение")
		}
		if item.Name == "" && item.Surname == "" {
			t.Error("Обнаружен контакт с пустыми Name и Surname, ожидались заполненные значения хотя бы для одного поля")
		}
	}
}
