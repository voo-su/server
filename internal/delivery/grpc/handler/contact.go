// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc/middleware"
	"voo.su/internal/usecase"
	"voo.su/pkg/locale"
)

type Contact struct {
	contactPb.UnimplementedContactServiceServer
	Conf            *config.Config
	Locale          locale.ILocale
	TokenMiddleware *middleware.TokenMiddleware
	ContactUseCase  *usecase.ContactUseCase
}

func NewContactHandler(
	conf *config.Config,
	locale locale.ILocale,
	tokenMiddleware *middleware.TokenMiddleware,
	contactUseCase *usecase.ContactUseCase,
) *Contact {
	return &Contact{
		Conf:            conf,
		Locale:          locale,
		TokenMiddleware: tokenMiddleware,
		ContactUseCase:  contactUseCase,
	}
}

func (c *Contact) List(ctx context.Context, in *contactPb.GetContactListRequest) (*contactPb.GetContactListResponse, error) {

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
