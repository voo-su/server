package usecase

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"log"
	"math"
	"mime/multipart"
	"path"
	"strings"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/strutil"
)

type UploadUseCase struct {
	Conf          *config.Config
	Locale        locale.ILocale
	Source        *infrastructure.Source
	FileRepo      *postgresRepo.FileRepository
	FileSplitRepo *postgresRepo.FileSplitRepository
	Minio         minio.IMinio
}

func NewUploadUseCase(
	conf *config.Config,
	locale locale.ILocale,
	source *infrastructure.Source,
	fileRepo *postgresRepo.FileRepository,
	fileSplitRepo *postgresRepo.FileSplitRepository,
	minio minio.IMinio,
) *UploadUseCase {
	return &UploadUseCase{
		Conf:          conf,
		Locale:        locale,
		Source:        source,
		FileRepo:      fileRepo,
		FileSplitRepo: fileSplitRepo,
		Minio:         minio,
	}
}

type MultipartInitiateOpt struct {
	UserId int
	Name   string
	Size   int64
}

func (u *UploadUseCase) InitiateMultipartUpload(ctx context.Context, params *MultipartInitiateOpt) (*postgresModel.FileSplit, error) {
	num := math.Ceil(float64(params.Size) / float64(5*1024*1024))

	now := time.Now()
	m := &postgresModel.FileSplit{
		Type:         1,
		Drive:        0,
		UserId:       params.UserId,
		OriginalName: params.Name,
		SplitNum:     int(num),
		FileExt:      strings.TrimPrefix(path.Ext(params.Name), "."),
		FileSize:     params.Size,
		Path:         fmt.Sprintf("multipart/%s/%s.tmp", now.Format("20060102"), uuid.New().String()),
		Attr:         "{}",
	}

	uploadId, err := u.Minio.InitiateMultipartUpload(u.Conf.Minio.GetBucket(), m.Path)
	if err != nil {
		return nil, err
	}

	m.UploadId = uploadId

	if err := u.Source.Postgres().
		WithContext(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

type MultipartUploadOpt struct {
	UserId     int
	UploadId   string
	SplitIndex int
	SplitNum   int
	File       *multipart.FileHeader
}

func (u *UploadUseCase) MultipartUpload(ctx context.Context, opt *MultipartUploadOpt) error {
	info, err := u.FileSplitRepo.FindByWhere(ctx, "upload_id = ? AND type = ?", opt.UploadId, 1)
	if err != nil {
		return err
	}

	stream, err := minio.ReadMultipartStream(opt.File)
	if err != nil {
		return err
	}

	data := &postgresModel.FileSplit{
		Type:         2,
		Drive:        info.Drive,
		UserId:       opt.UserId,
		UploadId:     opt.UploadId,
		OriginalName: info.OriginalName,
		SplitIndex:   opt.SplitIndex,
		SplitNum:     opt.SplitNum,
		Path:         "",
		FileExt:      info.FileExt,
		FileSize:     opt.File.Size,
		Attr:         "{}",
	}

	read := bytes.NewReader(stream)

	objectPart, err := u.Minio.PutObjectPart(
		u.Conf.Minio.GetBucket(),
		info.Path,
		info.UploadId,
		opt.SplitIndex,
		read,
		read.Size(),
	)
	if err != nil {
		return err
	}

	if objectPart.PartObjectName != "" {
		data.Path = objectPart.PartObjectName
	}

	data.Attr = jsonutil.Encode(objectPart)

	if err = u.Source.Postgres().Create(data).Error; err != nil {
		return err
	}

	if opt.SplitNum == opt.SplitIndex {
		err = u.merge(info)
	}

	return err
}

func (u *UploadUseCase) merge(info *postgresModel.FileSplit) error {
	items, err := u.FileSplitRepo.FindAll(context.Background(), func(db *gorm.DB) {
		db.Where("upload_id =? AND type = 2", info.UploadId).
			Order("split_index asc")
	})

	if err != nil {
		return err
	}

	parts := make([]minio.ObjectPart, 0)
	for _, item := range items {
		var obj minio.ObjectPart
		if err = jsonutil.Decode(item.Attr, &obj); err != nil {
			return err
		}

		parts = append(parts, obj)
	}

	if err := u.Minio.CompleteMultipartUpload(u.Conf.Minio.GetBucket(), info.Path, info.UploadId, parts); err != nil {
		return err
	}

	return nil
}

func (u *UploadUseCase) RetrySaveFilePart(maxRetries int, fileId int64, filePart int32, data []byte) error {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		if err := u.saveFilePart(fileId, filePart, data); err != nil {
			lastErr = err
			log.Printf("Ошибка при сохранении части %d (попытка %d/%d): %v", filePart, i+1, maxRetries, err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		return nil
	}
	return lastErr
}

func (u *UploadUseCase) saveFilePart(fileID int64, filePart int32, data []byte) error {
	partPath := fmt.Sprintf("%d/%d.part", fileID, filePart)
	if err := u.Minio.Write(u.Conf.Minio.Bucket, partPath, data); err != nil {
		return fmt.Errorf("не удалось сохранить часть %d в MinIO: %v", filePart, err)
	}

	return nil
}

type AssembleFilePart struct {
	FilePath string
	Size     int64
}

func (u *UploadUseCase) AssembleFileParts(ctx context.Context, uid int, fileId int64, totalParts int32, originalName string, fileExt string) (*AssembleFilePart, error) {
	objectName := strutil.GenMediaObjectName(fileExt, 200, 200)

	var totalSize int64
	uploadId, err := u.Minio.InitiateMultipartUpload(u.Conf.Minio.GetBucket(), objectName)
	if err != nil {
		return nil, err
	}

	var partsInfo = make([]minio.ObjectPart, 0)
	for i := int32(0); i < totalParts; i++ {
		objectPartPath := fmt.Sprintf("%d/%d.part", fileId, i)
		obj, err := u.Minio.GetObject(u.Conf.Minio.Bucket, objectPartPath)
		if err != nil {
			return nil, fmt.Errorf("ошибка получения части %d из MinIO: %v", i, err)
		}

		partData, err := io.ReadAll(obj)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения данных части %d: %v", i, err)
		}

		totalSize += int64(len(partData))

		uploadInfo, err := u.Minio.PutObjectPart(u.Conf.Minio.Bucket, objectName, uploadId, int(i+1), bytes.NewReader(partData), int64(len(partData)))
		if err != nil {
			return nil, fmt.Errorf("ошибка загрузки части %d: %v", i, err)
		}

		partsInfo = append(partsInfo, minio.ObjectPart{
			PartNumber: uploadInfo.PartNumber,
			ETag:       uploadInfo.ETag,
		})
	}

	if err := u.Minio.CompleteMultipartUpload(u.Conf.Minio.Bucket, objectName, uploadId, partsInfo); err != nil {
		return nil, fmt.Errorf("ошибка завершения сборки файла: %v", err)
	}

	if err := u.FileRepo.Create(ctx, &postgresModel.File{
		OriginalName: originalName,
		ObjectName:   objectName,
		Size:         int(totalSize),
		MimeType:     fileExt,
		CreatedBy:    uid,
		CreatedAt:    time.Now(),
	}); err != nil {
		return nil, fmt.Errorf("ошибка сохранения информации о файле в базе: %v", err)
	}

	return &AssembleFilePart{
		FilePath: objectName,
		Size:     totalSize,
	}, nil
}

func (u *UploadUseCase) GetFile(ctx context.Context, fileId uuid.UUID, offset int64, limit int32) ([]byte, error) {
	file, err := u.FileRepo.FindByWhere(ctx, "id = ?", fileId)
	if err != nil {
		return nil, err
	}

	obj, err := u.Minio.GetObject(u.Conf.Minio.Bucket, file.ObjectName)
	if err != nil {
		return nil, fmt.Errorf("файл не найден: %v", err)
	}
	defer obj.Close()

	if _, err := obj.Seek(offset, io.SeekStart); err != nil {
		return nil, fmt.Errorf("ошибка при установке offset: %v", err)
	}

	data := make([]byte, limit)
	n, err := obj.Read(data)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	return data[:n], nil
}
