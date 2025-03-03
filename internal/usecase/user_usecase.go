package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

func (u *UserUseCase) WebPushInit(ctx context.Context, uid int64, sessionId int64, webPush *entity.WebPush) error {

	existingToken, err := u.PushTokenRepo.FindByWhere(ctx, "user_session_id = ? AND is_active = ?", sessionId, true)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if existingToken != nil {
		_, err := u.PushTokenRepo.UpdateById(ctx, existingToken.Id, map[string]any{
			"web_endpoint": webPush.Endpoint,
			"web_p256dh":   webPush.Keys.P256dh,
			"web_auth":     webPush.Keys.Auth,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		if err := u.PushTokenRepo.Create(ctx, &model.PushToken{
			UserId:        uid,
			UserSessionId: sessionId,
			Platform:      constant.PushPlatformWeb,
			WebEndpoint:   webPush.Endpoint,
			WebP256dh:     webPush.Keys.P256dh,
			WebAuth:       webPush.Keys.Auth,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (u *UserUseCase) DevicePushInit(ctx context.Context, uid int64, sessionId int64, tokenType int32, token string) error {
	if token == "" {
		return errors.New(u.Locale.Localize("token_not_provided"))
	}

	existingToken, err := u.PushTokenRepo.FindByWhere(ctx, "user_session_id = ? AND is_active = ?", sessionId, true)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if existingToken != nil {
		_, err := u.PushTokenRepo.UpdateById(ctx, existingToken.Id, map[string]any{
			"token": token,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		if err := u.PushTokenRepo.Create(ctx, &model.PushToken{
			UserId:        uid,
			UserSessionId: sessionId,
			Platform:      constant.PushPlatformMobile,
			Token:         token,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (u *UserUseCase) GetNotifySettings(ctx context.Context, notifyType int, uid int) (*entity.NotifySettings, error) {
	user, err := u.UserRepo.FindById(ctx, uid)
	if err != nil {
		return nil, err
	}

	settings := entity.NotifySettings{}
	switch notifyType {
	case constant.ChatPrivateMode:
		settings = entity.NotifySettings{
			MuteUntil:    user.NotifyChatsMuteUntil,
			ShowPreviews: user.NotifyChatsShowPreviews,
			Silent:       user.NotifyChatsSilent,
		}
	case constant.ChatGroupMode:
		settings = entity.NotifySettings{
			MuteUntil:    user.NotifyGroupMuteUntil,
			ShowPreviews: user.NotifyGroupShowPreviews,
			Silent:       user.NotifyGroupSilent,
		}
	default:
		return nil, errors.New("")
	}

	return &settings, nil
}

func (u *UserUseCase) UpdateNotifySettings(ctx context.Context, notifyType int, uid int, data *entity.NotifySettings) error {
	settings := map[string]any{}
	switch notifyType {
	case constant.ChatPrivateMode:
		settings = map[string]any{
			"notify_chats_mute_until":    data.MuteUntil,
			"notify_chats_show_previews": data.ShowPreviews,
			"notify_chats_silent":        data.Silent,
		}
	case constant.ChatGroupMode:
		settings = map[string]any{
			"notify_group_mute_until":    data.MuteUntil,
			"notify_group_show_previews": data.ShowPreviews,
			"notify_group_silent":        data.Silent,
		}
	default:
		return errors.New("")
	}
	_, err := u.UserRepo.UpdateById(ctx, uid, settings)
	if err != nil {
		return err
	}

	return nil
}
