package handler

import (
	"context"
	"fmt"
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
	AuthUseCase *usecase.AuthUseCase
}

func NewAccountHandler(
	locale locale.ILocale,
	userUseCase *usecase.UserUseCase,
	authUseCase *usecase.AuthUseCase,
) *Account {
	return &Account{
		Locale:      locale,
		UserUseCase: userUseCase,
		AuthUseCase: authUseCase,
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
	switch entity := in.Entity.Entity.(type) {
	case *accountPb.NotifyEntity_Chat:
		fmt.Println(entity.Chat.ChatId)
		// TODO
	case *accountPb.NotifyEntity_Group:
		fmt.Println(entity.Group.GroupId)
		// TODO
	case *accountPb.NotifyEntity_Chats:
		notify, err := a.UserUseCase.GetNotifySettings(ctx, uid)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.ChatsMuteUntil,
			ShowPreviews: notify.ChatsShowPreviews,
			Silent:       notify.ChatsSilent,
		}
	case *accountPb.NotifyEntity_Groups:
		notify, err := a.UserUseCase.GetNotifySettings(ctx, uid)
		if err != nil {
			return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
		}

		res.Settings = &accountPb.EntityNotifySettings{
			MuteUntil:    notify.GroupMuteUntil,
			ShowPreviews: notify.GroupShowPreviews,
			Silent:       notify.GroupSilent,
		}
	default:
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return res, nil
}

func (a *Account) UpdateNotifySettings(ctx context.Context, in *accountPb.UpdateNotifySettingsRequest) (*accountPb.UpdateNotifySettingsResponse, error) {
	//uid := grpcutil.UserId(ctx)
	//if err := a.UserUseCase.UpdateNotifySettings(ctx, uid, &entity.NotifySettings{
	//	PersonalChats: in.PersonalChats,
	//	GroupChats:    in.GroupChats,
	//}); err != nil {
	//	return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	//}

	switch entity := in.Entity.Entity.(type) {
	case *accountPb.NotifyEntity_Chat:
		fmt.Println(entity.Chat.ChatId)
		// TODO
	case *accountPb.NotifyEntity_Group:
		fmt.Println(entity.Group.GroupId)
		// TODO
	case *accountPb.NotifyEntity_Chats:
		// TODO
	case *accountPb.NotifyEntity_Groups:
		// TODO
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

	if err := a.UserUseCase.RegisterDevice(ctx, int64(uid), session.Id, in.TokenType, in.Token); err != nil {
		return nil, status.Error(codes.Unknown, a.Locale.Localize("general_error"))
	}

	return &accountPb.RegisterDeviceResponse{
		Success: true,
	}, nil
}
