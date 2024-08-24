package log

import (
	"context"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler)
}

func Demo(ctx context.Context) {
	logger := NewLogger()

	logger.Info("hello, world")
	logger.Warn("hello, world", "foo", "bar")
	logger.WarnContext(ctx, "hello, world", "foo", "bar")

	logger.InfoContext(ctx, "hello, world")
	logger.WarnContext(ctx, "hello, world", slog.String("foo", "bar"))
}
