package v1

import (
	"github.com/gin-gonic/gin/binding"
	"voo.su/api/pb/v1"
	"voo.su/internal/service"
	"voo.su/pkg/core"
)

var mapping map[string]func(ctx *core.Context) error

type Publish struct {
	AuthService        *service.AuthService
	MessageSendService service.MessageSendService
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
	params := &api_v1.TextMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendText(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendImage(ctx *core.Context) error {
	params := &api_v1.ImageMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendImage(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVoice(ctx *core.Context) error {

	params := &api_v1.VoiceMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendVoice(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVideo(ctx *core.Context) error {

	params := &api_v1.VideoMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendVideo(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendFile(ctx *core.Context) error {
	params := &api_v1.FileMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendFile(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendCode(ctx *core.Context) error {
	params := &api_v1.CodeMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendCode(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendLocation(ctx *core.Context) error {
	params := &api_v1.LocationMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendLocation(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendForward(ctx *core.Context) error {
	params := &api_v1.ForwardMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendForward(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendVote(ctx *core.Context) error {
	params := &api_v1.VoteMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if len(params.Options) <= 1 {
		return ctx.InvalidParams("количество вариантов (options) должно быть больше 1!")
	}

	if len(params.Options) > 6 {
		return ctx.InvalidParams("количество вариантов (options) не может превышать 6!")
	}

	err := c.MessageSendService.SendVote(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onSendCard(ctx *core.Context) error {
	params := &api_v1.CardMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendBusinessCard(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) onMixedMessage(ctx *core.Context) error {
	params := &api_v1.MixedMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	err := c.MessageSendService.SendMixedMessage(ctx.Ctx(), ctx.UserId(), params)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (c *Publish) transfer(ctx *core.Context, typeValue string) error {
	if mapping == nil {
		mapping = make(map[string]func(ctx *core.Context) error)
		mapping["text"] = c.onSendText
		mapping["code"] = c.onSendCode
		mapping["location"] = c.onSendLocation
		mapping["vote"] = c.onSendVote
		mapping["image"] = c.onSendImage
		mapping["voice"] = c.onSendVoice
		mapping["video"] = c.onSendVideo
		mapping["file"] = c.onSendFile
		mapping["card"] = c.onSendCard
		mapping["forward"] = c.onSendForward
		mapping["mixed"] = c.onMixedMessage
	}
	if call, ok := mapping[typeValue]; ok {
		return call(ctx)
	}

	return nil
}
