package v1

import (
	"net/http"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/minio"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type Message struct {
	ChatUseCase            *usecase.ChatUseCase
	AuthUseCase            *usecase.AuthUseCase
	MessageUseCase         usecase.IMessageUseCase
	Minio                  minio.IMinio
	GroupChatMemberUseCase *usecase.GroupChatMemberUseCase
}

type VoteMessageRequest struct {
	ReceiverId int      `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	Mode       int      `form:"mode" json:"mode" binding:"oneof=0 1"`
	Anonymous  int      `form:"anonymous" json:"anonymous" binding:"oneof=0 1"`
	Title      string   `form:"title" json:"title" binding:"required"`
	Options    []string `form:"options" json:"options"`
}

func (m *Message) Vote(ctx *core.Context) error {
	params := &VoteMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if len(params.Options) <= 1 {
		return ctx.InvalidParams("количество вариантов должно быть больше 1")
	}

	if len(params.Options) > 10 {
		return ctx.InvalidParams("количество вариантов не может превышать 10")
	}

	uid := ctx.UserId()
	if err := m.AuthUseCase.IsAuth(ctx.Ctx(), &usecase.AuthOption{
		DialogType:        constant.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if err := m.MessageUseCase.SendVote(ctx.Ctx(), uid, &v1Pb.VoteMessageRequest{
		Mode:      int32(params.Mode),
		Title:     params.Title,
		Options:   params.Options,
		Anonymous: int32(params.Anonymous),
		Receiver: &v1Pb.MessageReceiver{
			DialogType: constant.ChatGroupMode,
			ReceiverId: int32(params.ReceiverId),
		},
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type RevokeMessageRequest struct {
	MsgId string `form:"msg_id" json:"msg_id" binding:"required" label:"msg_id"`
}

func (m *Message) Revoke(ctx *core.Context) error {
	params := &RevokeMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.MessageUseCase.Revoke(ctx.Ctx(), ctx.UserId(), params.MsgId); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type DeleteMessageRequest struct {
	DialogType int    `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
	ReceiverId int    `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	RecordIds  string `form:"record_id" json:"record_id" binding:"required,ids" label:"record_id"`
}

func (m *Message) Delete(ctx *core.Context) error {
	params := &DeleteMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.ChatUseCase.DeleteRecordList(ctx.Ctx(), &usecase.RemoveRecordListOpt{
		UserId:     ctx.UserId(),
		DialogType: params.DialogType,
		ReceiverId: params.ReceiverId,
		RecordIds:  params.RecordIds,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type VoteMessageHandleRequest struct {
	RecordId int    `form:"record_id" json:"record_id" binding:"required,gt=0"`
	Options  string `form:"options" json:"options" binding:"required"`
}

func (m *Message) HandleVote(ctx *core.Context) error {
	params := &VoteMessageHandleRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := m.MessageUseCase.Vote(ctx.Ctx(), ctx.UserId(), params.RecordId, params.Options)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(data)
}

type GetDialogRecordsRequest struct {
	DialogType int `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2"`
	MsgType    int `form:"msg_type" json:"msg_type" binding:"numeric"`
	ReceiverId int `form:"receiver_id" json:"receiver_id" binding:"required,numeric,min=1"`
	RecordId   int `form:"record_id" json:"record_id" binding:"min=0,numeric"`
	Limit      int `form:"limit" json:"limit" binding:"required,numeric,max=100"`
}

func (m *Message) GetRecords(ctx *core.Context) error {
	params := &GetDialogRecordsRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if params.DialogType == constant.ChatGroupMode {
		err := m.AuthUseCase.IsAuth(ctx.Ctx(), &usecase.AuthOption{
			DialogType: params.DialogType,
			UserId:     uid,
			ReceiverId: params.ReceiverId,
		})

		if err != nil {
			items := make([]map[string]any, 0)
			items = append(items, map[string]any{
				"content":     "Недостаточно прав для просмотра сообщений в группе",
				"created_at":  timeutil.DateTime(),
				"id":          1,
				"msg_id":      strutil.NewMsgId(),
				"msg_type":    constant.ChatMsgSysText,
				"receiver_id": params.ReceiverId,
				"dialog_type": params.DialogType,
				"user_id":     0,
			})

			return ctx.Success(map[string]any{
				"limit":     params.Limit,
				"record_id": 0,
				"items":     items,
			})
		}
	}

	records, err := m.MessageUseCase.GetDialogRecords(ctx.Ctx(), &usecase.QueryDialogRecordsOpt{
		DialogType: params.DialogType,
		UserId:     ctx.UserId(),
		ReceiverId: params.ReceiverId,
		RecordId:   params.RecordId,
		Limit:      params.Limit,
	})
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
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

type DownloadChatFileRequest struct {
	RecordId int `form:"cr_id" json:"cr_id" binding:"required,min=1"`
}

func (m *Message) Download(ctx *core.Context) error {
	params := &DownloadChatFileRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	record, err := m.MessageUseCase.GetMessageByRecordId(ctx.Ctx(), params.RecordId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	uid := ctx.UserId()
	if uid != record.UserId {
		if record.DialogType == constant.ChatPrivateMode {
			if record.ReceiverId != uid {
				return ctx.Forbidden("Отсутствует доступ")
			}
		} else {
			if !m.GroupChatMemberUseCase.MemberRepo.IsMember(ctx.Ctx(), record.ReceiverId, uid, false) {
				return ctx.Forbidden("Отсутствует доступ")
			}
		}
	}

	var fileInfo entity.DialogRecordExtraFile
	if err := jsonutil.Decode(record.Extra, &fileInfo); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	ctx.Context.Redirect(
		http.StatusFound,
		m.Minio.PrivateUrl(m.Minio.BucketPrivateName(), fileInfo.Path, fileInfo.Name, 60*time.Second),
	)

	return nil
}

type CollectMessageRequest struct {
	RecordId int `form:"record_id" json:"record_id" binding:"required,numeric,gt=0" label:"record_id"`
}

func (m *Message) Collect(ctx *core.Context) error {
	params := &CollectMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := m.ChatUseCase.Collect(ctx.Ctx(), ctx.UserId(), params.RecordId); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}
