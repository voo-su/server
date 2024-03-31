package v1

import (
	"bytes"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"
	"voo.su/api/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/strutil"
	"voo.su/pkg/utils"
)

type Upload struct {
	Config       *config.Config
	Filesystem   *filesystem.Filesystem
	SplitService *service.SplitService
}

func (u *Upload) Avatar(ctx *core.Context) error {
	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	stream, _ := filesystem.ReadMultipartStream(file)
	object := fmt.Sprintf("avatar/%s/%s", time.Now().Format("20060102"), strutil.GenImageName("png", 200, 200))
	if err := u.Filesystem.Default.Write(stream, object); err != nil {
		return ctx.ErrorBusiness("Ошибка загрузки файла")
	}

	return ctx.Success(api_v1.UploadAvatarResponse{
		Avatar: u.Filesystem.Default.PublicUrl(object),
	})
}

func (u *Upload) Image(ctx *core.Context) error {
	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams("Не удалось загрузить файл!")
	}

	var (
		ext       = strings.TrimPrefix(path.Ext(file.Filename), ".")
		width, _  = strconv.Atoi(ctx.Context.DefaultPostForm("width", "0"))
		height, _ = strconv.Atoi(ctx.Context.DefaultPostForm("height", "0"))
	)

	stream, _ := filesystem.ReadMultipartStream(file)
	if width == 0 || height == 0 {
		meta := utils.ReadImageMeta(bytes.NewReader(stream))
		width = meta.Width
		height = meta.Height
	}

	object := fmt.Sprintf("image/common/%s/%s", time.Now().Format("20060102"), strutil.GenImageName(ext, width, height))
	if err := u.Filesystem.Default.Write(stream, object); err != nil {
		return ctx.ErrorBusiness("Не удалось загрузить файл")
	}

	return ctx.Success(api_v1.UploadImageResponse{
		Src: u.Filesystem.Default.PublicUrl(object),
	})
}

func (u *Upload) InitiateMultipart(ctx *core.Context) error {
	params := &api_v1.UploadInitiateMultipartRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	info, err := u.SplitService.InitiateMultipartUpload(ctx.Ctx(), &service.MultipartInitiateOpt{
		Name:   params.FileName,
		Size:   params.FileSize,
		UserId: ctx.UserId(),
	})
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.UploadInitiateMultipartResponse{
		UploadId:    info.UploadId,
		UploadIdMd5: encrypt.Md5(info.UploadId),
		SplitSize:   2 << 20,
	})
}

func (u *Upload) MultipartUpload(ctx *core.Context) error {
	params := &api_v1.UploadMultipartRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	file, err := ctx.Context.FormFile("file")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	err = u.SplitService.MultipartUpload(ctx.Ctx(), &service.MultipartUploadOpt{
		UserId:     ctx.UserId(),
		UploadId:   params.UploadId,
		SplitIndex: int(params.SplitIndex),
		SplitNum:   int(params.SplitNum),
		File:       file,
	})
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if params.SplitIndex != params.SplitNum-1 {
		return ctx.Success(&api_v1.UploadMultipartResponse{
			IsMerge: false,
		})
	}

	return ctx.Success(&api_v1.UploadMultipartResponse{
		UploadId: params.UploadId,
		IsMerge:  true,
	})
}
