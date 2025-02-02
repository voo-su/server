package bot

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	botPb "voo.su/api/http/pb/bot"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/strutil"
)

type Message struct {
	Locale         locale.ILocale
	MessageUseCase usecase.IMessageUseCase
	BotUseCase     *usecase.BotUseCase
}

func (m *Message) checkBot(ctx *ginutil.Context) (*model.Bot, error) {
	token := ctx.Context.Param("token")
	if token == "" {
		return nil, ctx.Error(m.Locale.Localize("token_missing_or_empty"))
	}

	var bot, err = m.BotUseCase.GetBotByToken(ctx.Ctx(), token)
	if err != nil {
		return nil, ctx.Error(m.Locale.Localize("bot_data_request_error"))
	}

	if bot == nil {
		return nil, ctx.Error(m.Locale.Localize("bot_not_found_with_token"))
	}

	return bot, nil
}

func (m *Message) GroupChats(ctx *ginutil.Context) error {
	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	list, err := m.BotUseCase.Chats(ctx.Ctx(), bot.CreatorId)
	if err != nil {
		return ctx.Error(err.Error())
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

func (m *Message) Message(ctx *ginutil.Context) error {
	params := &botPb.MessageSendRequest{}
	if err := ctx.Context.Bind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendText(ctx.Ctx(), bot.UserId, &entity.SendText{
		Receiver: entity.MessageReceiver{
			ChatType:   2,
			ReceiverId: params.ChatId,
		},
		Content: params.Text,
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}

func (m *Message) Photo(ctx *ginutil.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("invalid_chat_id_format"))
	}

	file, err := ctx.Context.FormFile("photo")
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("image_upload_error"))
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotPhotoFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("image_upload_size_exceeded"), constant.BotPhotoFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams(m.Locale.Localize("image_upload_error"))
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendImage(ctx.Ctx(), bot.UserId, &entity.SendImage{
		Receiver: entity.MessageReceiver{
			ChatType:   2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Video(ctx *ginutil.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("invalid_chat_id_format"))
	}

	file, err := ctx.Context.FormFile("video")
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("video_upload_error"))
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotVideoFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("video_upload_size_exceeded"), constant.BotVideoFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams(m.Locale.Localize("video_upload_error"))
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendVideo(ctx.Ctx(), bot.UserId, &entity.SendVideo{
		Receiver: entity.MessageReceiver{
			ChatType:   2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Audio(ctx *ginutil.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("invalid_chat_id_format"))
	}

	file, err := ctx.Context.FormFile("audio")
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("audio_upload_error"))
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotAudioFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("audio_upload_size_exceeded"), constant.BotAudioFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams(m.Locale.Localize("audio_upload_error"))
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendAudio(ctx.Ctx(), bot.UserId, &entity.SendAudio{
		Receiver: entity.MessageReceiver{
			ChatType:   2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Document(ctx *ginutil.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("invalid_chat_id_format"))
	}

	file, err := ctx.Context.FormFile("document")
	if err != nil {
		return ctx.InvalidParams(m.Locale.Localize("document_upload_error"))
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotDocumentFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf(m.Locale.Localize("document_upload_size_exceeded"), constant.BotDocumentFileSize))
	}

	filePath, err := m.BotUseCase.FileDocumentUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams(m.Locale.Localize("document_upload_error"))
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendBotFile(ctx.Ctx(), bot.UserId, &entity.SendBotFile{
		Receiver: entity.MessageReceiver{
			ChatType:   2,
			ReceiverId: int32(chatId),
		},
		Drive:        1,
		OriginalName: file.Filename,
		FileExt:      strings.TrimPrefix(path.Ext(file.Filename), "."),
		FileSize:     int(file.Size),
		FilePath:     *filePath,
		Content:      strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}
