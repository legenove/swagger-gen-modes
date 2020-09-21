package services

import (
	"context"
	"github.com/legenove/nano-server-sdk/grpccore"
	"google.golang.org/grpc"
)

// 公共方法
func commonHandler(funcName string, handler grpc.UnaryHandler) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// before
		res, err := handler(ctx, req)
		// after
		return res, err
	}
}

func decoratorHandler(funcName string, handler grpc.UnaryHandler, decors ...grpccore.GrpcDecoratorFunc) grpc.UnaryHandler {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		handler = d(funcName, handler)
	}
	handler = commonHandler(funcName, handler)
	return grpccore.LoggerRecoveryHandler(funcName, handler)
}

