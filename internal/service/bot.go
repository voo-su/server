package service

import (
	"context"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/encrypt"
)

type BotService struct {
	BotRepo  *repo.Bot
	UserRepo *repo.User
}

func NewBotService(
	botRepo *repo.Bot,
	userRepo *repo.User,
) *BotService {
	return &BotService{
		BotRepo:  botRepo,
		UserRepo: userRepo,
	}
}

type BotCreateOpt struct {
	Name string
}

func (b *BotService) Create(ctx context.Context, opt *BotCreateOpt) (*string, error) {
	user, err := b.UserRepo.Create(&model.User{
		Name:  opt.Name,
		IsBot: 1,
	})
	if err != nil {
		return nil, err
	}

	token := encrypt.GenerateToken()

	return &token, b.BotRepo.Create(ctx, &model.Bot{
		UserId: user.Id,
		Token:  token,
		Name:   opt.Name,
	})
}

func (b *BotService) GetBotByToken(ctx context.Context, token string) (*model.Bot, error) {
	bot, err := b.BotRepo.FindByWhere(ctx, "token = ?", token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (b *BotService) List(ctx context.Context, uid int) ([]*model.ContactListItem, error) {
	tx := b.BotRepo.Model(ctx)
	tx.Select([]string{
		"bots.id",
		"bots.name",
	})

	var items []*model.ContactListItem
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
