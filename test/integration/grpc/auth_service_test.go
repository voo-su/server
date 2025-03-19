package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
	authPb "voo.su/api/grpc/pb"
)

func TestAuthLoginService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Не удалось подключиться к gRPC серверу")
	}

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
