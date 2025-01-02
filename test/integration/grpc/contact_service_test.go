// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package grpc

import (
	"testing"
	contactPb "voo.su/api/grpc/gen/go/pb"
)

func TestContactListService(t *testing.T) {
	conn, ctx, cancel := ConnectToGRPC("127.0.0.1:50051")
	defer cancel()
	defer conn.Close()

	client := contactPb.NewContactServiceClient(conn)
	res, err := client.List(ctx, &contactPb.GetContactListRequest{})
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
