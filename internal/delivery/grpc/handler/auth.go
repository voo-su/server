// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
	authPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/jwt"
	"voo.su/pkg/locale"
	"voo.su/pkg/utils"
)

type Auth struct {
	authPb.UnimplementedAuthServiceServer
	Conf              *config.Config
	Locale            locale.ILocale
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
	locale locale.ILocale,
	authUseCase *usecase.AuthUseCase,
	jwtTokenCacheRepo *redisRepo.JwtTokenCacheRepository,
	ipAddressUseCase *usecase.IpAddressUseCase,
	chatUseCase *usecase.ChatUseCase,
	botRepo *postgresRepo.BotRepository,
	userSession *postgresRepo.UserSessionRepository,
) *Auth {
	return &Auth{
		Conf:              conf,
		Locale:            locale,
		AuthUseCase:       authUseCase,
		JwtTokenCacheRepo: jwtTokenCacheRepo,
		IpAddressUseCase:  ipAddressUseCase,
		ChatUseCase:       chatUseCase,
		BotRepo:           botRepo,
		UserSession:       userSession,
	}
}

func (a *Auth) Login(ctx context.Context, in *authPb.AuthLoginRequest) (*authPb.AuthLoginResponse, error) {
	token, err := a.AuthUseCase.Login(ctx, constant.GuardGrpcAuth, in.Email)
	if err != nil {
		grpclog.Errorf("AuthHandler Login: %v", err)
		return nil, status.Error(codes.FailedPrecondition, a.Locale.Localize("token_creation_error"))
	}

	return &authPb.AuthLoginResponse{
		Token:     *token,
		ExpiresIn: constant.ExpiresTime,
	}, nil
}

func (a *Auth) Verify(ctx context.Context, in *authPb.AuthVerifyRequest) (*authPb.AuthVerifyResponse, error) {
	if !a.AuthUseCase.Verify(ctx, constant.LoginChannel, in.Token, in.Code) {
		return nil, status.Error(codes.Unauthenticated, a.Locale.Localize("invalid_code"))
	}

	claims, err := jwt.ParseToken(in.Token, a.Conf.App.Jwt.Secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, a.Locale.Localize("invalid_token"))
	}

	if claims.Guard != "grpc-auth" || claims.Valid() != nil {
		return nil, status.Error(codes.Unauthenticated, a.Locale.Localize("invalid_token"))
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
	token := jwt.GenerateToken(constant.GuardGrpcAuth, a.Conf.App.Jwt.Secret, &jwt.Options{
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
