package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/email"
	"voo.su/pkg/strutil"
	"voo.su/pkg/utils"
	"voo.su/resource"
)

type AuthUseCase struct {
	SmsStorage      *cache.SmsStorage
	ContactRepo     *repo.Contact
	GroupChatMember *repo.GroupChatMember
	GroupChat       *repo.GroupChat
	Conf            *config.Config
	Email           *email.Email
	User            *repo.User
}

func NewAuthUseCase(
	smsStorage *cache.SmsStorage,
	contactRepo *repo.Contact,
	groupChatMember *repo.GroupChatMember,
	groupChat *repo.GroupChat,
	conf *config.Config,
	email *email.Email,
	repo *repo.User,
) *AuthUseCase {
	return &AuthUseCase{
		SmsStorage:      smsStorage,
		ContactRepo:     contactRepo,
		GroupChatMember: groupChatMember,
		GroupChat:       groupChat,
		Conf:            conf,
		Email:           email,
		User:            repo,
	}
}

type AuthOption struct {
	DialogType        int
	UserId            int
	ReceiverId        int
	IsVerifyGroupMute bool
}

func (a *AuthUseCase) IsAuth(ctx context.Context, opt *AuthOption) error {
	if opt.DialogType == constant.ChatPrivateMode {
		if a.ContactRepo.IsFriend(ctx, opt.UserId, opt.ReceiverId, false) {
			return nil
		}
		return errors.New("нет прав на отправку сообщений")
	}

	groupInfo, err := a.GroupChat.FindById(ctx, opt.ReceiverId)
	if err != nil {
		return err
	}

	if groupInfo.IsDismiss == 1 {
		return errors.New("группа удалена")
	}

	memberInfo, err := a.GroupChatMember.FindByUserId(ctx, opt.ReceiverId, opt.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("нет прав на отправку сообщений")
		}
		return errors.New("система занята, пожалуйста, попробуйте позже")
	}

	if memberInfo.IsQuit == constant.GroupMemberQuitStatusYes {
		return errors.New("у вас нет прав на отправку сообщений")
	}

	if memberInfo.IsMute == constant.GroupMemberMuteStatusYes {
		return errors.New("отправка сообщений запрещена администратором или владельцем группы")
	}

	if opt.IsVerifyGroupMute && groupInfo.IsMute == 1 && memberInfo.Leader == 0 {
		return errors.New("в данной группе включено полное отключение возможности отправки сообщений для всех участников")
	}

	return nil
}

func (a *AuthUseCase) CodeTemplate(data map[string]string) (string, error) {
	fileContent, err := resource.Template().ReadFile("template/email_verify_code.tmpl")
	if err != nil {
		return "", err
	}

	return utils.RenderTemplate(fileContent, data)
}

func (a *AuthUseCase) Verify(ctx context.Context, channel string, token string, code string) bool {
	return a.SmsStorage.Verify(ctx, channel, token, code)
}

func (a *AuthUseCase) Delete(ctx context.Context, channel string, token string) {
	_ = a.SmsStorage.Del(ctx, channel, token)
}

func (a *AuthUseCase) Send(ctx context.Context, channel string, _email string, token string) error {
	code := strutil.GenValidateCode(6)
	if err := a.SmsStorage.Set(ctx, channel, token, code, 15*time.Minute); err != nil {
		return err
	}

	if a.Conf.App.Env == "dev" {
		fmt.Println("Почта: ", _email)
		fmt.Println("Код: ", code)
		return nil
	}

	body, err := a.CodeTemplate(map[string]string{"code": code})
	if err != nil {
		return err
	}

	_ = a.Email.SendMail(&email.Option{
		To:      _email,
		Subject: "Добро пожаловать",
		Body:    body,
	})

	return nil
}

func (a *AuthUseCase) Register(ctx context.Context, email string) (*model.User, error) {
	if a.User.IsEmailExist(ctx, email) {
		return a.User.FindByEmail(email)
	}

	for {
		username := uuid.New().String()
		var user model.User
		result := a.User.Db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				user.Email = email
				user.Username = username
				return a.User.Create(&user)
			}
			log.Fatal(result.Error)
		}
	}

	return nil, fmt.Errorf("Ошибка")
}
