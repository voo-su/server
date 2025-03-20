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
	commonPb "voo.su/api/grpc/pb/common"
)

func TestSaveFile(t *testing.T) {
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

	var fileId int64 = 1234567890

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

		if err := sendFilePart(ctx, client, fileId, int32(part), buffer[:n]); err != nil {
			log.Printf("Ошибка при отправке части файла: %v", err)
			return
		}
		part++
	}
}

func TestGetFile(t *testing.T) {
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

	if err = downloadFile(ctx, client, "ad7c81d4-d013-4a42-80cb-97e21e12439a", "logo.svg"); err != nil {
		log.Fatalf("Ошибка загрузки файла: %v", err)
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

func downloadFile(ctx context.Context, client uploadPb.UploadServiceClient, fileID string, outputPath string) error {
	chunkSize := 1024 * 1024

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	offset := int64(0)

	for {
		resp, err := client.GetFile(ctx, &uploadPb.GetFileRequest{
			Location: &uploadPb.GetFileRequest_DocumentLocation{
				DocumentLocation: &commonPb.InputDocumentFileLocation{
					Id: fileID,
				},
			},
			Offset: offset,
			Limit:  int32(chunkSize),
		})
		if err != nil {
			return fmt.Errorf("ошибка получения данных: %v", err)
		}

		if len(resp.Bytes) == 0 {
			break
		}

		_, err = file.Write(resp.Bytes)
		if err != nil {
			return fmt.Errorf("ошибка записи данных в файл: %v", err)
		}

		offset += int64(len(resp.Bytes))
		log.Printf("Загружено %d байт...\n", offset)
	}

	return nil
}
