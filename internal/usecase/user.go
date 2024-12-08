package usecase

import (
	"context"
	"log"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type UserUseCase struct {
	*repository.Source
	UserRepo        *repo.User
	UserSessionRepo *repo.UserSession
	PushTokenRepo   *repo.PushToken
}

func NewUserUseCase(
	source *repository.Source,
	userRepo *repo.User,
	userSessionRepo *repo.UserSession,
	pushTokenRepo *repo.PushToken,
) *UserUseCase {
	return &UserUseCase{
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
