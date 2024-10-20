package v1

import (
	"fmt"
	"strconv"
	"time"
	"voo.su/api/pb/v1"
	"voo.su/internal/config"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/pkg/core"
	"voo.su/pkg/jwt"
)

type Auth struct {
	Conf               *config.Config
	AuthService        *service.AuthService
	JwtTokenStorage    *cache.JwtTokenStorage
	RedisLock          *cache.RedisLock
	IpAddressService   *service.IpAddressService
	ChatService        *service.ChatService
	BotRepo            *repo.Bot
	MessageSendService service.MessageSendService
	UserSession        *repo.UserSession
}

func (a *Auth) Login(ctx *core.Context) error {
	params := &api_v1.AuthLoginRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(entity.ExpiresTime))
	token := jwt.GenerateToken("auth", a.Conf.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        params.Email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	err := a.AuthService.Send(ctx.Ctx(), entity.LoginChannel, params.Email, token)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(&api_v1.AuthLoginResponse{
		Token:     token,
		ExpiresIn: entity.ExpiresTime,
	})
}

func (a *Auth) Verify(ctx *core.Context) error {
	params := &api_v1.AuthVerifyRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if !a.AuthService.Verify(ctx.Ctx(), entity.LoginChannel, params.Token, params.Code) {
		return ctx.InvalidParams("Неверный код")
	}

	claims, err := jwt.ParseToken(params.Token, a.Conf.Jwt.Secret)
	if err != nil {
		return ctx.InvalidParams("Неверный токен")
	}

	if claims.Guard != "auth" || claims.Valid() != nil {
		return ctx.InvalidParams("Неверный токен")
	}

	user, err := a.AuthService.Register(ctx.Ctx(), claims.ID)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	a.AuthService.Delete(ctx.Ctx(), entity.LoginChannel, params.Token)

	ip := ctx.Context.ClientIP()
	root, _ := a.BotRepo.GetLoginBot(ctx.Ctx())
	if root != nil {
		address, err := a.IpAddressService.FindAddress(ip)
		if err != nil {
			fmt.Println(err)
		}

		_, _ = a.ChatService.Create(ctx.Ctx(), &service.CreateChatOpt{
			UserId:     user.Id,
			DialogType: entity.ChatPrivateMode,
			ReceiverId: root.UserId,
			IsBoot:     true,
		})
		_ = a.MessageSendService.SendLogin(ctx.Ctx(), user.Id, &api_v1.LoginMessageRequest{
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

	return ctx.Success(&api_v1.AuthVerifyResponse{
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
//	return ctx.Success(&api_v1.AuthRefreshResponse{
//		Type:        "Bearer",
//		AccessToken: c.token(ctx.UserId()),
//		ExpiresIn:   int32(c.config.Jwt.ExpiresTime),
//	})
//}
