package repository

import (
	"context"
	"fmt"
	"voo.su/internal/infrastructure/clickhouse/model"
	"voo.su/pkg/clickhouseutil"

	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type LoggerRepository struct {
	clickhouseutil.Repo[model.Logger]
}

func NewLoggerRepository(clickHouse *clickHouseDriver.Conn) *LoggerRepository {
	return &LoggerRepository{Repo: clickhouseutil.NewRepo[model.Logger](clickHouse)}
}

func (l *LoggerRepository) Create(ctx context.Context, logger *model.Logger) error {
	if err := l.Repo.Create(ctx, logger); err != nil {
		return fmt.Errorf("не удалось записать: %s", err)
	}

	return nil
}
