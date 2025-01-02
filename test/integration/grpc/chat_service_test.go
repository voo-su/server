// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package grpc

import (
	"testing"
	chatPb "voo.su/api/grpc/gen/go/pb"
)

func TestChatListService(t *testing.T) {
	conn, ctx, cancel := ConnectToGRPC("127.0.0.1:50051")
	defer cancel()
	defer conn.Close()

	client := chatPb.NewChatServiceClient(conn)
	res, err := client.List(ctx, &chatPb.GetChatListRequest{})
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
		if item.ChatType == 0 {
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
