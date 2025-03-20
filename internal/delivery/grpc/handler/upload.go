package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
	uploadPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg/locale"
)

type Upload struct {
	uploadPb.UnimplementedUploadServiceServer
	Conf          *config.Config
	Locale        locale.ILocale
	UploadUseCase *usecase.UploadUseCase
}

func NewUploadHandler(
	conf *config.Config,
	locale locale.ILocale,
	uploadUseCase *usecase.UploadUseCase,
) *Upload {
	return &Upload{
		Conf:          conf,
		Locale:        locale,
		UploadUseCase: uploadUseCase,
	}
}

func (u *Upload) SaveFilePart(ctx context.Context, in *uploadPb.SaveFilePartRequest) (*uploadPb.SaveFilePartResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err := u.UploadUseCase.RetrySaveFilePart(3, in.GetFileId(), in.GetFilePart(), in.GetBytes()); err != nil {
		log.Printf("Не удалось сохранить часть файла %d, ошибка: %v", in.GetFilePart(), err)
		return nil, errors.New("не удалось сохранить часть файла")
	}

	return &uploadPb.SaveFilePartResponse{
		Success: true,
	}, nil
}

func (u *Upload) GetFile(ctx context.Context, in *uploadPb.GetFileRequest) (*uploadPb.GetFileResponse, error) {
	if in.GetDocumentLocation() == nil {
		return nil, fmt.Errorf("некорректный запрос: отсутствует document_location")
	}

	fileId, err := uuid.Parse(in.GetDocumentLocation().GetId())
	if err != nil {
		return nil, errors.New(u.Locale.Localize("network_error"))
	}

	bytes, err := u.UploadUseCase.GetFile(ctx, fileId, in.GetOffset(), in.GetLimit())
	if err != nil {
		return nil, errors.New("ошибка")
	}

	return &uploadPb.GetFileResponse{
		Bytes: bytes,
	}, nil
}
