package usecase

import (
	"context"
	"errors"
	"fmt"
	"math"
	"mime/multipart"
	"path"
	"strings"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type SplitUseCase struct {
	*repo.Source
	Split      *repo.Split
	Config     *config.Config
	FileSystem *filesystem.Filesystem
}

func NewSplitUseCase(
	source *repo.Source,
	repo *repo.Split,
	conf *config.Config,
	fileSystem *filesystem.Filesystem,
) *SplitUseCase {
	return &SplitUseCase{
		Source:     source,
		Split:      repo,
		Config:     conf,
		FileSystem: fileSystem,
	}
}

type MultipartInitiateOpt struct {
	UserId int
	Name   string
	Size   int64
}

func (s *SplitUseCase) InitiateMultipartUpload(ctx context.Context, params *MultipartInitiateOpt) (*model.Split, error) {
	num := math.Ceil(float64(params.Size) / float64(3<<20))
	m := &model.Split{
		Type:         1,
		Drive:        entity.FileDriveMode(s.FileSystem.Driver()),
		UserId:       params.UserId,
		OriginalName: params.Name,
		SplitNum:     int(num),
		FileExt:      strings.TrimPrefix(path.Ext(params.Name), "."),
		FileSize:     params.Size,
		Path:         fmt.Sprintf("private-tmp/multipart/%s/%s.tmp", timeutil.DateNumber(), encrypt.Md5(strutil.Random(20))),
		Attr:         "{}",
	}
	uploadId, err := s.FileSystem.Default.InitiateMultipartUpload(m.Path, m.OriginalName)
	if err != nil {
		return nil, err
	}

	m.UploadId = uploadId
	if err := s.Source.Db().WithContext(ctx).Create(m).Error; err != nil {
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
	info, err := s.Split.FindByWhere(ctx, "upload_id = ? and type = 1", opt.UploadId)
	if err != nil {
		return err
	}

	stream, err := filesystem.ReadMultipartStream(opt.File)
	if err != nil {
		return err
	}

	dirPath := fmt.Sprintf("private-tmp/%s/%s/%d-%s.tmp", timeutil.DateNumber(), encrypt.Md5(opt.UploadId), opt.SplitIndex, opt.UploadId)
	data := &model.Split{
		Type:         2,
		Drive:        info.Drive,
		UserId:       opt.UserId,
		UploadId:     opt.UploadId,
		OriginalName: info.OriginalName,
		SplitIndex:   opt.SplitIndex,
		SplitNum:     opt.SplitNum,
		Path:         dirPath,
		FileExt:      info.FileExt,
		FileSize:     opt.File.Size,
		Attr:         "{}",
	}
	switch data.Drive {
	case constant.FileDriveLocal:
		_ = s.FileSystem.Default.Write(stream, data.Path)
	default:
		return errors.New("неизвестный тип драйвера файла")
	}
	if err := s.Source.Db().Create(data).Error; err != nil {
		return err
	}

	if opt.SplitNum == opt.SplitIndex+1 {
		err = s.merge(info)
	}
	return err
}

func (s *SplitUseCase) merge(info *model.Split) error {
	items, err := s.Split.GetSplitList(context.TODO(), info.UploadId)
	if err != nil {
		return err
	}

	switch info.Drive {
	case constant.FileDriveLocal:
		for _, item := range items {
			stream, err := s.FileSystem.Default.ReadStream(item.Path)
			if err != nil {
				return err
			}
			if err := s.FileSystem.Local.AppendWrite(stream, info.Path); err != nil {
				return err
			}
		}
	default:
		return errors.New("неизвестный тип драйвера файла")
	}
	return nil
}
