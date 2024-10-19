package v1

import (
	"errors"
	"gorm.io/gorm"
	"voo.su/api/pb/v1"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/core"
)

type Search struct {
	UserRepo *repo.User
}

func (s *Search) Search(ctx *core.Context) error {
	params := &api_v1.SearchUserRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	list, err := s.UserRepo.SearchByUsername(params.Username, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.ErrorBusiness("Ничего не найдено.")
		}
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*api_v1.SearchUserResponse_Item, 0)
	for _, item := range list {
		items = append(items, &api_v1.SearchUserResponse_Item{
			Id:       int32(item.Id),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return ctx.Success(&api_v1.SearchUserResponse{Items: items})
}
