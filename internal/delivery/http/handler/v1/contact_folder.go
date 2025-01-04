package v1

import (
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

type ContactFolder struct {
	Locale               locale.ILocale
	ContactFolderUseCase *usecase.ContactFolderUseCase
}

func (c *ContactFolder) List(ctx *ginutil.Context) error {
	uid := ctx.UserId()
	items := make([]*v1Pb.ContactFolderListResponse_Item, 0)
	count, err := c.ContactFolderUseCase.ContactRepo.QueryCount(ctx.Ctx(), "user_id = ? AND status = 1", uid)
	if err != nil {
		return ctx.Error(err)
	}

	items = append(items, &v1Pb.ContactFolderListResponse_Item{
		Name:  "Все",
		Count: int32(count),
	})

	group, err := c.ContactFolderUseCase.GetUserGroup(ctx.Ctx(), uid)
	if err != nil {
		return ctx.Error(err)
	}

	for _, v := range group {
		items = append(items, &v1Pb.ContactFolderListResponse_Item{
			Id:    int32(v.Id),
			Name:  v.Name,
			Count: int32(v.Num),
			Sort:  int32(v.Sort),
		})
	}

	return ctx.Success(&v1Pb.ContactFolderListResponse{Items: items})
}

func (c *ContactFolder) Save(ctx *ginutil.Context) error {
	params := &v1Pb.ContactFolderSaveRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	updateItems := make([]*model.ContactFolder, 0)
	deleteItems := make([]int, 0)
	insertItems := make([]*model.ContactFolder, 0)

	ids := make(map[int]struct{})
	for i, item := range params.GetItems() {
		if item.Id > 0 {
			ids[int(item.Id)] = struct{}{}
			updateItems = append(updateItems, &model.ContactFolder{
				Id:   int(item.Id),
				Sort: i + 1,
				Name: item.Name,
			})
		} else {
			insertItems = append(insertItems, &model.ContactFolder{
				Sort:   i + 1,
				Name:   item.Name,
				UserId: uid,
			})
		}
	}

	all, err := c.ContactFolderUseCase.ContactFolderRepo.FindAll(ctx.Ctx())
	if err != nil {
		return ctx.Error(err)
	}

	for _, m := range all {
		if _, ok := ids[m.Id]; !ok {
			deleteItems = append(deleteItems, m.Id)
		}
	}

	err = c.ContactFolderUseCase.Source.Postgres().Transaction(func(tx *gorm.DB) error {
		if len(insertItems) > 0 {
			if err := tx.Create(insertItems).Error; err != nil {
				return err
			}
		}

		if len(deleteItems) > 0 {
			if err := tx.Delete(model.ContactFolder{}, "id in (?) AND user_id = ?", deleteItems, uid).Error; err != nil {
				return err
			}
			tx.Table("contacts").
				Where("user_id = ? AND group_id in (?)", uid, deleteItems).
				UpdateColumn("group_id", 0)
		}

		for _, item := range updateItems {
			if err = tx.Table("contact_folders").
				Where("id = ? AND user_id = ?", item.Id, uid).
				Updates(map[string]any{
					"name": item.Name,
					"sort": item.Sort,
				}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Success(&v1Pb.ContactFolderSaveResponse{})
}

func (c *ContactFolder) Move(ctx *ginutil.Context) error {
	params := &v1Pb.ContactChangeGroupRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.ContactFolderUseCase.MoveGroup(ctx.Ctx(), ctx.UserId(), int(params.UserId), int(params.FolderId)); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ContactChangeGroupResponse{})
}
