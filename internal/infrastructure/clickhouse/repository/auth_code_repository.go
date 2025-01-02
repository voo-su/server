// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/infrastructure/clickhouse/model"
)

type AuthCodeRepository struct {
	ClickHouse clickHouseDriver.Conn
}

func NewAuthCodeRepository(clickHouse clickHouseDriver.Conn) *AuthCodeRepository {
	return &AuthCodeRepository{
		ClickHouse: clickHouse,
	}
}

func (a *AuthCodeRepository) Create(ctx context.Context, code *model.AuthCode) error {
	if err := a.ClickHouse.Exec(
		ctx,
		"INSERT INTO auth_codes (email, code, token, error_message) VALUES (?, ?, ?, ?)",
		code.Email,
		code.Code,
		code.Token,
		code.ErrorMessage,
	); err != nil {
		return fmt.Errorf("не удалось записать лог код аутентификации: %w", err)
	}

	return nil
}
