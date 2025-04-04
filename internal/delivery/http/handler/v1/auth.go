package v1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
)

type Auth struct {
	Conf             *config.Config
	Locale           locale.ILocale
	AuthUseCase      *usecase.AuthUseCase
	IpAddressUseCase *usecase.IpAddressUseCase
	ChatUseCase      *usecase.ChatUseCase
	BotUseCase       *usecase.BotUseCase
	UserUseCase      *usecase.UserUseCase
	MessageUseCase   usecase.IMessageUseCase
}

func (a *Auth) Login(ctx *ginutil.Context) error {
	params := &v1Pb.AuthLoginRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	token, err := a.AuthUseCase.Login(ctx.Ctx(), constant.GuardHttpAuth, params.Email)
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.AuthLoginResponse{
		Token:     *token,
		ExpiresIn: constant.ExpiresTime,
	})
}

func (a *Auth) Verify(ctx *ginutil.Context) error {
	params := &v1Pb.AuthVerifyRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !a.AuthUseCase.Verify(ctx.Ctx(), constant.LoginChannel, params.Token, params.Code) {
		return ctx.InvalidParams(a.Locale.Localize("invalid_code"))
	}

	claims, err := jwtutil.ParseToken(params.Token, a.Conf.App.Jwt.Secret)
	if err != nil {
		return ctx.InvalidParams(a.Locale.Localize("invalid_token"))
	}

	if claims.Guard != constant.GuardHttpAuth || claims.Valid() != nil {
		return ctx.InvalidParams(a.Locale.Localize("invalid_token"))
	}

	user, err := a.AuthUseCase.Register(ctx.Ctx(), claims.ID)
	if err != nil {
		return ctx.Error(err.Error())
	}

	if err := a.AuthUseCase.Delete(ctx.Ctx(), constant.LoginChannel, params.Token); err != nil {
		log.Printf("Error - Verify - AuthUseCase.Delete: %v", err)
	}

	ip := ctx.Context.ClientIP()
	bot, err := a.BotUseCase.BotRepo.GetLoginBot(ctx.Ctx())
	if err != nil {
		log.Printf("Error - Verify - BotRepo.GetLoginBot: %v", err)
	}
	if bot != nil {
		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			log.Printf("Error - Verify - IpAddressUseCase.FindAddress: %v", err)
		}

		_, err = a.ChatUseCase.Create(ctx.Ctx(), &usecase.CreateChatOpt{
			UserId:     user.Id,
			ChatType:   constant.ChatPrivateMode,
			ReceiverId: bot.UserId,
			IsBoot:     true,
		})
		if err != nil {
			log.Printf("Error - Verify - ChatUseCase.Create: %v", err)
		}

		if err := a.MessageUseCase.SendLogin(ctx.Ctx(), user.Id, &entity.SendLogin{
			Ip:      ip,
			Agent:   ctx.Context.GetHeader("user-agent"),
			Address: address,
		}); err != nil {
			log.Printf("Error - Verify - MessageUseCase.SendLogin: %v", err)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwtutil.GenerateToken(constant.GuardHttpAuth, a.Conf.App.Jwt.Secret, &jwtutil.Options{
		ExpiresAt: jwtutil.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "web",
		IssuedAt:  jwtutil.NewNumericDate(time.Now()),
	})

	if err := a.UserUseCase.UserSessionRepo.Create(ctx.Ctx(), &postgresModel.UserSession{
		UserId:      int64(user.Id),
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   ctx.Context.GetHeader("user-agent"),
	}); err != nil {
		log.Printf("Error - Verify - UserSessionRepo.Create: %v", err)
	}

	return ctx.Success(&v1Pb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.App.Jwt.ExpiresTime),
	})
}

func (a *Auth) Logout(ctx *ginutil.Context) error {
	token := ctx.UserToken()
	if err := a.AuthUseCase.Logout(ctx.Ctx(), token); err != nil {
		return ctx.Error(status.Error(codes.Unknown, a.Locale.Localize("general_error")))
	}

	return ctx.Success(nil)
}

func (a *Auth) Refresh(ctx *ginutil.Context) error {
	session := ctx.JwtSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = a.AuthUseCase.JwtTokenCacheRepo.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwtutil.GenerateToken(constant.GuardHttpAuth, a.Conf.App.Jwt.Secret, &jwtutil.Options{
		ExpiresAt: jwtutil.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(ctx.UserId()),
		Issuer:    "web",
		IssuedAt:  jwtutil.NewNumericDate(time.Now()),
	})

	return ctx.Success(&v1Pb.AuthRefreshResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.App.Jwt.ExpiresTime),
	})
}
