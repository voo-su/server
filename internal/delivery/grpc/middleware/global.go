package middleware

import (
	"context"
	"log"
	"reflect"
)

type GlobalServiceKey string

var (
	_context context.Context

	AuthMiddlewareKey     = GlobalServiceKey("authMiddleware")
	GrpcMethodsServiceKey = GlobalServiceKey("grpcMethodsService")

	globalServicesMap = map[reflect.Type]GlobalServiceKey{
		reflect.TypeOf(&AuthMiddleware{}):    AuthMiddlewareKey,
		reflect.TypeOf(&GrpcMethodService{}): GrpcMethodsServiceKey,
	}
)

func RegisterGlobalService(ctx context.Context, service interface{}) context.Context {
	serviceType := reflect.TypeOf(service)
	if _, ok := globalServicesMap[serviceType]; !ok {
		log.Fatalf("Unknown global service: %v", serviceType)
	}

	ctx = context.WithValue(ctx, globalServicesMap[serviceType], service)
	_context = ctx
	return ctx
}

func GetGlobalService(k GlobalServiceKey) interface{} {
	v := _context.Value(k)
	if v == nil {
		log.Fatalf("Value not found: %v", k)
	}
	return v
}
