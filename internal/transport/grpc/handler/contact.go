package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/service"
	"voo.su/internal/transport/grpc/middleware"
)

type ContactHandler struct {
	contactPb.UnimplementedContactServiceServer
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	ContactService  *service.ContactService
}

func NewContactHandler(
	conf *config.Config,
	tokenMiddleware *middleware.TokenMiddleware,
	contactService *service.ContactService,
) *ContactHandler {
	return &ContactHandler{
		Conf:            conf,
		TokenMiddleware: tokenMiddleware,
		ContactService:  contactService,
	}
}

func (c *ContactHandler) List(ctx context.Context, in *contactPb.GetContactListRequest) (*contactPb.GetContactListResponse, error) {

	// TODO
	uid := 1

	list, err := c.ContactService.List(ctx, uid)
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
