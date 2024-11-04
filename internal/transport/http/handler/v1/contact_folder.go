package v1

import (
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/repository/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
)

type ContactFolder struct {
	ContactFolderUseCase *usecase.ContactFolderUseCase
}

func (c *ContactFolder) List(ctx *core.Context) error {
	uid := ctx.UserId()
	items := make([]*v1Pb.ContactFolderListResponse_Item, 0)
	count, err := c.ContactFolderUseCase.ContactRepo.QueryCount(ctx.Ctx(), "user_id = ? and status = 1", uid)
	if err != nil {
		return ctx.Error(err.Error())
	}
	items = append(items, &v1Pb.ContactFolderListResponse_Item{
		Name:  "Все",
		Count: int32(count),
	})
	group, err := c.ContactFolderUseCase.GetUserGroup(ctx.Ctx(), uid)
	if err != nil {
		return ctx.Error(err.Error())
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

func (c *ContactFolder) Save(ctx *core.Context) error {
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
		return ctx.ErrorBusiness(err)
	}

	for _, m := range all {
		if _, ok := ids[m.Id]; !ok {
			deleteItems = append(deleteItems, m.Id)
		}
	}

	err = c.ContactFolderUseCase.Db().Transaction(func(tx *gorm.DB) error {
		if len(insertItems) > 0 {
			if err := tx.Create(insertItems).Error; err != nil {
				return err
			}
		}

		if len(deleteItems) > 0 {
			err := tx.Delete(model.ContactFolder{}, "id in (?) and user_id = ?", deleteItems, uid).Error
			if err != nil {
				return err
			}
			tx.Table("contacts").
				Where("user_id = ? and group_id in (?)", uid, deleteItems).
				UpdateColumn("group_id", 0)
		}

		for _, item := range updateItems {
			err = tx.Table("contact_folders").
				Where("id = ? and user_id = ?", item.Id, uid).
				Updates(map[string]any{
					"name": item.Name,
					"sort": item.Sort,
				}).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.ContactFolderSaveResponse{})
}

func (c *ContactFolder) Move(ctx *core.Context) error {
	params := &v1Pb.ContactChangeGroupRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.ContactFolderUseCase.MoveGroup(ctx.Ctx(), ctx.UserId(), int(params.UserId), int(params.FolderId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.ContactChangeGroupResponse{})
}
