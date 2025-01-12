package clickhouseutil

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"reflect"
	"strings"
	"voo.su/pkg"
)

type ITable interface {
	TableName() string
}

type Repo[T ITable] struct {
	model      T
	ClickHouse clickHouseDriver.Conn
}

func NewRepo[T ITable](clickHouse *clickHouseDriver.Conn) Repo[T] {
	return Repo[T]{ClickHouse: *clickHouse}
}

func (r *Repo[T]) Create(ctx context.Context, model ITable) error {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("модель должна быть структурой или указателем на структуру")
	}

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		chTag := field.Tag.Get("ch")
		if chTag == "" {
			continue
		}

		columns = append(columns, chTag)
		placeholders = append(placeholders, "?")

		value := v.Field(i).Interface()

		if pkg.IsZeroValue(v.Field(i)) {
			values = append(values, nil)
		} else {
			values = append(values, value)
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", model.TableName(), strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	if err := r.ClickHouse.Exec(ctx, query, values...); err != nil {
		return fmt.Errorf("не удалось записать: %s", err)
	}

	return nil
}
