package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/transport/grpc/middleware"
	"voo.su/internal/usecase"
)

type ContactHandler struct {
	contactPb.UnimplementedContactServiceServer
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	ContactUseCase  *usecase.ContactUseCase
}

func NewContactHandler(
	conf *config.Config,
	tokenMiddleware *middleware.TokenMiddleware,
	contactUseCase *usecase.ContactUseCase,
) *ContactHandler {
	return &ContactHandler{
		Conf:            conf,
		TokenMiddleware: tokenMiddleware,
		ContactUseCase:  contactUseCase,
	}
}

func (c *ContactHandler) List(ctx context.Context, in *contactPb.GetContactListRequest) (*contactPb.GetContactListResponse, error) {

	// TODO
	uid := 1

	list, err := c.ContactUseCase.List(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*contactPb.ContactItem, 0)
	for _, item := range list {
		items = append(items, &contactPb.ContactItem{
			Id:       int32(item.Id),
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return &contactPb.GetContactListResponse{
		Items: items,
	}, nil
}
