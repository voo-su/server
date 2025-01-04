package repo

import (
	"context"
	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
}

type Repo[T ITable] struct {
	model T
	Db    *gorm.DB
}

func NewRepo[T ITable](db *gorm.DB) Repo[T] {
	return Repo[T]{Db: db}
}

func (r *Repo[T]) Model(ctx context.Context) *gorm.DB {
	return r.Db.WithContext(ctx).Model(r.model)
}

func (r *Repo[T]) FindById(ctx context.Context, id int) (*T, error) {
	var item *T
	err := r.Db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repo[T]) FindByIds(ctx context.Context, ids []int) ([]*T, error) {
	var items []*T
	err := r.Db.WithContext(ctx).Find(&items, ids).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repo[T]) FindAll(ctx context.Context, arg ...func(*gorm.DB)) ([]*T, error) {
	bd := r.Model(ctx)
	for _, fn := range arg {
		fn(bd)
	}
	var items []*T
	if err := bd.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repo[T]) FindByWhere(ctx context.Context, where string, args ...any) (*T, error) {
	var item *T
	err := r.Db.WithContext(ctx).Where(where, args...).First(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repo[T]) QueryCount(ctx context.Context, where string, args ...any) (int64, error) {
	var count int64
	err := r.Model(ctx).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repo[T]) QueryExist(ctx context.Context, where string, args ...any) (bool, error) {
	var count int64
	err := r.Model(ctx).Select("1").Where(where, args...).Limit(1).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *Repo[T]) UpdateById(ctx context.Context, id any, data map[string]any) (int64, error) {
	res := r.Model(ctx).Where("id = ?", id).Updates(data)

	return res.RowsAffected, res.Error
}

func (r *Repo[T]) UpdateWhere(ctx context.Context, data map[string]any, where string, args ...any) (int64, error) {
	res := r.Model(ctx).Where(where, args...).Updates(data)

	return res.RowsAffected, res.Error
}

func (r *Repo[T]) Create(ctx context.Context, data *T) error {
	return r.Db.WithContext(ctx).Create(data).Error
}

func (r *Repo[T]) BatchCreate(ctx context.Context, data []*T) error {
	return r.Db.WithContext(ctx).Create(data).Error
}

func (r *Repo[T]) Txx(ctx context.Context, fc func(tx *gorm.DB) error) error {
	return r.Db.WithContext(ctx).Transaction(fc)
}
