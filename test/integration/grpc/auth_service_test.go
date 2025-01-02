// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package grpc

import (
	"fmt"
	"testing"
	authPb "voo.su/api/grpc/gen/go/pb"
)

func TestAuthLoginService(t *testing.T) {
	conn, ctx, cancel := ConnectToGRPC("127.0.0.1:50051")
	defer cancel()
	defer conn.Close()

	client := authPb.NewAuthServiceClient(conn)
	res, err := client.Login(ctx, &authPb.AuthLoginRequest{
		Email:    "test@test.test",
		Platform: "android",
	})
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса Login: %v", err)
		return
	}

	if res.Token == "" {
		t.Error("Получен пустой токен, ожидался валидный токен")
	}
	if res.ExpiresIn == 0 {
		t.Error("Поле ExpiresIn равно 0, ожидалось ненулевое время жизни токена")
	}

	fmt.Printf("Token: %s\n", res.Token)
	fmt.Printf("Expires In: %d секунд\n", res.ExpiresIn)
}
