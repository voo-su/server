package v1

import (
	"github.com/gin-gonic/gin/binding"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
)

var mapping map[string]func(ctx *core.Context) error

type Publish struct {
	AuthUseCase    *usecase.AuthUseCase
	MessageUseCase usecase.IMessageUseCase
}

func (p *Publish) transfer(ctx *core.Context, typeValue string) error {
	if mapping == nil {
		mapping = make(map[string]func(ctx *core.Context) error)
		mapping["text"] = p.onSendText
		mapping["image"] = p.onSendImage
		mapping["vote"] = p.onSendVote
		mapping["voice"] = p.onSendVoice
		mapping["video"] = p.onSendVideo
		mapping["file"] = p.onSendFile
		mapping["forward"] = p.onSendForward
		mapping["mixed"] = p.onMixedMessage
		mapping["sticker"] = p.onSendSticker
		mapping["code"] = p.onSendCode
	}
	if call, ok := mapping[typeValue]; ok {
		return call(ctx)
	}

	return nil
}

type PublishBaseMessageRequest struct {
	Type     string `json:"type" binding:"required"`
	Receiver struct {
		DialogType int `json:"dialog_type" binding:"required,gt=0"`
		ReceiverId int `json:"receiver_id" binding:"required,gt=0"`
	} `json:"receiver" binding:"required"`
}

func (p *Publish) Publish(ctx *core.Context) error {
	params := &PublishBaseMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.AuthUseCase.IsAuth(ctx.Ctx(), &usecase.AuthOption{
		DialogType:        params.Receiver.DialogType,
		UserId:            ctx.UserId(),
		ReceiverId:        params.Receiver.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return p.transfer(ctx, params.Type)
}

func (p *Publish) onSendText(ctx *core.Context) error {
	params := &v1Pb.TextMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendText(ctx.Ctx(), ctx.UserId(), &usecase.SendText{
		Receiver: usecase.Receiver{
			DialogType: params.Receiver.DialogType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Content: params.Content,
		QuoteId: params.QuoteId,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendImage(ctx *core.Context) error {
	params := &v1Pb.ImageMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendImage(ctx.Ctx(), ctx.UserId(), &usecase.SendImage{
		Receiver: usecase.Receiver{
			DialogType: params.Receiver.DialogType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Url:     params.Url,
		Width:   params.Width,
		Height:  params.Height,
		QuoteId: params.QuoteId,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendVoice(ctx *core.Context) error {
	params := &v1Pb.VoiceMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendVoice(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendVideo(ctx *core.Context) error {
	params := &v1Pb.VideoMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendVideo(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendFile(ctx *core.Context) error {
	params := &v1Pb.FileMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendFile(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendForward(ctx *core.Context) error {
	params := &v1Pb.ForwardMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendForward(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendVote(ctx *core.Context) error {
	params := &v1Pb.VoteMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if len(params.Options) <= 1 {
		return ctx.InvalidParams("количество вариантов (options) должно быть больше 1!")
	}

	if len(params.Options) > 6 {
		return ctx.InvalidParams("количество вариантов (options) не может превышать 6!")
	}

	if err := p.MessageUseCase.SendVote(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onMixedMessage(ctx *core.Context) error {
	params := &v1Pb.MixedMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendMixedMessage(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendSticker(ctx *core.Context) error {
	params := &v1Pb.StickerMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendSticker(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (p *Publish) onSendCode(ctx *core.Context) error {
	params := &v1Pb.CodeMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := p.MessageUseCase.SendCode(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}
