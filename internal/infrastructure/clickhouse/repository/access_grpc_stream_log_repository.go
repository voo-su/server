package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/infrastructure/clickhouse/model"
	"voo.su/pkg/clickhouseutil"
)

type AccessGrpcStreamLogRepository struct {
	clickhouseutil.Repo[model.AccessGrpcStreamLog]
}

func NewAccessGrpcStreamLogRepository(clickHouse *clickHouseDriver.Conn) *AccessGrpcStreamLogRepository {
	return &AccessGrpcStreamLogRepository{Repo: clickhouseutil.NewRepo[model.AccessGrpcStreamLog](clickHouse)}
}

func (a *AccessGrpcStreamLogRepository) Create(ctx context.Context, accessGrpcStreamLog *model.AccessGrpcStreamLog) error {
	if err := a.Repo.Create(ctx, accessGrpcStreamLog); err != nil {
		return fmt.Errorf("не удалось записать access grpc stream log: %s", err)
	}

	return nil
}
