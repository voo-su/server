package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type UserRepository struct {
	gormutil.Repo[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repo: gormutil.NewRepo[model.User](db)}
}

func (u *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := u.Repo.Db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (*model.User, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("пуст")
	}

	return u.Repo.FindByWhere(context.TODO(), "username = ?", username)
}

func (u *UserRepository) IsEmailExist(ctx context.Context, email string) bool {
	if len(email) == 0 {
		return false
	}

	exist, _ := u.Repo.QueryExist(ctx, "email = ?", email)
	return exist
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("пуст")
	}

	return u.Repo.FindByWhere(context.TODO(), "email = ?", email)
}

func (u *UserRepository) Search(q string, id int, limit int) ([]*model.User, error) {
	return u.Repo.FindAll(context.TODO(), func(db *gorm.DB) {
		query := "%" + q + "%"
		db.Where("id <> ? AND is_bot = 0 AND lower(username) LIKE lower(?) OR lower(name) LIKE lower(?) OR lower(surname) LIKE lower(?)", id, query, query, query).Limit(limit)
	})
}

func (u *UserRepository) SearchByUsername(username string, id int) ([]*model.User, error) {
	if len(username) == 0 {
		return u.Repo.FindAll(context.TODO(), func(db *gorm.DB) {
			db.Where("id <> ? AND is_bot = ?", id, 0)
		})
	}

	return u.Repo.FindAll(context.TODO(), func(db *gorm.DB) {
		db.Where("id <> ? AND lower(username) LIKE lower(?)", id, "%"+username+"%")
	})
}
