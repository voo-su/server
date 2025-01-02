// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package v1

import (
	"fmt"
	"regexp"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/locale"
)

type Bot struct {
	Locale         locale.ILocale
	BotUseCase     *usecase.BotUseCase
	MessageUseCase usecase.IMessageUseCase
}

func (b *Bot) Create(ctx *core.Context) error {
	params := &v1Pb.BotCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	re := regexp.MustCompile(`(_bot|bot)$`)
	if !re.MatchString(params.Username) {
		return ctx.ErrorBusiness(fmt.Sprintf(b.Locale.Localize("username_does_not_end_with_bot"), params.Username))
	}

	token, err := b.BotUseCase.Create(ctx.Ctx(), &usecase.BotCreateOpt{
		Username:  params.Username,
		CreatorId: ctx.UserId(),
	})
	if err != nil {
		return nil
	}

	return ctx.Success(&v1Pb.BotCreateResponse{
		Token: *token,
	})
}

func (b *Bot) List(ctx *core.Context) error {
	list, err := b.BotUseCase.List(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*v1Pb.BotListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &v1Pb.BotListResponse_Item{
			Id:       int32(item.Id),
			Username: item.Name,
			Token:    item.Token,
		})
	}

	return ctx.Success(&v1Pb.BotListResponse{
		Items: items,
	})
}
