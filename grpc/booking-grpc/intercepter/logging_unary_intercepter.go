package intercepter

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func UnaryServerLoggingIntercepter(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		res, err := handler(ctx, req)

		end := time.Since(start)
		method := info.FullMethod

		logger.Info("Unary call has completed", zap.String("method", method), zap.String("duration", end.String()))
		return res, err
	}
}
