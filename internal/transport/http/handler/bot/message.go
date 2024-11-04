package bot

import (
	botPb "voo.su/api/http/pb/bot"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
)

type Message struct {
	MessageSendUseCase usecase.MessageSendUseCase
	BotUseCase         *usecase.BotUseCase
}

func (m *Message) Send(ctx *core.Context) error {
	params := &botPb.MessageSendRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	token := ctx.Context.Param("token")

	var bot, err = m.BotUseCase.GetBotByToken(ctx.Ctx(), token)
	if err != nil {
		return ctx.ErrorBusiness("")
	}

	if err := m.MessageSendUseCase.SendText(ctx.Ctx(), bot.UserId, &usecase.SendText{
		Receiver: usecase.Receiver{
			DialogType: 2,
			ReceiverId: params.ChatId,
		},
		Content: params.Content,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}

func (m *Message) GroupChats(ctx *core.Context) error {
	token := ctx.Context.Param("token")

	var bot, err = m.BotUseCase.GetBotByToken(ctx.Ctx(), token)
	if err != nil {
		return ctx.ErrorBusiness("")
	}

	list, err := m.BotUseCase.Chats(ctx.Ctx(), bot.CreatorId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*botPb.MessageChatsResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &botPb.MessageChatsResponse_Item{
			Id:   int32(item.Id),
			Name: item.GroupName,
		})
	}

	return ctx.Success(&botPb.MessageChatsResponse{
		Items: items,
	})
}
