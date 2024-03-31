package v1

import (
	"bytes"
	"fmt"
	"voo.su/api/pb/v1"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
	"voo.su/pkg/utils"
)

type Message struct {
	DialogService          *service.DialogService
	AuthService            *service.AuthService
	MessageSendService     service.MessageSendService
	Filesystem             *filesystem.Filesystem
	MessageService         *service.MessageService
	GroupChatMemberService *service.GroupChatMemberService
	GroupMemberRepo        *repo.GroupChatMember
	MessageRepo            *repo.Message
}

type TextMessageRequest struct {
	DialogType int    `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
	ReceiverId int    `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	Text       string `form:"text" json:"text" binding:"required,max=3000" label:"text"`
}

func (c *Message) Text(ctx *core.Context) error {
	params := &TextMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType:        params.DialogType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if err := c.MessageSendService.SendText(ctx.Ctx(), uid, &api_v1.TextMessageRequest{
		Content: params.Text,
		Receiver: &api_v1.MessageReceiver{
			DialogType: int32(params.DialogType),
			ReceiverId: int32(params.ReceiverId),
		},
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

//type CodeMessageRequest struct {
//	DialogType   int    `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
//	ReceiverId int    `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
//	Lang       string `form:"lang" json:"lang" binding:"required"`
//	Code       string `form:"code" json:"code" binding:"required,max=65535"`
//}
//
//func (c *Message) Code(ctx *core.Context) error {
//	params := &CodeMessageRequest{}
//	if err := ctx.Context.ShouldBind(params); err != nil {
//		return ctx.InvalidParams(err)
//	}
//	uid := ctx.UserId()
//	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
//		DialogType:   params.DialogType,
//		UserId:     uid,
//		ReceiverId: params.ReceiverId,
//	}); err != nil {
//		return ctx.ErrorBusiness(err.Error())
//	}
//	if err := c.MessageSendService.SendCode(ctx.Ctx(), uid, &api_v1.CodeMessageRequest{
//		Lang: params.Lang,
//		Code: params.Code,
//		Receiver: &api_v1.MessageReceiver{
//			DialogType:   int32(params.DialogType),
//			ReceiverId: int32(params.ReceiverId),
//		},
//	}); err != nil {
//		return ctx.ErrorBusiness(err.Error())
//	}
//	return ctx.Success(nil)
//}

type ImageMessageRequest struct {
	DialogType int `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
	ReceiverId int `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
}

func (c *Message) Image(ctx *core.Context) error {
	params := &ImageMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	file, err := ctx.Context.FormFile("image")
	if err != nil {
		return ctx.InvalidParams("поле 'image' обязательно!")
	}

	if !sliceutil.Include(strutil.FileSuffix(file.Filename), []string{"png", "jpg", "jpeg", "gif", "webp", "svg", "PNG", "JPG", "JPEG", "GIF", "WEBP", "SVG"}) {
		return ctx.InvalidParams("Некорректный формат загружаемого файла. Поддерживаются только форматы: png, jpg, jpeg, svg, gif и webp.")
	}

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5МБ")
	}

	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType:        params.DialogType,
		UserId:            ctx.UserId(),
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	stream, err := filesystem.ReadMultipartStream(file)
	if err != nil {
		return err
	}

	ext := strutil.FileSuffix(file.Filename)
	meta := utils.ReadImageMeta(bytes.NewReader(stream))
	filePath := fmt.Sprintf("dialog/%s/%s", timeutil.DateNumber(), strutil.GenImageName(ext, meta.Width, meta.Height))
	if err := c.Filesystem.Default.Write(stream, filePath); err != nil {
		return err
	}

	if err := c.MessageSendService.SendImage(ctx.Ctx(), ctx.UserId(), &api_v1.ImageMessageRequest{
		Url:    c.Filesystem.Default.PublicUrl(filePath),
		Width:  int32(meta.Width),
		Height: int32(meta.Height),
		Size:   int32(file.Size),
		Receiver: &api_v1.MessageReceiver{
			DialogType: int32(params.DialogType),
			ReceiverId: int32(params.ReceiverId),
		},
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type FileMessageRequest struct {
	DialogType int    `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
	ReceiverId int    `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	UploadId   string `form:"upload_id" json:"upload_id" binding:"required"`
}

func (c *Message) File(ctx *core.Context) error {
	params := &FileMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType:        params.DialogType,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if err := c.MessageSendService.SendFile(ctx.Ctx(), uid, &api_v1.FileMessageRequest{
		UploadId: params.UploadId,
		Receiver: &api_v1.MessageReceiver{
			DialogType: int32(params.DialogType),
			ReceiverId: int32(params.ReceiverId),
		},
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type VoteMessageRequest struct {
	ReceiverId int      `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	Mode       int      `form:"mode" json:"mode" binding:"oneof=0 1"`
	Anonymous  int      `form:"anonymous" json:"anonymous" binding:"oneof=0 1"`
	Title      string   `form:"title" json:"title" binding:"required"`
	Options    []string `form:"options" json:"options"`
}

func (c *Message) Vote(ctx *core.Context) error {
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
	if err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
		DialogType:        entity.ChatGroupMode,
		UserId:            uid,
		ReceiverId:        params.ReceiverId,
		IsVerifyGroupMute: true,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if err := c.MessageSendService.SendVote(ctx.Ctx(), uid, &api_v1.VoteMessageRequest{
		Mode:      int32(params.Mode),
		Title:     params.Title,
		Options:   params.Options,
		Anonymous: int32(params.Anonymous),
		Receiver: &api_v1.MessageReceiver{
			DialogType: entity.ChatGroupMode,
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

func (c *Message) Revoke(ctx *core.Context) error {
	params := &RevokeMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.MessageSendService.Revoke(ctx.Ctx(), ctx.UserId(), params.MsgId); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

type DeleteMessageRequest struct {
	DialogType int    `form:"dialog_type" json:"dialog_type" binding:"required,oneof=1 2" label:"dialog_type"`
	ReceiverId int    `form:"receiver_id" json:"receiver_id" binding:"required,numeric,gt=0" label:"receiver_id"`
	RecordIds  string `form:"record_id" json:"record_id" binding:"required,ids" label:"record_id"`
}

func (c *Message) Delete(ctx *core.Context) error {
	params := &DeleteMessageRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.DialogService.DeleteRecordList(ctx.Ctx(), &service.RemoveRecordListOpt{
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

func (c *Message) HandleVote(ctx *core.Context) error {
	params := &VoteMessageHandleRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	data, err := c.MessageSendService.Vote(ctx.Ctx(), ctx.UserId(), params.RecordId, params.Options)
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

func (c *Message) GetRecords(ctx *core.Context) error {
	params := &GetDialogRecordsRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()
	if params.DialogType == entity.ChatGroupMode {
		err := c.AuthService.IsAuth(ctx.Ctx(), &service.AuthOption{
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
				"msg_type":    entity.ChatMsgSysText,
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

	records, err := c.MessageService.GetDialogRecords(ctx.Ctx(), &service.QueryDialogRecordsOpt{
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

func (c *Message) Download(ctx *core.Context) error {
	params := &DownloadChatFileRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	record, err := c.MessageRepo.FindById(ctx.Ctx(), params.RecordId)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	uid := ctx.UserId()
	if uid != record.UserId {
		if record.DialogType == entity.ChatPrivateMode {
			if record.ReceiverId != uid {
				return ctx.Forbidden("Отсутствует доступ")
			}
		} else {
			if !c.GroupMemberRepo.IsMember(ctx.Ctx(), record.ReceiverId, uid, false) {
				return ctx.Forbidden("Отсутствует доступ")
			}
		}
	}

	var fileInfo model.DialogRecordExtraFile
	if err := jsonutil.Decode(record.Extra, &fileInfo); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	switch fileInfo.Drive {
	case entity.FileDriveLocal:
		ctx.Context.FileAttachment(c.Filesystem.Local.Path(fileInfo.Path), fileInfo.Name)
	default:
		return ctx.ErrorBusiness("Неизвестный тип драйвера файла")
	}

	return nil
}
