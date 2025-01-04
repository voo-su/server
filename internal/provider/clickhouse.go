package provider

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"voo.su/internal/config"
	"voo.su/pkg/locale"
)

func NewClickHouseClient(conf *config.Config, locale locale.ILocale) clickHouseDriver.Conn {
	conn, err := clickhouse.Open(conf.ClickHouse.Options())
	if err != nil {
		panic(fmt.Errorf(locale.Localize("connection_error"), "ClickHouse", err))
	}

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		if err != nil {
			panic(fmt.Errorf(locale.Localize("connection_error"), "ClickHouse", err))
		}
	}

	return conn
}
