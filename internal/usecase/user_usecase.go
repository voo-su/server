package usecase

import (
	"context"
	"log"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	"voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/locale"
)

type UserUseCase struct {
	Locale          locale.ILocale
	Source          *infrastructure.Source
	UserRepo        *postgresRepo.UserRepository
	UserSessionRepo *postgresRepo.UserSessionRepository
	PushTokenRepo   *postgresRepo.PushTokenRepository
}

func NewUserUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	userRepo *postgresRepo.UserRepository,
	userSessionRepo *postgresRepo.UserSessionRepository,
	pushTokenRepo *postgresRepo.PushTokenRepository,
) *UserUseCase {
	return &UserUseCase{
		Locale:          locale,
		Source:          source,
		UserRepo:        userRepo,
		UserSessionRepo: userSessionRepo,
		PushTokenRepo:   pushTokenRepo,
	}
}

func (u *UserUseCase) WebPushInit(ctx context.Context, uid int64, webPush entity.WebPush) {
	if err := u.PushTokenRepo.Create(ctx, &model.PushToken{
		UserId:      uid,
		Platform:    "webpush",
		WebEndpoint: webPush.Endpoint,
		WebP256dh:   webPush.Keys.P256dh,
		WebAuth:     webPush.Keys.Auth,
	}); err != nil {
		log.Println(err)
	}
}
