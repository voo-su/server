package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"testing"
	"time"
	uploadPb "voo.su/api/grpc/pb"
)

func TestSaveFilePart(t *testing.T) {
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

	client := uploadPb.NewUploadServiceClient(conn)

	file, err := os.Open("../../../assets/logo.svg")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var fileID int64 = 1234567890
	buffer := make([]byte, 512*1024)
	part := 0
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Ошибка чтения файла: %v", err)
		}

		if err := sendFilePart(ctx, client, fileID, int32(part), buffer[:n]); err != nil {
			log.Printf("Ошибка при отправке части файла: %v", err)
			return
		}
		part++
	}
}

func sendFilePart(ctx context.Context, client uploadPb.UploadServiceClient, fileId int64, part int32, data []byte) error {
	req := &uploadPb.SaveFilePartRequest{
		FileId:   fileId,
		FilePart: part,
		Bytes:    data,
	}

	resp, err := client.SaveFilePart(ctx, req)
	if err != nil {
		return fmt.Errorf("не удалось отправить часть файла: %v", err)
	}

	if !resp.GetSuccess() {
		return fmt.Errorf("не удалось сохранить часть файла: file_id=%d, part=%d", fileId, part)
	}

	return nil
}
