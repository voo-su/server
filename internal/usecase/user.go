package usecase

import (
	"voo.su/internal/repository/repo"
)

type UserUseCase struct {
	*repo.Source
	UserRepo        *repo.User
	UserSessionRepo *repo.UserSession
}

func NewUserUseCase(
	source *repo.Source,
	userRepo *repo.User,
	userSessionRepo *repo.UserSession,
) *UserUseCase {
	return &UserUseCase{
		Source:          source,
		UserRepo:        userRepo,
		UserSessionRepo: userSessionRepo,
	}
}
