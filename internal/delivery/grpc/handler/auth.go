package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
	authPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/delivery/grpc/middleware"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/jwt"
	"voo.su/pkg/utils"
)

type AuthHandler struct {
	authPb.UnimplementedAuthServiceServer
	Conf              *config.Config
	TokenMiddleware   *middleware.TokenMiddleware
	AuthUseCase       *usecase.AuthUseCase
	JwtTokenCacheRepo *redisRepo.JwtTokenCacheRepository
	IpAddressUseCase  *usecase.IpAddressUseCase
	ChatUseCase       *usecase.ChatUseCase
	BotRepo           *postgresRepo.BotRepository
	MessageUseCase    usecase.IMessageUseCase
	UserSession       *postgresRepo.UserSessionRepository
}

func NewAuthHandler(
	conf *config.Config,
	tokenMiddleware *middleware.TokenMiddleware,
	authUseCase *usecase.AuthUseCase,
	jwtTokenCacheRepo *redisRepo.JwtTokenCacheRepository,
	ipAddressUseCase *usecase.IpAddressUseCase,
	chatUseCase *usecase.ChatUseCase,
	botRepo *postgresRepo.BotRepository,
	userSession *postgresRepo.UserSessionRepository,
) *AuthHandler {
	return &AuthHandler{
		Conf:              conf,
		TokenMiddleware:   tokenMiddleware,
		AuthUseCase:       authUseCase,
		JwtTokenCacheRepo: jwtTokenCacheRepo,
		IpAddressUseCase:  ipAddressUseCase,
		ChatUseCase:       chatUseCase,
		BotRepo:           botRepo,
		UserSession:       userSession,
	}
}

func (a *AuthHandler) Login(ctx context.Context, in *authPb.AuthLoginRequest) (*authPb.AuthLoginResponse, error) {
	token, err := a.AuthUseCase.Login(ctx, "grpc-auth", in.Email)
	if err != nil {
		grpclog.Errorf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return &authPb.AuthLoginResponse{
		Token:     *token,
		ExpiresIn: constant.ExpiresTime,
	}, nil
}

func (a *AuthHandler) Verify(ctx context.Context, in *authPb.AuthVerifyRequest) (*authPb.AuthVerifyResponse, error) {
	if !a.AuthUseCase.Verify(ctx, constant.LoginChannel, in.Token, in.Code) {
		return nil, status.Error(codes.Unauthenticated, "Неверный код")
	}

	claims, err := jwt.ParseToken(in.Token, a.Conf.App.Jwt.Secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный токен")
	}

	if claims.Guard != "grpc-auth" || claims.Valid() != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный токен")
	}

	ip := utils.GetGrpcClientIp(ctx)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {

	}

	userAgent := md.Get("user-agent")
	if len(userAgent) == 0 {
		userAgent[0] = "unknown"
	}

	user, err := a.AuthUseCase.Register(ctx, claims.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := a.AuthUseCase.Delete(ctx, constant.LoginChannel, in.Token); err != nil {
		fmt.Println(err)
	}

	bot, err := a.BotRepo.GetLoginBot(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if bot != nil {
		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		_, err = a.ChatUseCase.Create(ctx, &usecase.CreateChatOpt{
			UserId:     user.Id,
			DialogType: constant.ChatPrivateMode,
			ReceiverId: bot.UserId,
			IsBoot:     true,
		})
		if err != nil {
			fmt.Println(err)
		}

		if err := a.MessageUseCase.SendLogin(ctx, user.Id, &usecase.SendLogin{
			Ip:      ip,
			Agent:   userAgent[0],
			Address: address,
		}); err != nil {
			fmt.Println(err)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwt.GenerateToken("grpc-auth", a.Conf.App.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "grpc",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	if err := a.UserSession.Create(ctx, &postgresModel.UserSession{
		UserId:      user.Id,
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   userAgent[0],
	}); err != nil {
		fmt.Println(err)
	}

	return &authPb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.App.Jwt.ExpiresTime),
	}, nil
}
