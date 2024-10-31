package v1

import (
	"fmt"
	"regexp"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type Bot struct {
	BotService         *service.BotService
	MessageSendService service.MessageSendService
}

func (b *Bot) Create(ctx *core.Context) error {
	params := &v1Pb.BotCreateRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	re := regexp.MustCompile(`(_bot|bot)$`)
	if !re.MatchString(params.Username) {
		return ctx.ErrorBusiness(fmt.Sprintf("Имя пользователя '%s' не заканчивается на '_bot' или 'bot'", params.Username))
	}

	token, err := b.BotService.Create(ctx.Ctx(), &service.BotCreateOpt{
		Username:  params.Username,
		CreatorId: ctx.UserId(),
	})
	if err != nil {
		return nil
	}

	//_ = b.MessageSendService.SendLogin(ctx.Ctx(), ctx.UserId(), &v1Pb.LoginMessageRequest{})

	return ctx.Success(&v1Pb.BotCreateResponse{
		Token: *token,
	})
}

func (b *Bot) List(ctx *core.Context) error {
	list, err := b.BotService.List(ctx.Ctx(), ctx.UserId())
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
