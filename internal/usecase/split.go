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
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/minio"
)

type SplitUseCase struct {
	*repo.Source
	SplitRepo *repo.Split
	Config    *config.Config
	Minio     minio.IMinio
}

func NewSplitUseCase(
	source *repo.Source,
	splitRepo *repo.Split,
	conf *config.Config,
	minio minio.IMinio,
) *SplitUseCase {
	return &SplitUseCase{
		Source:    source,
		SplitRepo: splitRepo,
		Config:    conf,
		Minio:     minio,
	}
}

type MultipartInitiateOpt struct {
	UserId int
	Name   string
	Size   int64
}

func (s *SplitUseCase) InitiateMultipartUpload(ctx context.Context, params *MultipartInitiateOpt) (*model.Split, error) {
	num := math.Ceil(float64(params.Size) / float64(5*1024*1024))

	now := time.Now()
	m := &model.Split{
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

	uploadId, err := s.Minio.InitiateMultipartUpload(s.Minio.BucketPrivateName(), m.Path)
	if err != nil {
		return nil, err
	}

	m.UploadId = uploadId

	if err := s.Source.Db().
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

func (s *SplitUseCase) MultipartUpload(ctx context.Context, opt *MultipartUploadOpt) error {
	info, err := s.SplitRepo.FindByWhere(ctx, "upload_id = ? AND type = 1", opt.UploadId)
	if err != nil {
		return err
	}

	stream, err := minio.ReadMultipartStream(opt.File)
	if err != nil {
		return err
	}

	data := &model.Split{
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

	objectPart, err := s.Minio.PutObjectPart(
		s.Minio.BucketPrivateName(),
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

	if err = s.Source.Db().Create(data).Error; err != nil {
		return err
	}

	if opt.SplitNum == opt.SplitIndex {
		err = s.merge(info)
	}

	return err
}

func (s *SplitUseCase) merge(info *model.Split) error {
	items, err := s.SplitRepo.FindAll(context.Background(), func(db *gorm.DB) {
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

	if err := s.Minio.CompleteMultipartUpload(s.Minio.BucketPrivateName(), info.Path, info.UploadId, parts); err != nil {
		return err
	}

	return nil
}
