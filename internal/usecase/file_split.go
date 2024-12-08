package usecase

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
	"mime/multipart"
	"path"
	"strings"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/repository"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/minio"
)

type FileSplitUseCase struct {
	*repository.Source
	Conf          *config.Config
	FileSplitRepo *repo.FileSplit
	Minio         minio.IMinio
}

func NewFileSplitUseCase(
	source *repository.Source,
	conf *config.Config,
	fileSplitRepo *repo.FileSplit,
	minio minio.IMinio,
) *FileSplitUseCase {
	return &FileSplitUseCase{
		Source:        source,
		Conf:          conf,
		FileSplitRepo: fileSplitRepo,
		Minio:         minio,
	}
}

type MultipartInitiateOpt struct {
	UserId int
	Name   string
	Size   int64
}

func (f *FileSplitUseCase) InitiateMultipartUpload(ctx context.Context, params *MultipartInitiateOpt) (*model.FileSplit, error) {
	num := math.Ceil(float64(params.Size) / float64(5*1024*1024))

	now := time.Now()
	m := &model.FileSplit{
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

	uploadId, err := f.Minio.InitiateMultipartUpload(f.Conf.Minio.GetBucket(), m.Path)
	if err != nil {
		return nil, err
	}

	m.UploadId = uploadId

	if err := f.Source.Db().
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

func (f *FileSplitUseCase) MultipartUpload(ctx context.Context, opt *MultipartUploadOpt) error {
	info, err := f.FileSplitRepo.FindByWhere(ctx, "upload_id = ? AND type = 1", opt.UploadId)
	if err != nil {
		return err
	}

	stream, err := minio.ReadMultipartStream(opt.File)
	if err != nil {
		return err
	}

	data := &model.FileSplit{
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

	objectPart, err := f.Minio.PutObjectPart(
		f.Conf.Minio.GetBucket(),
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

	if err = f.Source.Db().Create(data).Error; err != nil {
		return err
	}

	if opt.SplitNum == opt.SplitIndex {
		err = f.merge(info)
	}

	return err
}

func (f *FileSplitUseCase) merge(info *model.FileSplit) error {
	items, err := f.FileSplitRepo.FindAll(context.Background(), func(db *gorm.DB) {
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

	if err := f.Minio.CompleteMultipartUpload(f.Conf.Minio.GetBucket(), info.Path, info.UploadId, parts); err != nil {
		return err
	}

	return nil
}
