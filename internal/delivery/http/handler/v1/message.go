package v1

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type Message struct {
	Conf                   *config.Config
	Locale                 locale.ILocale
	ChatUseCase            *usecase.ChatUseCase
	MessageUseCase         usecase.IMessageUseCase
	GroupChatMemberUseCase *usecase.GroupChatMemberUseCase
	StorageUseCase         *usecase.StorageUseCase
}

var mapping map[string]func(ctx *ginutil.Context) error

func (m *Message) transfer(ctx *ginutil.Context, typeValue string) error {
	if mapping == nil {
		mapping = make(map[string]func(ctx *ginutil.Context) error)
		mapping["text"] = m.onSendText
		mapping["image"] = m.onSendImage
		mapping["voice"] = m.onSendAudio
		mapping["video"] = m.onSendVideo
		mapping["file"] = m.onSendFile
		mapping["vote"] = m.onSendVote
		mapping["forward"] = m.onSendForward
		mapping["mixed"] = m.onMixedMessage
		mapping["sticker"] = m.onSendSticker
		mapping["code"] = m.onSendCode
	}
	if call, ok := mapping[typeValue]; ok {
		return call(ctx)
	}

	return nil
}

func (m *Message) Send(ctx *ginutil.Context) error {
	params := &v1Pb.PublishBaseMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.IsAccess(ctx.Ctx(), &entity.MessageAccess{
		ChatType:          int(params.Receiver.ChatType),
		UserId:            ctx.UserId(),
		ReceiverId:        int(params.Receiver.ReceiverId),
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return m.transfer(ctx, params.Type)
}

func (m *Message) onSendText(ctx *ginutil.Context) error {
	params := &v1Pb.TextMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendText(ctx.Ctx(), ctx.UserId(), &entity.SendText{
		Receiver: entity.MessageReceiver{
			ChatType:   params.Receiver.ChatType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Content: params.Content,
		QuoteId: params.QuoteId,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendImage(ctx *ginutil.Context) error {
	params := &v1Pb.ImageMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendImage(ctx.Ctx(), ctx.UserId(), &entity.SendImage{
		Receiver: entity.MessageReceiver{
			ChatType:   params.Receiver.ChatType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Url:     params.Url,
		Width:   params.Width,
		Height:  params.Height,
		QuoteId: params.QuoteId,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendVideo(ctx *ginutil.Context) error {
	params := &v1Pb.VideoMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendVideo(ctx.Ctx(), ctx.UserId(), &entity.SendVideo{
		Receiver: entity.MessageReceiver{
			ChatType:   params.Receiver.ChatType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Url:      params.Url,
		Duration: params.Duration,
		Size:     params.Size,
		Cover:    params.Cover,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendAudio(ctx *ginutil.Context) error {
	params := &v1Pb.AudioMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendAudio(ctx.Ctx(), ctx.UserId(), &entity.SendAudio{
		Receiver: entity.MessageReceiver{
			ChatType:   params.Receiver.ChatType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		Url:  params.Url,
		Size: params.Size,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendFile(ctx *ginutil.Context) error {
	params := &v1Pb.FileMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendFile(ctx.Ctx(), ctx.UserId(), &entity.SendFile{
		Receiver: entity.MessageReceiver{
			ChatType:   params.Receiver.ChatType,
			ReceiverId: params.Receiver.ReceiverId,
		},
		UploadId: params.UploadId,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendForward(ctx *ginutil.Context) error {
	params := &v1Pb.ForwardMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendForward(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendVote(ctx *ginutil.Context) error {
	params := &v1Pb.VoteMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if len(params.Options) <= 1 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("options_count_must_be_greater_than_one"), 1))
	}

	if len(params.Options) > 6 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("options_count_cannot_exceed_ten"), 6))
	}

	if err := m.MessageUseCase.SendVote(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onMixedMessage(ctx *ginutil.Context) error {
	params := &v1Pb.MixedMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendMixedMessage(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendSticker(ctx *ginutil.Context) error {
	params := &v1Pb.StickerMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendSticker(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) onSendCode(ctx *ginutil.Context) error {
	params := &v1Pb.CodeMessageRequest{}
	if err := ctx.Context.ShouldBindBodyWith(params, binding.JSON); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.SendCode(ctx.Ctx(), ctx.UserId(), params); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Vote(ctx *ginutil.Context) error {
	params := &v1Pb.VoteSendMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if len(params.Options) <= 1 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("options_count_must_be_greater_than_one"), 1))
	}

	if len(params.Options) > 10 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("options_count_cannot_exceed_ten"), 10))
	}

	uid := ctx.UserId()
	if err := m.MessageUseCase.IsAccess(ctx.Ctx(), &entity.MessageAccess{
		ChatType:          constant.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        int(params.ReceiverId),
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	if err := m.MessageUseCase.SendVote(ctx.Ctx(), uid, &v1Pb.VoteMessageRequest{
		Mode:      params.Mode,
		Title:     params.Title,
		Options:   params.Options,
		Anonymous: params.Anonymous,
		Receiver: &v1Pb.MessageReceiver{
			ChatType:   constant.ChatGroupMode,
			ReceiverId: int32(params.ReceiverId),
		},
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Revoke(ctx *ginutil.Context) error {
	params := &v1Pb.RevokeMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.Revoke(ctx.Ctx(), ctx.UserId(), params.MsgId); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Delete(ctx *ginutil.Context) error {
	params := &v1Pb.DeleteMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.ChatUseCase.DeleteRecordList(ctx.Ctx(), &usecase.RemoveRecordListOpt{
		UserId:     ctx.UserId(),
		ChatType:   int(params.ChatType),
		ReceiverId: int(params.ReceiverId),
		MsgIds:     sliceutil.ParseIdsToInt64(params.MsgIds),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) HandleVote(ctx *ginutil.Context) error {
	params := &v1Pb.VoteMessageHandleRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := m.MessageUseCase.Vote(ctx.Ctx(), ctx.UserId(), params.RecordId, params.Options)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(data)
}

func (m *Message) GetHistory(ctx *ginutil.Context) error {
	params := &v1Pb.GetRecordsRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if params.ChatType == constant.ChatGroupMode {
		if err := m.MessageUseCase.IsAccess(ctx.Ctx(), &entity.MessageAccess{
			ChatType:   int(params.ChatType),
			UserId:     uid,
			ReceiverId: int(params.ReceiverId),
		}); err != nil {
			items := make([]map[string]any, 0)
			items = append(items, map[string]any{
				"content":     m.Locale.Localize("insufficient_permissions_to_view_messages"),
				"created_at":  timeutil.DateTime(),
				"id":          1,
				"msg_id":      strutil.NewMsgId(),
				"msg_type":    constant.ChatMsgSysText,
				"receiver_id": params.ReceiverId,
				"chat_type":   params.ChatType,
				"user_id":     0,
			})

			return ctx.Success(map[string]any{
				"limit":     params.Limit,
				"record_id": 0,
				"items":     items,
			})
		}
	}

	records, err := m.MessageUseCase.GetHistory(ctx.Ctx(), &entity.QueryGetHistoryOpt{
		ChatType:   int(params.ChatType),
		UserId:     uid,
		ReceiverId: int(params.ReceiverId),
		RecordId:   int(params.RecordId),
		Limit:      int(params.Limit),
	})
	if err != nil {
		return ctx.Error(err.Error())
	}

	rid := 0
	if length := len(records); length > 0 {
		rid = records[length-1].Sequence
	}

	return ctx.Success(map[string]any{
		"limit":     params.Limit,
		"record_id": rid,
		"items":     records,
	})
}

func (m *Message) Download(ctx *ginutil.Context) error {
	params := &v1Pb.DownloadChatFileRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	record, err := m.MessageUseCase.GetMessageByRecordId(ctx.Ctx(), params.RecordId)
	if err != nil {
		return ctx.Error(err.Error())
	}

	uid := ctx.UserId()
	if uid != record.UserId {
		if record.ChatType == constant.ChatPrivateMode {
			if record.ReceiverId != uid {
				return ctx.Forbidden(m.Locale.Localize("access_rights_missing"))
			}
		} else {
			if !m.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), record.ReceiverId, uid, false) {
				return ctx.Forbidden(m.Locale.Localize("access_rights_missing"))
			}
		}
	}

	var fileInfo entity.MessageExtraFile
	if err := jsonutil.Decode(record.Extra, &fileInfo); err != nil {
		return ctx.Error(err.Error())
	}

	object, err := m.StorageUseCase.Minio.GetObject(m.Conf.Minio.GetBucket(), fileInfo.Path)
	if err != nil {
		ctx.Context.AbortWithStatusJSON(http.StatusNotFound, &ginutil.Response{
			Code:    http.StatusNotFound,
			Message: m.Locale.Localize("file_not_found"),
		})

		return nil
	}

	defer object.Close()

	stat, err := m.StorageUseCase.Minio.Stat(m.Conf.Minio.GetBucket(), fileInfo.Path)
	if err != nil {
		ctx.Context.AbortWithStatusJSON(http.StatusInternalServerError, &ginutil.Response{
			Code:    http.StatusInternalServerError,
			Message: m.Locale.Localize("unable_to_get_file_information"),
		})

		return nil
	}

	ctx.Context.Header("Content-Type", stat.MimeType)
	ctx.Context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name))
	ctx.Context.Header("Content-Length", fmt.Sprintf("%d", stat.Size))

	_, err = io.Copy(ctx.Context.Writer, object)
	if err != nil {
		ctx.Context.AbortWithStatusJSON(http.StatusInternalServerError, &ginutil.Response{
			Code:    http.StatusInternalServerError,
			Message: m.Locale.Localize("failed_to_send_file"),
		})
	}

	return nil
}

func (m *Message) Collect(ctx *ginutil.Context) error {
	params := &v1Pb.CollectMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.ChatUseCase.Collect(ctx.Ctx(), ctx.UserId(), params.RecordId); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}
