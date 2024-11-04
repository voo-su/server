package v1

import (
	"gorm.io/gorm"
	"regexp"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/repository/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/timeutil"
)

type Account struct {
	UserUseCase *usecase.UserUseCase
}

func (a *Account) Get(ctx *core.Context) error {
	user, err := a.UserUseCase.UserRepo.FindById(ctx.Ctx(), ctx.UserId())
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.AccountResponse{
		Id:       int32(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Name:     user.Name,
		Surname:  user.Surname,
		Gender:   int32(user.Gender),
		Birthday: user.Birthday,
		About:    user.About,
	})
}

func (a *Account) ChangeDetail(ctx *core.Context) error {
	params := &v1Pb.AccountDetailUpdateRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if params.Birthday != "" {
		if !timeutil.IsDateFormat(params.Birthday) {
			return ctx.InvalidParams("Неверный формат даты рождения")
		}
	}

	_, err := a.UserUseCase.UserRepo.UpdateById(ctx.Ctx(), ctx.UserId(), map[string]any{
		//"username": strings.TrimSpace(strings.Replace(params.Username, " ", "", -1)),
		"avatar":   params.Avatar,
		"name":     params.Name,
		"surname":  params.Surname,
		"gender":   params.Gender,
		"about":    params.About,
		"birthday": params.Birthday,
	})
	if err != nil {
		return ctx.ErrorBusiness("Ошибка при изменении личной информации")
	}

	return ctx.Success(nil, "Личная информация успешно изменена")
}

func (a *Account) ChangeUsername(ctx *core.Context) error {
	params := &v1Pb.AccountUsernameUpdateRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", params.Username); !match {
		return ctx.ErrorBusiness("Имя пользователя содержит недопустимые символы")
	}

	uid := ctx.UserId()
	var user model.User
	result := a.UserUseCase.UserRepo.Db.Where("username = ?", params.Username).First(&user)
	if result.Error != gorm.ErrRecordNotFound && user.Id != uid {
		return ctx.ErrorBusiness("Имя пользователя уже существует")
	}

	_, err := a.UserUseCase.UserRepo.UpdateById(ctx.Ctx(), ctx.UserId(), map[string]interface{}{
		"username": params.Username,
	})
	if err != nil {
		return ctx.ErrorBusiness("Ошибка")
	}

	return ctx.Success("Успешно")
}
