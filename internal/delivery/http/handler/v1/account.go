package v1

import (
	"encoding/json"
	"gorm.io/gorm"
	"regexp"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/internal/usecase"
	"voo.su/pkg/core"
	"voo.su/pkg/locale"
	"voo.su/pkg/logger"
	"voo.su/pkg/timeutil"
)

type Account struct {
	Locale      locale.ILocale
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
			return ctx.InvalidParams(a.Locale.Localize("invalid_birth_date_format"))
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
		return ctx.ErrorBusiness(a.Locale.Localize("personal_info_update_error"))
	}

	return ctx.Success(nil, a.Locale.Localize("personal_info_updated_success"))
}

func (a *Account) ChangeUsername(ctx *core.Context) error {
	params := &v1Pb.AccountUsernameUpdateRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", params.Username); !match {
		return ctx.ErrorBusiness(a.Locale.Localize("invalid_username_symbols"))
	}

	uid := ctx.UserId()
	var user model.User
	result := a.UserUseCase.UserRepo.Db.Where("username = ?", params.Username).First(&user)
	if result.Error != gorm.ErrRecordNotFound && user.Id != uid {
		return ctx.ErrorBusiness(a.Locale.Localize("username_already_exists"))
	}

	_, err := a.UserUseCase.UserRepo.UpdateById(ctx.Ctx(), ctx.UserId(), map[string]interface{}{
		"username": params.Username,
	})
	if err != nil {
		return ctx.ErrorBusiness(a.Locale.Localize("general_error"))
	}

	return ctx.Success(a.Locale.Localize("success"))
}

func (a *Account) Push(ctx *core.Context) error {
	params := &v1Pb.AccountPushRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	var in entity.WebPush
	if err := json.Unmarshal([]byte(params.Subscription), &in); err != nil {
		logger.Errorf("%s: %s", a.Locale.Localize("decode_error"), err)
		return ctx.ErrorBusiness(a.Locale.Localize("general_error"))
	}

	a.UserUseCase.WebPushInit(ctx.Ctx(), int64(uid), in)

	return ctx.Success(a.Locale.Localize("success"))
}
