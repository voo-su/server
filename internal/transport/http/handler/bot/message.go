package bot

import (
	"fmt"
	botPb "voo.su/api/http/pb/bot"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

type Message struct {
	MessageSendService service.MessageSendService
}

func (m *Message) Send(ctx *core.Context) error {
	params := &botPb.MessageSendRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	token := ctx.Context.Param("token")

	fmt.Println(token)
	fmt.Println(params)

	if err := m.MessageSendService.SendText(ctx.Ctx(), 2, &service.SendText{
		Receiver: service.Receiver{
			DialogType: 2,
			ReceiverId: params.ChatId,
		},
		Content: params.Content,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}
