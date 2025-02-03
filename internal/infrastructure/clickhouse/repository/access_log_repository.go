package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/infrastructure/clickhouse/model"
	"voo.su/pkg/clickhouseutil"
)

type AccessLogRepository struct {
	clickhouseutil.Repo[model.AccessLog]
}

func NewAccessLogRepository(clickHouse *clickHouseDriver.Conn) *AccessLogRepository {
	return &AccessLogRepository{Repo: clickhouseutil.NewRepo[model.AccessLog](clickHouse)}
}

func (a *AccessLogRepository) Create(ctx context.Context, accessLog *model.AccessLog) error {
	if err := a.Repo.Create(ctx, accessLog); err != nil {
		return fmt.Errorf("не удалось записать access log: %s", err)
	}

	return nil
}
