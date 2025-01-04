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
	"voo.su/pkg"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"

	clickhouseModel "voo.su/internal/infrastructure/clickhouse/model"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/email"
	"voo.su/pkg/strutil"
	"voo.su/resource"
)

type AuthUseCase struct {
	Conf                *config.Config
	Locale              locale.ILocale
	Email               *email.Email
	SmsCache            *redisRepo.SmsCacheRepository
	ContactRepo         *postgresRepo.ContactRepository
	GroupChatMemberRepo *postgresRepo.GroupChatMemberRepository
	GroupChatRepo       *postgresRepo.GroupChatRepository
	UserRepo            *postgresRepo.UserRepository
	AuthCodeRepo        *clickhouseRepo.AuthCodeRepository
	JwtTokenCacheRepo   *redisRepo.JwtTokenCacheRepository
}

func NewAuthUseCase(
	conf *config.Config,
	locale locale.ILocale,
	email *email.Email,
	smsCache *redisRepo.SmsCacheRepository,
	contactRepo *postgresRepo.ContactRepository,
	groupChatMember *postgresRepo.GroupChatMemberRepository,
	groupChat *postgresRepo.GroupChatRepository,
	repo *postgresRepo.UserRepository,
	authCodeRepo *clickhouseRepo.AuthCodeRepository,
	jwtTokenCacheRepository *redisRepo.JwtTokenCacheRepository,
) *AuthUseCase {
	return &AuthUseCase{
		Conf:                conf,
		Locale:              locale,
		Email:               email,
		SmsCache:            smsCache,
		ContactRepo:         contactRepo,
		GroupChatMemberRepo: groupChatMember,
		GroupChatRepo:       groupChat,
		UserRepo:            repo,
		AuthCodeRepo:        authCodeRepo,
		JwtTokenCacheRepo:   jwtTokenCacheRepository,
	}
}

func (a *AuthUseCase) Login(ctx context.Context, guard string, _email string) (*string, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(constant.ExpiresTime))

	token := jwt.GenerateToken(guard, a.Conf.App.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        _email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	code, err := a.Send(ctx, constant.LoginChannel, _email, token)
	if err != nil {
		return nil, err
	}

	if err := a.AuthCodeRepo.Create(ctx, &clickhouseModel.AuthCode{
		Email:        _email,
		Code:         *code,
		Token:        token,
		ErrorMessage: "",
	}); err != nil {
		log.Println(err)
	}

	return &token, nil
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
		return errors.New(a.Locale.Localize("no_permission_to_send_messages"))
	}

	groupInfo, err := a.GroupChatRepo.FindById(ctx, opt.ReceiverId)
	if err != nil {
		return err
	}

	if groupInfo.IsDismiss == 1 {
		return errors.New(a.Locale.Localize("group_deleted"))
	}

	memberInfo, err := a.GroupChatMemberRepo.FindByUserId(ctx, opt.ReceiverId, opt.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(a.Locale.Localize("no_permission_to_send_messages"))
		}
		return errors.New(a.Locale.Localize("system_busy_try_later"))
	}

	if memberInfo.IsQuit == constant.GroupMemberQuitStatusYes {
		return errors.New(a.Locale.Localize("no_permission_to_send_messages"))
	}

	if memberInfo.IsMute == constant.GroupMemberMuteStatusYes {
		return errors.New(a.Locale.Localize("message_sending_prohibited_by_admin"))
	}

	if opt.IsVerifyGroupMute && groupInfo.IsMute == 1 && memberInfo.Leader == 0 {
		return errors.New(a.Locale.Localize("group_message_sending_disabled"))
	}

	return nil
}

func (a *AuthUseCase) CodeTemplate(data map[string]string) (string, error) {
	fileContent, err := resource.Template().ReadFile("template/email_verify_code.tmpl")
	if err != nil {
		return "", err
	}

	return pkg.RenderTemplate(fileContent, data)
}

func (a *AuthUseCase) Verify(ctx context.Context, channel string, token string, code string) bool {
	return a.SmsCache.Verify(ctx, channel, token, code)
}

func (a *AuthUseCase) Delete(ctx context.Context, channel string, token string) error {
	return a.SmsCache.Del(ctx, channel, token)
}

func (a *AuthUseCase) Send(ctx context.Context, channel string, _email string, token string) (*string, error) {
	code := strutil.GenValidateCode(6)
	if err := a.SmsCache.Set(ctx, channel, token, code, 15*time.Minute); err != nil {
		return nil, err
	}

	if a.Conf.App.Env == "dev" {
		fmt.Println("Mail: ", _email)
		fmt.Println("Code: ", code)
		return &code, nil
	}

	body, err := a.CodeTemplate(map[string]string{
		"code": code,
	})
	if err != nil {
		return nil, err
	}

	if err := a.Email.SendMail(&email.Option{
		To:      _email,
		Subject: a.Locale.Localize("welcome"),
		Body:    body,
	}); err != nil {
		fmt.Println(err)
	}

	return &code, nil
}

func (a *AuthUseCase) Register(ctx context.Context, email string) (*postgresModel.User, error) {
	if a.UserRepo.IsEmailExist(ctx, email) {
		return a.UserRepo.FindByEmail(email)
	}

	for {
		username := uuid.New().String()
		var user postgresModel.User
		result := a.UserRepo.Db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				user.Email = email
				user.Username = username
				return a.UserRepo.Create(&user)
			}
			log.Fatal(result.Error)
		}
	}

	return nil, errors.New(a.Locale.Localize("general_error"))
}
