package logger

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkAdapters(b *testing.B) {
	// Zap
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(io.Discard), zap.InfoLevel)
	zapLogger := NewZapLogger(zap.New(core))

	// Logrus
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})
	lr.SetOutput(io.Discard)
	logrusLogger := NewLogrusLogger(lr)

	// Slog
	sl := slog.New(slog.NewJSONHandler(io.Discard, nil))
	slogLogger := NewSlogLogger(sl)

	// Zerolog
	zl := zerolog.New(io.Discard)
	zerologLogger := NewZerologLogger(zl)

	b.Run("Zap", func(b *testing.B) {
		for b.Loop() {
			zapLogger.Info("test message", "key", "value", "int", 42)
		}
	})

	b.Run("Logrus", func(b *testing.B) {
		for b.Loop() {
			logrusLogger.Info("test message", "key", "value", "int", 42)
		}
	})

	b.Run("Slog", func(b *testing.B) {
		for b.Loop() {
			slogLogger.Info("test message", "key", "value", "int", 42)
		}
	})

	b.Run("Zerolog", func(b *testing.B) {
		for b.Loop() {
			zerologLogger.Info("test message", "key", "value", "int", 42)
		}
	})

	b.Run("Slog-WithContext", func(b *testing.B) {
		ctx := context.Background()
		logger := slogLogger.WithContext(ctx)
		for b.Loop() {
			logger.Info("test message", "key", "value", "int", 42)
		}
	})
}
