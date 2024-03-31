package v1

import (
	"bytes"
	"fmt"
	"time"
	"voo.su/api/pb/v1"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/utils"
)

type Sticker struct {
	StickerRepo    *repo.Sticker
	Filesystem     *filesystem.Filesystem
	StickerService *service.StickerService
	RedisLock      *cache.RedisLock
}

func (c *Sticker) CollectList(ctx *core.Context) error {
	var (
		uid  = ctx.UserId()
		resp = &api_v1.StickerListResponse{
			SysSticker:     make([]*api_v1.StickerListResponse_SysSticker, 0),
			CollectSticker: make([]*api_v1.StickerListItem, 0),
		}
	)
	if ids := c.StickerRepo.GetUserInstallIds(uid); len(ids) > 0 {
		if items, err := c.StickerRepo.FindByIds(ctx.Ctx(), ids); err == nil {
			for _, item := range items {
				data := &api_v1.StickerListResponse_SysSticker{
					StickerId: int32(item.Id),
					Url:       item.Icon,
					Name:      item.Name,
					List:      make([]*api_v1.StickerListItem, 0),
				}
				if list, err := c.StickerRepo.GetDetailsAll(item.Id, 0); err == nil {
					for _, v := range list {
						data.List = append(data.List, &api_v1.StickerListItem{
							MediaId: int32(v.Id),
							Src:     v.Url,
						})
					}
				}
				resp.SysSticker = append(resp.SysSticker, data)
			}
		}
	}

	if items, err := c.StickerRepo.GetDetailsAll(0, uid); err == nil {
		for _, item := range items {
			resp.CollectSticker = append(resp.CollectSticker, &api_v1.StickerListItem{
				MediaId: int32(item.Id),
				Src:     item.Url,
			})
		}
	}

	return ctx.Success(resp)
}

func (c *Sticker) DeleteCollect(ctx *core.Context) error {
	params := &api_v1.StickerDeleteRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.StickerService.DeleteCollect(ctx.UserId(), sliceutil.ParseIds(params.Ids)); err != nil {
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
		Description: "Пользовательский набор смайликов",
		Url:         c.Filesystem.Default.PublicUrl(src),
		FileSuffix:  ext,
		FileSize:    int(file.Size),
	}
	if err := c.StickerService.Db().Create(m).Error; err != nil {
		return ctx.ErrorBusiness("Ошибка загрузки")
	}

	return ctx.Success(&api_v1.StickerUploadResponse{
		MediaId: int32(m.Id),
		Src:     m.Url,
	})
}

func (c *Sticker) SystemList(ctx *core.Context) error {
	items, err := c.StickerRepo.GetSystemStickerList(ctx.Ctx())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	ids := c.StickerRepo.GetUserInstallIds(ctx.UserId())
	data := make([]*api_v1.StickerSysListResponse_Item, 0)
	for _, item := range items {
		data = append(data, &api_v1.StickerSysListResponse_Item{
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
		params = &api_v1.StickerSetSystemRequest{}
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
		if err = c.StickerService.RemoveUserSysSticker(uid, int(params.StickerId)); err != nil {
			return ctx.ErrorBusiness(err.Error())
		}

		return ctx.Success(nil)
	}

	info, err := c.StickerRepo.FindById(ctx.Ctx(), int(params.StickerId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}
	if err := c.StickerService.AddUserSysSticker(uid, int(params.StickerId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*api_v1.StickerListItem, 0)
	if list, err := c.StickerRepo.GetDetailsAll(int(params.StickerId), 0); err == nil {
		for _, item := range list {
			items = append(items, &api_v1.StickerListItem{
				MediaId: int32(item.Id),
				Src:     item.Url,
			})
		}
	}

	return ctx.Success(&api_v1.StickerSetSystemResponse{
		StickerId: int32(info.Id),
		Url:       info.Icon,
		Name:      info.Name,
		List:      items,
	})
}
