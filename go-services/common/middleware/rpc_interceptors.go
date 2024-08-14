package middleware

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type slogGrpcLogger struct {
	logger *slog.Logger
}

func (l *slogGrpcLogger) Log(ctx context.Context, level logging.Level, msg string, fields ...any) {
	// Convert logging.Level to slog.Level
	var slogLevel slog.Level
	switch level {
	case logging.LevelDebug:
		slogLevel = slog.LevelDebug
	case logging.LevelInfo:
		slogLevel = slog.LevelInfo
	case logging.LevelWarn:
		slogLevel = slog.LevelWarn
	case logging.LevelError:
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	l.logger.Log(ctx, slogLevel, msg, fields...)
}

func NewSlogGrpcLogger(logger *slog.Logger) *slogGrpcLogger {
	return &slogGrpcLogger{logger: logger}
}
func UnaryServerLogger(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(NewSlogGrpcLogger(logger))
}

func UnaryClientLogger(logger *slog.Logger) grpc.UnaryClientInterceptor {
	return logging.UnaryClientInterceptor(NewSlogGrpcLogger(logger))
}

func grpcPanicRecoveryHandler(p any) (err error) {
	slog.Error("recovered from panic", "panic", p)
	return status.Errorf(codes.Internal, "Internal server error: %v", p)
}

func UnaryRecoverer() grpc.UnaryServerInterceptor {
	return recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler))
}
