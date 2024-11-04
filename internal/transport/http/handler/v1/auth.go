package v1

import (
	"fmt"
	"strconv"
	"time"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/jwt"
)

type Auth struct {
	Conf               *config.Config
	AuthUseCase        *usecase.AuthUseCase
	JwtTokenStorage    *cache.JwtTokenStorage
	IpAddressUseCase   *usecase.IpAddressUseCase
	ChatUseCase        *usecase.ChatUseCase
	BotRepo            *repo.Bot
	MessageSendUseCase usecase.MessageSendUseCase
	UserSession        *repo.UserSession
}

func (a *Auth) Login(ctx *core.Context) error {
	params := &v1Pb.AuthLoginRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(constant.ExpiresTime))
	token := jwt.GenerateToken("auth", a.Conf.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        params.Email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	if err := a.AuthUseCase.Send(ctx.Ctx(), constant.LoginChannel, params.Email, token); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&v1Pb.AuthLoginResponse{
		Token:     token,
		ExpiresIn: constant.ExpiresTime,
	})
}

func (a *Auth) Verify(ctx *core.Context) error {
	params := &v1Pb.AuthVerifyRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !a.AuthUseCase.Verify(ctx.Ctx(), constant.LoginChannel, params.Token, params.Code) {
		return ctx.InvalidParams("Неверный код")
	}

	claims, err := jwt.ParseToken(params.Token, a.Conf.Jwt.Secret)
	if err != nil {
		return ctx.InvalidParams("Неверный токен")
	}

	if claims.Guard != "auth" || claims.Valid() != nil {
		return ctx.InvalidParams("Неверный токен")
	}

	user, err := a.AuthUseCase.Register(ctx.Ctx(), claims.ID)
	ip := ctx.Context.ClientIP()

	user, err := a.AuthService.Register(ctx.Ctx(), claims.ID)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	a.AuthUseCase.Delete(ctx.Ctx(), constant.LoginChannel, params.Token)

	root, _ := a.BotRepo.GetLoginBot(ctx.Ctx())
	if root != nil {
		address, err := a.IpAddressUseCase.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		_, _ = a.ChatUseCase.Create(ctx.Ctx(), &usecase.CreateChatOpt{
			UserId:     user.Id,
			DialogType: constant.ChatPrivateMode,
			ReceiverId: root.UserId,
			IsBoot:     true,
		})
		_ = a.MessageSendUseCase.SendLogin(ctx.Ctx(), user.Id, &v1Pb.LoginMessageRequest{

		_ = a.MessageSendService.SendLogin(ctx.Ctx(), user.Id, &service.SendLogin{
			Ip:      ip,
			Agent:   ctx.Context.GetHeader("user-agent"),
			Address: address,
		})
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(a.Conf.Jwt.ExpiresTime))
	token := jwt.GenerateToken("api", a.Conf.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(user.Id),
		Issuer:    "web",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	_err := a.UserSession.Create(ctx.Ctx(), &model.UserSession{
		UserId:      user.Id,
		AccessToken: token,
		UserIp:      ip,
		UserAgent:   ctx.Context.GetHeader("user-agent"),
	})
	if _err != nil {
		fmt.Println(_err)
	}

	return ctx.Success(&v1Pb.AuthVerifyResponse{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   int32(a.Conf.Jwt.ExpiresTime),
	})
}

func (a *Auth) Logout(ctx *core.Context) error {
	session := ctx.JwtSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = a.JwtTokenStorage.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}

	return ctx.Success(nil)
}

//func (c *Auth) Refresh(ctx *core.Context) error {
//	c.toBlackList(ctx)
//	return ctx.Success(&v1Pb.AuthRefreshResponse{
//		Type:        "Bearer",
//		AccessToken: c.token(ctx.UserId()),
//		ExpiresIn:   int32(c.config.Jwt.ExpiresTime),
//	})
//}
