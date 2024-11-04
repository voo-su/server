package usecase

import (
	"voo.su/internal/repository/repo"
)

type UserUseCase struct {
	*repo.Source
	UserRepo *repo.User
}

func NewUserUseCase(
	source *repo.Source,
	userRepo *repo.User,
) *UserUseCase {
	return &UserUseCase{
		Source:   source,
		UserRepo: userRepo,
	}
}
