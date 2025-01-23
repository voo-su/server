package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accountPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
)

type Account struct {
	accountPb.UnimplementedAccountServiceServer
	UserUseCase *usecase.UserUseCase
}

func NewAccountHandler(userUseCase *usecase.UserUseCase) *Account {
	return &Account{
		UserUseCase: userUseCase,
	}
}

func (a *Account) List(ctx context.Context, in *accountPb.GetAccountRequest) (*accountPb.GetAccountResponse, error) {
	uid := grpcutil.UserId(ctx)
	user, err := a.UserUseCase.UserRepo.FindById(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &accountPb.GetAccountResponse{
		Id:       int32(user.Id),
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
