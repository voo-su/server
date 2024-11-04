package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
	"time"
	authPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/transport/grpc/middleware"
	"voo.su/internal/usecase"
	"voo.su/pkg/jwt"
)

type AuthHandler struct {
	authPb.UnimplementedAuthServiceServer
	Conf               *config.Config
	TokenMiddleware    *middleware.TokenMiddleware
	AuthUseCase        *usecase.AuthUseCase
	JwtTokenStorage    *cache.JwtTokenStorage
	IpAddressUseCase   *usecase.IpAddressUseCase
	ChatUseCase        *usecase.ChatUseCase
	BotRepo            *repo.Bot
	MessageSendUseCase usecase.MessageSendUseCase
	UserSession        *repo.UserSession
}

func NewAuthHandler(
	conf *config.Config,
	tokenMiddleware *middleware.TokenMiddleware,
	authUseCase *usecase.AuthUseCase,
	jwtTokenStorage *cache.JwtTokenStorage,
	ipAddressUseCase *usecase.IpAddressUseCase,
	chatUseCase *usecase.ChatUseCase,
	botRepo *repo.Bot,
	messageSendUseCase usecase.MessageSendUseCase,
	userSession *repo.UserSession,
) *AuthHandler {
	return &AuthHandler{
		Conf:               conf,
		TokenMiddleware:    tokenMiddleware,
		AuthUseCase:        authUseCase,
		JwtTokenStorage:    jwtTokenStorage,
		IpAddressUseCase:   ipAddressUseCase,
		ChatUseCase:        chatUseCase,
		BotRepo:            botRepo,
		MessageSendUseCase: messageSendUseCase,
		UserSession:        userSession,
	}
}

func (a *AuthHandler) Login(ctx context.Context, in *authPb.AuthLoginRequest) (*authPb.AuthLoginResponse, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(constant.ExpiresTime))
	token := jwt.GenerateToken("grpc-auth", a.Conf.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        in.Email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	if err := a.AuthUseCase.Send(ctx, constant.LoginChannel, in.Email, token); err != nil {
		grpclog.Errorf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return &authPb.AuthLoginResponse{
		Token:     token,
		ExpiresIn: constant.ExpiresTime,
	}, nil
}

func getClientIP(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Попробуем извлечь IP из метаданных
		if ipList := md["x-forwarded-for"]; len(ipList) > 0 {
			return ipList[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		addr := p.Addr.String()
		// Извлекаем только IP-адрес
		return strings.Split(addr, ":")[0]
	}

	return ""
}

func (a *AuthHandler) Verify(ctx context.Context, in *authPb.AuthVerifyRequest) (*authPb.AuthVerifyResponse, error) {
	if !a.AuthUseCase.Verify(ctx, constant.LoginChannel, in.Token, in.Code) {
		return nil, status.Error(codes.Unauthenticated, "Неверный код")
	}

	claims, err := jwt.ParseToken(in.Token, a.Conf.Jwt.Secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный токен")
	}

	if claims.Guard != "grpc-auth" || claims.Valid() != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный токен")
	}

	ip := getClientIP(ctx)

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

	a.AuthUseCase.Delete(ctx, constant.LoginChannel, in.Token)

	root, _ := a.BotRepo.GetLoginBot(ctx)
	if root != nil {
		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		_, _ = a.ChatUseCase.Create(ctx, &usecase.CreateChatOpt{
			UserId:     user.Id,
			DialogType: constant.ChatPrivateMode,
			ReceiverId: root.UserId,
			IsBoot:     true,
		})

		_ = a.MessageSendUseCase.SendLogin(ctx, user.Id, &usecase.SendLogin{
			Ip:      ip,
			Agent:   userAgent[0],
			Address: address,
		})
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.Jwt.ExpiresTime))
	token := jwt.GenerateToken("grpc-auth", a.Conf.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "grpc",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	_err := a.UserSession.Create(ctx, &model.UserSession{
		UserId:      user.Id,
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   userAgent[0],
	})
	if _err != nil {
		fmt.Println(_err)
	}

	return &authPb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.Jwt.ExpiresTime),
	}, nil
}
