package usecase

import (
	"bytes"
	"context"
	"mime/multipart"
	"voo.su/internal/config"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/strutil"
	"voo.su/pkg/utils"
)

type BotUseCase struct {
	Conf     *config.Config
	Locale   locale.ILocale
	Source   *infrastructure.Source
	BotRepo  *postgresRepo.BotRepository
	UserRepo *postgresRepo.UserRepository
	Minio    minio.IMinio
}

func NewBotUseCase(
	conf *config.Config,
	locale locale.ILocale,
	source *infrastructure.Source,
	botRepo *postgresRepo.BotRepository,
	userRepo *postgresRepo.UserRepository,
	minio minio.IMinio,
) *BotUseCase {
	return &BotUseCase{
		Conf:     conf,
		Locale:   locale,
		Source:   source,
		BotRepo:  botRepo,
		UserRepo: userRepo,
		Minio:    minio,
	}
}

type BotCreateOpt struct {
	Username  string
	CreatorId int
}

func (b *BotUseCase) Create(ctx context.Context, opt *BotCreateOpt) (*string, error) {
	user, err := b.UserRepo.Create(&postgresModel.User{
		Username: opt.Username,
		Name:     opt.Username,
		IsBot:    1,
	})
	if err != nil {
		return nil, err
	}

	token := encrypt.GenerateToken()

	return &token, b.BotRepo.Create(ctx, &postgresModel.Bot{
		UserId:    user.Id,
		Token:     token,
		Name:      opt.Username,
		CreatorId: opt.CreatorId,
	})
}

func (b *BotUseCase) GetBotByToken(ctx context.Context, token string) (*postgresModel.Bot, error) {
	bot, err := b.BotRepo.FindByWhere(ctx, "token = ?", token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (b *BotUseCase) List(ctx context.Context, uid int) ([]*postgresModel.Bot, error) {
	fields := []string{
		"bots.id",
		"bots.name",
		"bots.token",
	}
	tx := b.BotRepo.Model(ctx).
		Select(fields).
		Where("bots.creator_id = ?", uid)

	var items []*postgresModel.Bot
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

	query := b.Source.Postgres().WithContext(ctx).Table("chats c").
		Joins("LEFT JOIN group_chats AS g ON c.receiver_id = g.id").
		Where("c.user_id = ? AND c.dialog_type = ? AND c.is_delete = ?", botId, 2, 0).
		Order("c.updated_at DESC")

	var items []*entity.SearchChat
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (b *BotUseCase) FileUpload(ctx context.Context, file *multipart.FileHeader) (*string, error) {
	stream, err := minio.ReadMultipartStream(file)
	if err != nil {
		return nil, err
	}

	meta := utils.ReadImageMeta(bytes.NewReader(stream))
	ext := strutil.FileSuffix(file.Filename)

	src := strutil.GenMediaObjectName(ext, meta.Width, meta.Height)
	if err = b.Minio.Write(b.Conf.Minio.GetBucket(), src, stream); err != nil {
		return nil, err
	}

	path := b.Minio.PublicUrl(b.Conf.Minio.GetBucket(), src)

	return &path, nil
}

func (b *BotUseCase) FileDocumentUpload(ctx context.Context, file *multipart.FileHeader) (*string, error) {
	stream, err := minio.ReadMultipartStream(file)
	if err != nil {
		return nil, err
	}

	meta := utils.ReadImageMeta(bytes.NewReader(stream))
	ext := strutil.FileSuffix(file.Filename)

	src := strutil.GenMediaObjectName(ext, meta.Width, meta.Height)
	if err = b.Minio.Write(b.Conf.Minio.GetBucket(), src, stream); err != nil {
		return nil, err
	}

	return &src, nil
}
