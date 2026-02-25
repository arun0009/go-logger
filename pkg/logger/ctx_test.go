package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestContextPropagation(t *testing.T) {
	ctx := WithFields(context.Background(), "request_id", "req-123", "user_id", 456)

	t.Run("Zap", func(t *testing.T) {
		var buffer bytes.Buffer
		encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		core := zapcore.NewCore(encoder, zapcore.AddSync(&buffer), zap.InfoLevel)
		l := NewZapLogger(zap.New(core)).WithContext(ctx)

		l.Info("test")

		var fields map[string]any
		json.Unmarshal(buffer.Bytes(), &fields)
		assert.Equal(t, "req-123", fields["request_id"])
		assert.Equal(t, float64(456), fields["user_id"])
	})

	t.Run("Logrus", func(t *testing.T) {
		var buffer bytes.Buffer
		lr := logrus.New()
		lr.SetFormatter(&logrus.JSONFormatter{})
		lr.SetOutput(&buffer)
		l := NewLogrusLogger(lr).WithContext(ctx)

		l.Info("test")

		var fields map[string]any
		json.Unmarshal(buffer.Bytes(), &fields)
		assert.Equal(t, "req-123", fields["request_id"])
		assert.Equal(t, float64(456), fields["user_id"])
	})

	t.Run("Slog", func(t *testing.T) {
		var buffer bytes.Buffer
		sl := slog.New(slog.NewJSONHandler(&buffer, nil))
		l := NewSlogLogger(sl).WithContext(ctx)

		l.Info("test")

		var fields map[string]any
		json.Unmarshal(buffer.Bytes(), &fields)
		assert.Equal(t, "req-123", fields["request_id"])
		assert.Equal(t, float64(456), fields["user_id"])
	})

	t.Run("Zerolog", func(t *testing.T) {
		var buffer bytes.Buffer
		zl := zerolog.New(&buffer)
		l := NewZerologLogger(zl).WithContext(ctx)

		l.Info("test")

		var fields map[string]any
		json.Unmarshal(buffer.Bytes(), &fields)
		assert.Equal(t, "req-123", fields["request_id"])
		assert.Equal(t, float64(456), fields["user_id"])
	})
}
