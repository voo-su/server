// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package v1

import (
	"fmt"
	"strconv"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/jwt"
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

func (a *Auth) Login(ctx *core.Context) error {
	params := &v1Pb.AuthLoginRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	token, err := a.AuthUseCase.Login(ctx.Ctx(), "auth", params.Email)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.AuthLoginResponse{
		Token:     *token,
		ExpiresIn: constant.ExpiresTime,
	})
}

func (a *Auth) Verify(ctx *core.Context) error {
	params := &v1Pb.AuthVerifyRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !a.AuthUseCase.Verify(ctx.Ctx(), constant.LoginChannel, params.Token, params.Code) {
		return ctx.InvalidParams(a.Locale.Localize("invalid_code"))
	}

	claims, err := jwt.ParseToken(params.Token, a.Conf.App.Jwt.Secret)
	if err != nil {
		return ctx.InvalidParams(a.Locale.Localize("invalid_code"))
	}

	if claims.Guard != "auth" || claims.Valid() != nil {
		return ctx.InvalidParams(a.Locale.Localize("invalid_code"))
	}

	user, err := a.AuthUseCase.Register(ctx.Ctx(), claims.ID)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	if err := a.AuthUseCase.Delete(ctx.Ctx(), constant.LoginChannel, params.Token); err != nil {
		fmt.Println(err)
	}

	ip := ctx.Context.ClientIP()
	bot, err := a.BotUseCase.BotRepo.GetLoginBot(ctx.Ctx())
	if err != nil {
		fmt.Println(err)
	}
	if bot != nil {
		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		_, err = a.ChatUseCase.Create(ctx.Ctx(), &usecase.CreateChatOpt{
			UserId:     user.Id,
			DialogType: constant.ChatPrivateMode,
			ReceiverId: bot.UserId,
			IsBoot:     true,
		})
		if err != nil {
			fmt.Println(err)
		}

		if err := a.MessageUseCase.SendLogin(ctx.Ctx(), user.Id, &usecase.SendLogin{
			Ip:      ip,
			Agent:   ctx.Context.GetHeader("user-agent"),
			Address: address,
		}); err != nil {
			fmt.Println(err)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwt.GenerateToken(constant.GuardHttpAuth, a.Conf.App.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "web",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	if err := a.UserUseCase.UserSessionRepo.Create(ctx.Ctx(), &postgresModel.UserSession{
		UserId:      user.Id,
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   ctx.Context.GetHeader("user-agent"),
	}); err != nil {
		fmt.Println(err)
	}

	return ctx.Success(&v1Pb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.App.Jwt.ExpiresTime),
	})
}

func (a *Auth) Logout(ctx *core.Context) error {
	session := ctx.JwtSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = a.AuthUseCase.JwtTokenCacheRepo.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}

	return ctx.Success(nil)
}

func (a *Auth) Refresh(ctx *core.Context) error {
	session := ctx.JwtSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = a.AuthUseCase.JwtTokenCacheRepo.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.App.Jwt.ExpiresTime))
	token := jwt.GenerateToken(constant.GuardHttpAuth, a.Conf.App.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(ctx.UserId()),
		Issuer:    "web",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return ctx.Success(&v1Pb.AuthRefreshResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.App.Jwt.ExpiresTime),
	})
}
