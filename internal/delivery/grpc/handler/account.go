package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	accountPb "voo.su/api/grpc/pb"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type Account struct {
	accountPb.UnimplementedAccountServiceServer
	Locale      locale.ILocale
	UserUseCase *usecase.UserUseCase
	AuthUseCase *usecase.AuthUseCase
	ChatUseCase *usecase.ChatUseCase
}

func NewAccountHandler(
	locale locale.ILocale,
	userUseCase *usecase.UserUseCase,
	authUseCase *usecase.AuthUseCase,
	chatUseCase *usecase.ChatUseCase,
) *Account {
	return &Account{
		Locale:      locale,
		UserUseCase: userUseCase,
		AuthUseCase: authUseCase,
		ChatUseCase: chatUseCase,
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

func (a *Account) GetNotifySettings(ctx context.Context, in *accountPb.GetNotifySettingsRequest) (*accountPb.GetNotifySettingsResponse, error) {
	uid := grpcutil.UserId(ctx)

	res := &accountPb.GetNotifySettingsResponse{}
	switch e := in.Entity.Entity.(type) {
	case *accountPb.NotifyEntity_Chat:
		chatId := e.Chat.ChatId
		notify, err := a.ChatUseCase.GetNotifySettings(ctx, constant.ChatPrivateMode, uid, chatId)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.MuteUntil,
			ShowPreviews: notify.ShowPreviews,
			Silent:       notify.Silent,
		}

	case *accountPb.NotifyEntity_Group:
		groupId := e.Group.GroupId
		notify, err := a.ChatUseCase.GetNotifySettings(ctx, constant.ChatGroupMode, uid, groupId)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.MuteUntil,
			ShowPreviews: notify.ShowPreviews,
			Silent:       notify.Silent,
		}
	case *accountPb.NotifyEntity_Chats:
		notify, err := a.UserUseCase.GetNotifySettings(ctx, constant.ChatPrivateMode, uid)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.MuteUntil,
			ShowPreviews: notify.ShowPreviews,
			Silent:       notify.Silent,
		}

	case *accountPb.NotifyEntity_Groups:
		notify, err := a.UserUseCase.GetNotifySettings(ctx, constant.ChatGroupMode, uid)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.MuteUntil,
			ShowPreviews: notify.ShowPreviews,
			Silent:       notify.Silent,
		}
	default:
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return res, nil
}

func (a *Account) UpdateNotifySettings(ctx context.Context, in *accountPb.UpdateNotifySettingsRequest) (*accountPb.UpdateNotifySettingsResponse, error) {
	uid := grpcutil.UserId(ctx)
	settings := &entity.NotifySettings{
		MuteUntil:    in.Settings.MuteUntil,
		Silent:       in.Settings.Silent,
		ShowPreviews: in.Settings.ShowPreviews,
	}

	switch e := in.Entity.Entity.(type) {
	case *accountPb.NotifyEntity_Chat:
		chatId := e.Chat.ChatId
		if err := a.ChatUseCase.UpdateNotifySettings(ctx, constant.ChatPrivateMode, uid, chatId, settings); err != nil {
			log.Println(err)
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

	case *accountPb.NotifyEntity_Group:
		groupId := e.Group.GroupId
		if err := a.ChatUseCase.UpdateNotifySettings(ctx, constant.ChatGroupMode, uid, groupId, settings); err != nil {
			log.Println(err)
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

	case *accountPb.NotifyEntity_Chats:
		if err := a.UserUseCase.UpdateNotifySettings(ctx, constant.ChatPrivateMode, uid, settings); err != nil {
			log.Println(err)
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

	case *accountPb.NotifyEntity_Groups:
		if err := a.UserUseCase.UpdateNotifySettings(ctx, constant.ChatGroupMode, uid, settings); err != nil {
			log.Println(err)
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

	default:
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return &accountPb.UpdateNotifySettingsResponse{
		Success: true,
	}, nil
}

func (a *Account) RegisterDevice(ctx context.Context, in *accountPb.RegisterDeviceRequest) (*accountPb.RegisterDeviceResponse, error) {
	uid := grpcutil.UserId(ctx)
	token := grpcutil.UserToken(ctx)
	session, err := a.AuthUseCase.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	if in.Token == "" {
		log.Printf("%s user: %d - token: %s", a.Locale.Localize("token_not_provided"), uid, token)
		return nil, status.Error(codes.InvalidArgument, a.Locale.Localize("token_not_provided"))
	}

	if err := a.UserUseCase.DevicePushInit(ctx, int64(uid), session.Id, in.TokenType, in.Token); err != nil {
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return &accountPb.RegisterDeviceResponse{
		Success: true,
	}, nil
}
