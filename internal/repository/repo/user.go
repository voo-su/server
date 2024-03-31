package repo

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type User struct {
	core.Repo[model.User]
}

func NewUser(db *gorm.DB) *User {
	return &User{Repo: core.NewRepo[model.User](db)}
}

func (u *User) Create(user *model.User) (*model.User, error) {
	if err := u.Repo.Db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) FindByUsername(username string) (*model.User, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("пуст")
	}

	return u.Repo.FindByWhere(context.TODO(), "username = ?", username)
}

func (u *User) IsEmailExist(ctx context.Context, email string) bool {
	if len(email) == 0 {
		return false
	}

	exist, _ := u.Repo.QueryExist(ctx, "email = ?", email)
	return exist
}

func (u *User) FindByEmail(email string) (*model.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("пуст")
	}

	return u.Repo.FindByWhere(context.TODO(), "email = ?", email)
}

func (u *User) SearchByUsername(username string, id int) ([]*model.User, error) {
	if len(username) == 0 {
		return u.Repo.FindAll(context.TODO(), func(db *gorm.DB) {
			db.Where("id <> ? AND is_bot = 0", id)
		})
	}

	return u.Repo.FindAll(context.TODO(), func(db *gorm.DB) {
		db.Where("id <> ? AND lower(username) LIKE lower(?)", id, "%"+username+"%")
	})
}
