package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/infrastructure/clickhouse/model"
	"voo.su/pkg/clickhouseutil"
)

type AccessGrpcLogRepository struct {
	clickhouseutil.Repo[model.AccessGrpcLog]
}

func NewAccessGrpcLogRepository(clickHouse *clickHouseDriver.Conn) *AccessGrpcLogRepository {
	return &AccessGrpcLogRepository{Repo: clickhouseutil.NewRepo[model.AccessGrpcLog](clickHouse)}
}

func (a *AccessGrpcLogRepository) Create(ctx context.Context, accessGrpcLog *model.AccessGrpcLog) error {
	if err := a.Repo.Create(ctx, accessGrpcLog); err != nil {
		return fmt.Errorf("не удалось записать access grpc log: %s", err)
	}

	return nil
}
