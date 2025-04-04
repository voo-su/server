package v1

import (
	"bytes"
	"math"
	"path"
	"strconv"
	"strings"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/strutil"
)

type Upload struct {
	Conf          *config.Config
	Locale        locale.ILocale
	Minio         minio.IMinio
	UploadUseCase *usecase.UploadUseCase
}

func (u *Upload) Avatar(ctx *ginutil.Context) error {
	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams(u.Locale.Localize("image_upload_error"))
	}

	stream, _ := minio.ReadMultipartStream(file)
	object := strutil.GenMediaObjectName("png", 200, 200)
	if err := u.Minio.Write(u.Conf.Minio.GetBucket(), object, stream); err != nil {
		return ctx.Error(u.Locale.Localize("image_upload_error"))
	}

	return ctx.Success(v1Pb.UploadAvatarResponse{
		Avatar: u.Minio.PublicUrl(u.Conf.Minio.GetBucket(), object),
	})
}

func (u *Upload) Upload(ctx *ginutil.Context) error {
	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams(u.Locale.Localize("file_upload_error"))
	}

	var (
		ext       = strings.TrimPrefix(path.Ext(file.Filename), ".")
		width, _  = strconv.Atoi(ctx.Context.DefaultPostForm("width", "0"))
		height, _ = strconv.Atoi(ctx.Context.DefaultPostForm("height", "0"))
	)

	stream, _ := minio.ReadMultipartStream(file)
	if width == 0 || height == 0 {
		meta := pkg.ReadImageMeta(bytes.NewReader(stream))
		width = meta.Width
		height = meta.Height
	}

	object := strutil.GenMediaObjectName(ext, width, height)
	if err := u.Minio.Write(u.Conf.Minio.GetBucket(), object, stream); err != nil {
		return ctx.Error(u.Locale.Localize("file_upload_error"))
	}

	return ctx.Success(v1Pb.UploadImageResponse{
		Src: u.Minio.PublicUrl(u.Conf.Minio.GetBucket(), object),
	})
}

func (u *Upload) InitiateMultipart(ctx *ginutil.Context) error {
	params := &v1Pb.UploadInitiateMultipartRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	info, err := u.UploadUseCase.InitiateMultipartUpload(ctx.Ctx(), &usecase.MultipartInitiateOpt{
		Name:   params.FileName,
		Size:   params.FileSize,
		UserId: ctx.UserId(),
	})
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.UploadInitiateMultipartResponse{
		UploadId:  info.UploadId,
		ShardSize: 5 << 20,
		ShardNum:  int32(math.Ceil(float64(params.FileSize) / float64(5<<20))),
	})
}

func (u *Upload) MultipartUpload(ctx *ginutil.Context) error {
	params := &v1Pb.UploadMultipartRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams(u.Locale.Localize("file_upload_failure"))
	}

	if err = u.UploadUseCase.MultipartUpload(ctx.Ctx(), &usecase.MultipartUploadOpt{
		UserId:     ctx.UserId(),
		UploadId:   params.UploadId,
		SplitIndex: int(params.SplitIndex),
		SplitNum:   int(params.SplitNum),
		File:       file,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	if params.SplitIndex != params.SplitNum-1 {
		return ctx.Success(&v1Pb.UploadMultipartResponse{
			IsMerge: false,
		})
	}

	return ctx.Success(&v1Pb.UploadMultipartResponse{
		UploadId: params.UploadId,
		IsMerge:  true,
	})
}
