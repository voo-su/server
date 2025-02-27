package usecase

import (
	"context"
	"log"
	"voo.su/internal/constant"
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

func (u *UserUseCase) WebPushInit(ctx context.Context, uid int64, webPush *entity.WebPush) error {
	if err := u.PushTokenRepo.Create(ctx, &model.PushToken{
		UserId:      uid,
		Platform:    constant.PushPlatformWeb,
		WebEndpoint: webPush.Endpoint,
		WebP256dh:   webPush.Keys.P256dh,
		WebAuth:     webPush.Keys.Auth,
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *UserUseCase) GetNotifySettings(ctx context.Context, uid int) (*entity.NotifySettings, error) {
	user, err := u.UserRepo.FindById(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &entity.NotifySettings{
		ChatsMuteUntil:    user.NotifyChatsMuteUntil,
		ChatsShowPreviews: user.NotifyChatsShowPreviews,
		ChatsSilent:       user.NotifyChatsSilent,
		GroupMuteUntil:    user.NotifyGroupMuteUntil,
		GroupShowPreviews: user.NotifyGroupShowPreviews,
		GroupSilent:       user.NotifyGroupSilent,
	}, nil
}

func (u *UserUseCase) UpdateNotifySettings(ctx context.Context, uid int, data *entity.NotifySettings) error {
	//_, err := u.UserRepo.UpdateById(ctx, uid, map[string]any{
	//	"personal_chats": data.PersonalChats,
	//	"group_chats":    data.GroupChats,
	//})
	//if err != nil {
	//	return err
	//}
	//
	return nil
}

func (u *UserUseCase) RegisterDevice(ctx context.Context, uid int64, tokenType int32, token string) error {
	if err := u.PushTokenRepo.Create(ctx, &model.PushToken{
		UserId:   uid,
		Platform: constant.PushPlatformMobile,
		Token:    token,
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
