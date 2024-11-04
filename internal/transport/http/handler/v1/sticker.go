package v1

import (
	"bytes"
	"fmt"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/utils"
)

type Sticker struct {
	StickerUseCase *usecase.StickerUseCase
	Filesystem     *filesystem.Filesystem
	RedisLock      *cache.RedisLock
}

func (c *Sticker) CollectList(ctx *core.Context) error {
	var (
		uid  = ctx.UserId()
		resp = &v1Pb.StickerListResponse{
			SysSticker:     make([]*v1Pb.StickerListResponse_SysSticker, 0),
			CollectSticker: make([]*v1Pb.StickerListItem, 0),
		}
	)
	if ids := c.StickerUseCase.StickerRepo.GetUserInstallIds(uid); len(ids) > 0 {
		if items, err := c.StickerUseCase.StickerRepo.FindByIds(ctx.Ctx(), ids); err == nil {
			for _, item := range items {
				data := &v1Pb.StickerListResponse_SysSticker{
					StickerId: int32(item.Id),
					Url:       item.Icon,
					Name:      item.Name,
					List:      make([]*v1Pb.StickerListItem, 0),
				}
				if list, err := c.StickerUseCase.StickerRepo.GetDetailsAll(item.Id, 0); err == nil {
					for _, v := range list {
						data.List = append(data.List, &v1Pb.StickerListItem{
							MediaId: int32(v.Id),
							Src:     v.Url,
						})
					}
				}
				resp.SysSticker = append(resp.SysSticker, data)
			}
		}
	}

	if items, err := c.StickerUseCase.StickerRepo.GetDetailsAll(0, uid); err == nil {
		for _, item := range items {
			resp.CollectSticker = append(resp.CollectSticker, &v1Pb.StickerListItem{
				MediaId: int32(item.Id),
				Src:     item.Url,
			})
		}
	}

	return ctx.Success(resp)
}

func (c *Sticker) DeleteCollect(ctx *core.Context) error {
	params := &v1Pb.StickerDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.StickerUseCase.DeleteCollect(ctx.UserId(), sliceutil.ParseIds(params.Ids)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Sticker) Upload(ctx *core.Context) error {
	file, err := ctx.Context.FormFile("sticker")
	if err != nil {
		return ctx.InvalidParams("Поле 'sticker' обязательно для загрузки!")
	}

	if !sliceutil.Include(strutil.FileSuffix(file.Filename), []string{"png", "jpg", "jpeg", "gif"}) {
		return ctx.InvalidParams("Неверный формат загружаемого файла, поддерживаются только файлы в формате png, jpg, jpeg и gif")
	}

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5 МБ")
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		return ctx.ErrorBusiness("Ошибка загрузки")
	}

	meta := utils.ReadImageMeta(bytes.NewReader(stream))
	ext := strutil.FileSuffix(file.Filename)
	src := fmt.Sprintf("sticker/%s/%s", time.Now().Format("20060102"), strutil.GenImageName(ext, meta.Width, meta.Height))
	if err = c.Filesystem.Default.Write(stream, src); err != nil {
		return ctx.ErrorBusiness("Ошибка загрузки")
	}

	m := &model.StickerItem{
		UserId:      ctx.UserId(),
		Description: "Пользовательский набор",
		Url:         c.Filesystem.Default.PublicUrl(src),
		FileSuffix:  ext,
		FileSize:    int(file.Size),
	}
	if err := c.StickerUseCase.Db().Create(m).Error; err != nil {
		return ctx.ErrorBusiness("Ошибка загрузки")
	}

	return ctx.Success(&v1Pb.StickerUploadResponse{
		MediaId: int32(m.Id),
		Src:     m.Url,
	})
}

func (c *Sticker) SystemList(ctx *core.Context) error {
	items, err := c.StickerUseCase.StickerRepo.GetSystemStickerList(ctx.Ctx())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	ids := c.StickerUseCase.StickerRepo.GetUserInstallIds(ctx.UserId())
	data := make([]*v1Pb.StickerSysListResponse_Item, 0)
	for _, item := range items {
		data = append(data, &v1Pb.StickerSysListResponse_Item{
			Id:     int32(item.Id),
			Name:   item.Name,
			Icon:   item.Icon,
			Status: int32(strutil.BoolToInt(sliceutil.Include(item.Id, ids))),
		})
	}

	return ctx.Success(data)
}

func (c *Sticker) SetSystemSticker(ctx *core.Context) error {
	var (
		err    error
		params = &v1Pb.StickerSetSystemRequest{}
		uid    = ctx.UserId()
		key    = fmt.Sprintf("sys-sticker:%d", uid)
	)
	if err = ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !c.RedisLock.Lock(ctx.Ctx(), key, 5) {
		return ctx.ErrorBusiness("Слишком частые запросы")
	}

	defer c.RedisLock.UnLock(ctx.Ctx(), key)
	if params.Type == 2 {
		if err = c.StickerUseCase.RemoveUserSysSticker(uid, int(params.StickerId)); err != nil {
			return ctx.ErrorBusiness(err.Error())
		}

		return ctx.Success(nil)
	}

	info, err := c.StickerUseCase.StickerRepo.FindById(ctx.Ctx(), int(params.StickerId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}
	if err := c.StickerUseCase.AddUserSysSticker(uid, int(params.StickerId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*v1Pb.StickerListItem, 0)
	if list, err := c.StickerUseCase.StickerRepo.GetDetailsAll(int(params.StickerId), 0); err == nil {
		for _, item := range list {
			items = append(items, &v1Pb.StickerListItem{
				MediaId: int32(item.Id),
				Src:     item.Url,
			})
		}
	}

	return ctx.Success(&v1Pb.StickerSetSystemResponse{
		StickerId: int32(info.Id),
		Url:       info.Icon,
		Name:      info.Name,
		List:      items,
	})
}
