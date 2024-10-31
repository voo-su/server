package service

import (
	"context"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/encrypt"
)

type BotService struct {
	*repo.Source
	BotRepo  *repo.Bot
	UserRepo *repo.User
}

func NewBotService(
	source *repo.Source,
	botRepo *repo.Bot,
	userRepo *repo.User,
) *BotService {
	return &BotService{
		Source:   source,
		BotRepo:  botRepo,
		UserRepo: userRepo,
	}
}

type BotCreateOpt struct {
	Username string
}

func (b *BotService) Create(ctx context.Context, opt *BotCreateOpt) (*string, error) {
	user, err := b.UserRepo.Create(&model.User{
		Name:  opt.Username,
		IsBot: 1,
	})
	if err != nil {
		return nil, err
	}

	token := encrypt.GenerateToken()

	return &token, b.BotRepo.Create(ctx, &model.Bot{
		UserId: user.Id,
		Token:  token,
		Name:   opt.Username,
	})
}

func (b *BotService) GetBotByToken(ctx context.Context, token string) (*model.Bot, error) {
	bot, err := b.BotRepo.FindByWhere(ctx, "token = ?", token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (b *BotService) List(ctx context.Context, uid int) ([]*model.Bot, error) {
	tx := b.BotRepo.Model(ctx)
	tx.Select([]string{
		"bots.id",
		"bots.name",
		"bots.token",
	})

	var items []*model.Bot
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (b *BotService) Chats(ctx context.Context, botId int) ([]*model.SearchChat, error) {
	fields := []string{
		"g.id",
		"d.dialog_type",
		"d.receiver_id",
		"d.updated_at",
		"d.is_disturb",
		"d.is_top",
		"d.is_bot",
		"g.group_name",
		"g.avatar as group_avatar",
	}

	query := b.Source.Db().WithContext(ctx).Table("dialogs d")
	//query.Joins("LEFT JOIN users AS u ON d.receiver_id = u.id AND d.dialog_type = 1")
	query.Joins("LEFT JOIN group_chats AS g ON d.receiver_id = g.id")
	query.Where("d.user_id = ? AND d.dialog_type = 2 AND d.is_delete = 0", 2)
	query.Order("d.updated_at DESC")

	var items []*model.SearchChat
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
