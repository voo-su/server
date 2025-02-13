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
}

func NewContactHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
) *Contact {
	return &Contact{
		Conf:           conf,
		Locale:         locale,
		ContactUseCase: contactUseCase,
	}
}

func (c *Contact) List(ctx context.Context, in *contactPb.GetContactListRequest) (*contactPb.GetContactListResponse, error) {
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

	return &contactPb.GetContactListResponse{
		Items: items,
	}, nil
}
