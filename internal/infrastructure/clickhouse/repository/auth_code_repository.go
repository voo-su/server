package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/infrastructure/clickhouse/model"
	"voo.su/pkg/clickhouseutil"
)

type AuthCodeRepository struct {
	clickhouseutil.Repo[model.AuthCode]
}

func NewAuthCodeRepository(clickHouse *clickHouseDriver.Conn) *AuthCodeRepository {
	return &AuthCodeRepository{Repo: clickhouseutil.NewRepo[model.AuthCode](clickHouse)}
}

func (a *AuthCodeRepository) Create(ctx context.Context, authCode *model.AuthCode) error {
	if err := a.Repo.Create(ctx, authCode); err != nil {
		return fmt.Errorf("не удалось записать лог код аутентификации: %s", err)
	}

	return nil
}
