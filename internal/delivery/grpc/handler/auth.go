package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
	authPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
)

type Auth struct {
	authPb.UnimplementedAuthServiceServer
	Conf             *config.Config
	Locale           locale.ILocale
	AuthUseCase      *usecase.AuthUseCase
	IpAddressUseCase *usecase.IpAddressUseCase
	ChatUseCase      *usecase.ChatUseCase
	BotUseCase       *usecase.BotUseCase
	MessageUseCase   usecase.IMessageUseCase
}

func NewAuthHandler(
	conf *config.Config,
	locale locale.ILocale,
	authUseCase *usecase.AuthUseCase,
	ipAddressUseCase *usecase.IpAddressUseCase,
	chatUseCase *usecase.ChatUseCase,
	botUseCase *usecase.BotUseCase,
	messageUseCase usecase.IMessageUseCase,
) *Auth {
	return &Auth{
		Conf:             conf,
		Locale:           locale,
		AuthUseCase:      authUseCase,
		IpAddressUseCase: ipAddressUseCase,
		ChatUseCase:      chatUseCase,
		BotUseCase:       botUseCase,
		MessageUseCase:   messageUseCase,
	}
}

func (a *Auth) Login(ctx context.Context, in *authPb.AuthLoginRequest) (*authPb.AuthLoginResponse, error) {
	token, err := a.AuthUseCase.Login(ctx, constant.GuardGrpcAuth, in.Email)
	if err != nil {
		log.Printf("AuthHandler Login: %v", err)
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

	claims, err := jwtutil.ParseToken(in.Token, a.Conf.App.Jwt.Secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, a.Locale.Localize("invalid_token"))
	}

	if claims.Guard != constant.GuardGrpcAuth || claims.Valid() != nil {
		return nil, status.Error(codes.Unauthenticated, a.Locale.Localize("invalid_token"))
	}

	ip := grpcutil.ClientIp(ctx)

	userAgent := grpcutil.UserAgent(ctx)

	user, err := a.AuthUseCase.Register(ctx, claims.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := a.AuthUseCase.Delete(ctx, constant.LoginChannel, in.Token); err != nil {
		fmt.Println(err)
	}

	bot, err := a.BotUseCase.BotRepo.GetLoginBot(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if bot != nil {
		_, err = a.ChatUseCase.Create(ctx, &usecase.CreateChatOpt{
			UserId:     user.Id,
			ChatType:   constant.ChatPrivateMode,
			ReceiverId: bot.UserId,
			IsBoot:     true,
		})
		if err != nil {
			fmt.Println(err)
		}

		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		if err := a.MessageUseCase.SendLogin(ctx, user.Id, &entity.SendLogin{
			Ip:      ip,
			Agent:   userAgent,
			Address: &address,
		}); err != nil {
			fmt.Println(err)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwtutil.GenerateToken(constant.GuardGrpcAuth, a.Conf.App.Jwt.Secret, &jwtutil.Options{
		ExpiresAt: jwtutil.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "grpc",
		IssuedAt:  jwtutil.NewNumericDate(time.Now()),
	})

	if err := a.AuthUseCase.UserSessionRepo.Create(ctx, &postgresModel.UserSession{
		UserId:      int64(user.Id),
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   userAgent,
	}); err != nil {
		fmt.Println(err)
	}

	return &authPb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   a.Conf.App.Jwt.ExpiresTime,
	}, nil
}

func (a *Auth) Logout(ctx context.Context, in *authPb.AuthLogoutRequest) (*authPb.AuthLogoutResponse, error) {
	token := grpcutil.UserToken(ctx)

	if err := a.AuthUseCase.Logout(ctx, token); err != nil {
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return &authPb.AuthLogoutResponse{
		Success: true,
	}, nil
}
