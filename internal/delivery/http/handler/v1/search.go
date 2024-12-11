package v1

import (
	"errors"
	"gorm.io/gorm"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/infrastructure/postgres/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/timeutil"
)

type Search struct {
	UserUseCase      *usecase.UserUseCase
	GroupChatUseCase *usecase.GroupChatUseCase
}

func (s *Search) Users(ctx *core.Context) error {
	params := &v1Pb.SearchUsersRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	list, err := s.UserUseCase.UserRepo.Search(params.Q, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.ErrorBusiness("Ничего не найдено.")
		}
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*v1Pb.SearchUserResponse_Item, 0)
	for _, item := range list {
		items = append(items, &v1Pb.SearchUserResponse_Item{
			Id:       int32(item.Id),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return ctx.Success(&v1Pb.SearchUserResponse{Items: items})
}

func (s *Search) GroupChats(ctx *core.Context) error {
	params := &v1Pb.SearchGroupChatsRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	list, err := s.GroupChatUseCase.GroupChatRepo.SearchOvertList(ctx.Ctx(), &repository.SearchOvertListOpt{
		Name:   params.Name,
		UserId: uid,
		Page:   int(params.Page),
		Size:   20,
	})

	if err != nil {
		return ctx.ErrorBusiness("Ошибка запроса")
	}
	resp := &v1Pb.SearchGroupChatsResponse{}
	resp.Items = make([]*v1Pb.SearchGroupChatsResponse_Item, 0)
	if len(list) == 0 {
		return ctx.Success(resp)
	}

	ids := make([]int, 0)
	for _, val := range list {
		ids = append(ids, val.Id)
	}

	count, err := s.GroupChatUseCase.MemberRepo.CountGroupMemberNum(ids)
	if err != nil {
		return ctx.ErrorBusiness("Ошибка запроса")
	}

	countMap := make(map[int]int)
	for _, member := range count {
		countMap[member.GroupId] = member.Count
	}

	//checks, err := s.groupChatMemberService.Dao().CheckUserGroup(ids, ctx.UserId())
	//if err != nil {
	//	return ctx.ErrorBusiness("Ошибка запроса")
	//}

	for i, value := range list {
		if i >= 40 {
			break
		}
		resp.Items = append(resp.Items, &v1Pb.SearchGroupChatsResponse_Item{
			Id:          int32(value.Id),
			Type:        int32(value.Type),
			Name:        value.Name,
			Avatar:      value.Avatar,
			Description: value.Description,
			Count:       int32(countMap[value.Id]),
			MaxNum:      int32(value.MaxNum),
			//IsMember:  sliceutil.Include(value.Id, checks),
			CreatedAt: timeutil.FormatDatetime(value.CreatedAt),
		})
	}

	resp.Next = len(list) > 40

	return ctx.Success(resp)
}
