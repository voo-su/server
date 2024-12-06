package bot

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	botPb "voo.su/api/http/pb/bot"
	"voo.su/internal/constant"
	"voo.su/internal/repository/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/strutil"
)

type Message struct {
	MessageUseCase usecase.IMessageUseCase
	BotUseCase     *usecase.BotUseCase
}

func (m *Message) checkBot(ctx *core.Context) (*model.Bot, error) {
	token := ctx.Context.Param("token")
	if token == "" {
		return nil, ctx.ErrorBusiness("Токен не передан или пуст")
	}

	var bot, err = m.BotUseCase.GetBotByToken(ctx.Ctx(), token)
	if err != nil {
		return nil, ctx.ErrorBusiness("Ошибка при запросе данных о боте. Попробуйте позже.")
	}

	if bot == nil {
		return nil, ctx.ErrorBusiness("Не удалось найти бота с данным токеном")
	}

	return bot, nil
}

func (m *Message) GroupChats(ctx *core.Context) error {
	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
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

func (m *Message) Message(ctx *core.Context) error {
	params := &botPb.MessageSendRequest{}
	if err := ctx.Context.Bind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendText(ctx.Ctx(), bot.UserId, &usecase.SendText{
		Receiver: usecase.MessageReceiver{
			DialogType: 2,
			ReceiverId: params.ChatId,
		},
		Content: params.Text,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}

func (m *Message) Photo(ctx *core.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams("Неверный формат chat_id")
	}

	file, err := ctx.Context.FormFile("photo")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла изображения")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotPhotoFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf("Размер загружаемого изображения не может превышать %v МБ", constant.BotPhotoFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки изображения")
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendImage(ctx.Ctx(), bot.UserId, &usecase.SendImage{
		Receiver: usecase.MessageReceiver{
			DialogType: 2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Video(ctx *core.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams("Неверный формат chat_id")
	}

	file, err := ctx.Context.FormFile("video")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки видеофайла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotVideoFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf("Размер загружаемого видео не может превышать %v МБ", constant.BotVideoFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки видеофайла")
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendVideo(ctx.Ctx(), bot.UserId, &usecase.SendVideo{
		Receiver: usecase.MessageReceiver{
			DialogType: 2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Audio(ctx *core.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams("Неверный формат chat_id")
	}

	file, err := ctx.Context.FormFile("audio")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки аудиофайла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotAudioFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf("Размер загружаемого аудиофайла не может превышать %v МБ", constant.BotAudioFileSize))
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки аудиофайла")
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendAudio(ctx.Ctx(), bot.UserId, &usecase.SendAudio{
		Receiver: usecase.MessageReceiver{
			DialogType: 2,
			ReceiverId: int32(chatId),
		},
		Url:     *filePath,
		Content: strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil)
}

func (m *Message) Document(ctx *core.Context) error {
	chatId, err := strconv.Atoi(ctx.Context.PostForm("chat_id"))
	if err != nil {
		return ctx.InvalidParams("Неверный формат chat_id")
	}

	file, err := ctx.Context.FormFile("document")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки документа")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > constant.BotDocumentFileSize<<20 {
		return ctx.InvalidParams(fmt.Sprintf("Размер загружаемого документа не может превышать %v МБ", constant.BotDocumentFileSize))
	}

	filePath, err := m.BotUseCase.FileDocumentUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки документа")
	}

	bot, err := m.checkBot(ctx)
	if err != nil {
		return err
	}

	if err := m.MessageUseCase.SendBotFile(ctx.Ctx(), bot.UserId, &usecase.SendBotFile{
		Receiver: usecase.MessageReceiver{
			DialogType: 2,
			ReceiverId: int32(chatId),
		},
		Drive:        1,
		OriginalName: file.Filename,
		FileExt:      strings.TrimPrefix(path.Ext(file.Filename), "."),
		FileSize:     int(file.Size),
		FilePath:     *filePath,
		Content:      strutil.EscapeHtml(caption),
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}
