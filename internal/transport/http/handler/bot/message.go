package bot

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	botPb "voo.su/api/http/pb/bot"
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

	var bot, err = m.BotUseCase.GetBotByToken(ctx.Ctx(), token)
	if err != nil {
		return nil, ctx.ErrorBusiness("")
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

// TODO DELETE
func (m *Message) Send(ctx *core.Context) error {
	params := &botPb.MessageRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
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
		Content: params.Content,
	}); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&botPb.MessageSendResponse{})
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
		return ctx.InvalidParams("")
	}

	file, err := ctx.Context.FormFile("photo")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5МБ")
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки")
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
		return ctx.InvalidParams("")
	}

	file, err := ctx.Context.FormFile("video")
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5МБ")
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки")
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
		return ctx.InvalidParams("")
	}

	file, err := ctx.Context.FormFile("audio")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5МБ")
	}

	filePath, err := m.BotUseCase.FileUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки")
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
		return ctx.InvalidParams("")
	}

	file, err := ctx.Context.FormFile("document")
	if err != nil {
		return ctx.InvalidParams("Ошибка загрузки файла")
	}

	caption := ctx.Context.DefaultPostForm("caption", "")

	if file.Size > 5<<20 {
		return ctx.InvalidParams("Размер загружаемого файла не может превышать 5МБ")
	}

	filePath, err := m.BotUseCase.FileDocumentUpload(ctx.Ctx(), file)
	if err != nil {
		fmt.Println(err)
		return ctx.InvalidParams("Ошибка загрузки")
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
		return err
	}

	return ctx.Success(&botPb.MessageSendResponse{})
}
