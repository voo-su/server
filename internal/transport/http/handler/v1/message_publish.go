package v1

import (
	"github.com/gin-gonic/gin/binding"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

var mapping map[string]func(ctx *core.Context) error

type Publish struct {
	AuthService        *service.AuthService
	MessageSendService service.MessageSendService
}

func (c *Publish) transfer(ctx *core.Context, typeValue string) error {
	if mapping == nil {
		mapping = make(map[string]func(ctx *core.Context) error)
		mapping["text"] = c.onSendText
		mapping["image"] = c.onSendImage
		mapping["vote"] = c.onSendVote
		mapping["voice"] = c.onSendVoice
		mapping["video"] = c.onSendVideo
		mapping["file"] = c.onSendFile
		mapping["forward"] = c.onSendForward
		mapping["mixed"] = c.onMixedMessage
		mapping["sticker"] = c.onSendSticker
		mapping["code"] = c.onSendCode
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

func (c *Publish) Publish(ctx *core.Context) error {
	params := &PublishBaseMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType:        params.Receiver.DialogType,
		UserId:            ctx.UserId(),
		ReceiverId:        params.Receiver.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return c.transfer(ctx, params.Type)
}

func (c *Publish) onSendText(ctx *core.Context) error {
	params := &v1Pb.TextMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendText(ctx.Ctx(), ctx.UserId(), &service.SendText{
		Receiver: service.Receiver{
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

func (c *Publish) onSendImage(ctx *core.Context) error {
	params := &v1Pb.ImageMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendImage(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVoice(ctx *core.Context) error {

	params := &v1Pb.VoiceMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendVoice(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVideo(ctx *core.Context) error {

	params := &v1Pb.VideoMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendVideo(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendFile(ctx *core.Context) error {
	params := &v1Pb.FileMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendFile(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendForward(ctx *core.Context) error {
	params := &v1Pb.ForwardMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendForward(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVote(ctx *core.Context) error {
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

	if err := c.MessageSendService.SendVote(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onMixedMessage(ctx *core.Context) error {
	params := &v1Pb.MixedMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendMixedMessage(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendSticker(ctx *core.Context) error {
	params := &v1Pb.StickerMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendSticker(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendCode(ctx *core.Context) error {
	params := &v1Pb.CodeMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.SendCode(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}
