package middleware

import (
	"context"
	"log"
	"reflect"
)

type GlobalServiceKey string

var (
	_context context.Context

	AuthServiceKey        = GlobalServiceKey("authService")
	GrpcMethodsServiceKey = GlobalServiceKey("grpcMethodsService")

	globalServicesMap = map[reflect.Type]GlobalServiceKey{
		reflect.TypeOf(&TokenMiddleware{}):   AuthServiceKey,
		reflect.TypeOf(&GrpcMethodService{}): GrpcMethodsServiceKey,
	}
)

func RegisterGlobalService(ctx context.Context, service interface{}) context.Context {
	serviceType := reflect.TypeOf(service)
	if _, ok := globalServicesMap[serviceType]; !ok {
		log.Fatalf("Неизвестный глобальный сервис: %v", serviceType)
	}

	ctx = context.WithValue(ctx, globalServicesMap[serviceType], service)
	_context = ctx
	return ctx
}

func GetGlobalService(k GlobalServiceKey) interface{} {
	v := _context.Value(k)
	if v == nil {
		log.Fatalf("Значение не найдено: %v", k)
	}
	return v
}
