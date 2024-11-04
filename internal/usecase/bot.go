package usecase

import (
	"context"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/encrypt"
)

type BotUseCase struct {
	*repo.Source
	BotRepo  *repo.Bot
	UserRepo *repo.User
}

func NewBotUseCase(
	source *repo.Source,
	botRepo *repo.Bot,
	userRepo *repo.User,
) *BotUseCase {
	return &BotUseCase{
		Source:   source,
		BotRepo:  botRepo,
		UserRepo: userRepo,
	}
}

type BotCreateOpt struct {
	Username  string
	CreatorId int
}

func (b *BotUseCase) Create(ctx context.Context, opt *BotCreateOpt) (*string, error) {
	user, err := b.UserRepo.Create(&model.User{
		Username: opt.Username,
		Name:     opt.Username,
		IsBot:    1,
	})
	if err != nil {
		return nil, err
	}

	token := encrypt.GenerateToken()

	return &token, b.BotRepo.Create(ctx, &model.Bot{
		UserId:    user.Id,
		Token:     token,
		Name:      opt.Username,
		CreatorId: opt.CreatorId,
	})
}

func (b *BotUseCase) GetBotByToken(ctx context.Context, token string) (*model.Bot, error) {
	bot, err := b.BotRepo.FindByWhere(ctx, "token = ?", token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (b *BotUseCase) List(ctx context.Context, uid int) ([]*model.Bot, error) {
	tx := b.BotRepo.Model(ctx)
	tx.Select([]string{
		"bots.id",
		"bots.name",
		"bots.token",
	})
	tx.Where("bots.creator_id = ?", uid)

	var items []*model.Bot
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (b *BotUseCase) Chats(ctx context.Context, botId int) ([]*entity.SearchChat, error) {
	fields := []string{
		"g.id",
		"c.dialog_type",
		"c.receiver_id",
		"c.updated_at",
		"c.is_disturb",
		"c.is_top",
		"c.is_bot",
		"g.group_name",
		"g.avatar as group_avatar",
	}

	query := b.Source.Db().WithContext(ctx).Table("chats c")
	//query.Joins("LEFT JOIN users AS u ON c.receiver_id = u.id AND c.dialog_type = 1")
	query.Joins("LEFT JOIN group_chats AS g ON c.receiver_id = g.id")
	query.Where("c.user_id = ? AND c.dialog_type = 2 AND c.is_delete = 0", 2)
	query.Order("c.updated_at DESC")

	var items []*entity.SearchChat
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
