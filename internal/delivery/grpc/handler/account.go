package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accountPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type Account struct {
	accountPb.UnimplementedAccountServiceServer
	Locale      locale.ILocale
	UserUseCase *usecase.UserUseCase
}

func NewAccountHandler(locale locale.ILocale, userUseCase *usecase.UserUseCase) *Account {
	return &Account{
		Locale:      locale,
		UserUseCase: userUseCase,
	}
}

func (a *Account) GetProfile(ctx context.Context, in *accountPb.GetProfileRequest) (*accountPb.GetProfileResponse, error) {
	uid := grpcutil.UserId(ctx)
	user, err := a.UserUseCase.UserRepo.FindById(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &accountPb.GetProfileResponse{
		Id:       int64(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Name:     user.Name,
		Surname:  user.Surname,
		Gender:   int32(user.Gender),
		Birthday: user.Birthday,
		About:    user.About,
	}, nil
}

func (a *Account) UpdateProfile(ctx context.Context, in *accountPb.UpdateProfileRequest) (*accountPb.UpdateProfileResponse, error) {
	uid := grpcutil.UserId(ctx)

	if in.Birthday != "" {
		if !timeutil.IsDateFormat(in.Birthday) {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("invalid_birth_date_format"))
		}
	}

	_, err := a.UserUseCase.UserRepo.UpdateById(ctx, uid, map[string]any{
		"name":     in.Name,
		"surname":  in.Surname,
		"gender":   in.Gender,
		"about":    in.About,
		"birthday": in.Birthday,
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, a.Locale.Localize("personal_info_update_error"))
	}

	return &accountPb.UpdateProfileResponse{
		Success: true,
	}, nil
}

func (a *Account) UpdateProfilePhoto(ctx context.Context, in *accountPb.UpdateProfilePhotoRequest) (*accountPb.UpdateProfilePhotoResponse, error) {
	// TODO
	return &accountPb.UpdateProfilePhotoResponse{}, nil
}

func (a *Account) getNotifySettings(ctx context.Context, in *accountPb.GetNotifySettingsRequest) (*accountPb.GetNotifySettingsResponse, error) {
	// TODO
	return &accountPb.GetNotifySettingsResponse{}, nil
}

func (a *Account) updateNotifySettings(ctx context.Context, in *accountPb.UpdateNotifySettingsRequest) (*accountPb.UpdateNotifySettingsResponse, error) {
	// TODO
	return &accountPb.UpdateNotifySettingsResponse{}, nil
}

func (a *Account) RegisterDevice(ctx context.Context, in *accountPb.RegisterDeviceRequest) (*accountPb.RegisterDeviceResponse, error) {
	// TODO
	return &accountPb.RegisterDeviceResponse{}, nil
}
