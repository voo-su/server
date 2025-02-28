package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
)

type Contact struct {
	contactPb.UnimplementedContactServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	ContactUseCase *usecase.ContactUseCase
	UserUseCase    *usecase.UserUseCase
}

func NewContactHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
	userUseCase *usecase.UserUseCase,
) *Contact {
	return &Contact{
		Conf:           conf,
		Locale:         locale,
		ContactUseCase: contactUseCase,
		UserUseCase:    userUseCase,
	}
}

func (c *Contact) GetContacts(ctx context.Context, in *contactPb.GetContactsRequest) (*contactPb.GetContactsResponse, error) {
	uid := grpcutil.UserId(ctx)

	list, err := c.ContactUseCase.List(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*contactPb.ContactItem, 0)
	for _, item := range list {
		items = append(items, &contactPb.ContactItem{
			Id:       int64(item.Id),
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return &contactPb.GetContactsResponse{
		Items: items,
	}, nil
}

func (c *Contact) GetUser(ctx context.Context, in *contactPb.GetUserRequest) (*contactPb.GetUserResponse, error) {
	user, err := c.UserUseCase.UserRepo.FindById(ctx, int(in.Id))
	if err != nil {
		return nil, nil
	}

	return &contactPb.GetUserResponse{
		Id:       int64(user.Id),
		Username: user.Username,
		Avatar:   user.Avatar,
		Name:     user.Name,
		Surname:  user.Surname,
		Gender:   int32(user.Gender),
		About:    user.About,
	}, nil
}

func (c *Contact) Search(ctx context.Context, in *contactPb.SearchRequest) (*contactPb.SearchResponse, error) {
	uid := grpcutil.UserId(ctx)

	if len(in.Q) <= 1 {
		return nil, nil
	}

	list, err := c.UserUseCase.UserRepo.Search(in.Q, uid, int(in.Id))
	if err != nil {
		return nil, nil
	}

	items := make([]*contactPb.ContactItem, 0)
	for _, item := range list {
		items = append(items, &contactPb.ContactItem{
			Id:       int64(item.Id),
			Username: item.Username,
			Avatar:   item.Avatar,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}
	return &contactPb.SearchResponse{
		Items: items,
	}, nil
}
