package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	clickhouseModel "voo.su/internal/infrastructure/clickhouse/model"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg"
	"voo.su/pkg/email"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/strutil"
	"voo.su/resource"
)

type AuthUseCase struct {
	Conf                *config.Config
	Locale              locale.ILocale
	Email               *email.Email
	SmsCache            *redisRepo.SmsCacheRepository
	GroupChatMemberRepo *postgresRepo.GroupChatMemberRepository
	UserRepo            *postgresRepo.UserRepository
	UserSessionRepo     *postgresRepo.UserSessionRepository
	AuthCodeRepo        *clickhouseRepo.AuthCodeRepository
	JwtTokenCacheRepo   *redisRepo.JwtTokenCacheRepository
}

func NewAuthUseCase(
	conf *config.Config,
	locale locale.ILocale,
	email *email.Email,
	smsCache *redisRepo.SmsCacheRepository,
	groupChatMember *postgresRepo.GroupChatMemberRepository,
	userRepo *postgresRepo.UserRepository,
	userSessionRepo *postgresRepo.UserSessionRepository,
	authCodeRepo *clickhouseRepo.AuthCodeRepository,
	jwtTokenCacheRepository *redisRepo.JwtTokenCacheRepository,
) *AuthUseCase {
	return &AuthUseCase{
		Conf:                conf,
		Locale:              locale,
		Email:               email,
		SmsCache:            smsCache,
		GroupChatMemberRepo: groupChatMember,
		UserRepo:            userRepo,
		UserSessionRepo:     userSessionRepo,
		AuthCodeRepo:        authCodeRepo,
		JwtTokenCacheRepo:   jwtTokenCacheRepository,
	}
}

func (a *AuthUseCase) Login(ctx context.Context, guard string, _email string) (*string, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(constant.ExpiresTime))

	token := jwtutil.GenerateToken(guard, a.Conf.App.Jwt.Secret, &jwtutil.Options{
		ExpiresAt: jwtutil.NewNumericDate(expiresAt),
		ID:        _email,
		IssuedAt:  jwtutil.NewNumericDate(time.Now()),
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
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("email не может быть пустым")
	}

	lowerEmail := strings.ToLower(email)
	user, err := a.UserRepo.FindByEmail(lowerEmail)
	if err == nil {
		return user, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}
	username := uuid.New().String()
	user = &postgresModel.User{
		Email:    email,
		Username: username,
	}

	return a.UserRepo.Create(user)
}

func (a *AuthUseCase) GetSessionByToken(ctx context.Context, token string) (*postgresModel.UserSession, error) {
	session, err := a.UserSessionRepo.FindByWhere(ctx, "access_token = ? AND is_logout = ?", token, false)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *AuthUseCase) Logout(ctx context.Context, token string) error {
	if err := a.JwtTokenCacheRepo.SetBlackList(ctx, token, 0); err != nil {
		log.Println(err)
	}

	session, err := a.GetSessionByToken(ctx, token)
	if err != nil {
		return status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	_, err = a.UserSessionRepo.UpdateById(ctx, session.Id, map[string]any{
		"is_logout": true,
		"logout_at": time.Now(),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
