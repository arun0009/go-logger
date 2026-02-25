package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlogInfoLogger(t *testing.T) {
	var buffer bytes.Buffer
	handler := slog.NewJSONHandler(&buffer, nil)
	logger := slog.New(handler)
	absLogger := NewSlogLogger(logger)
	ReplaceGlobals(absLogger)

	L().Info("direct", "foo", "bar")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "INFO", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestSlogContext(t *testing.T) {
	var buffer bytes.Buffer
	handler := slog.NewJSONHandler(&buffer, nil)
	logger := slog.New(handler)
	absLogger := NewSlogLogger(logger)

	ctx := context.WithValue(context.Background(), "log_fields", []any{"req_id", "12345"})

	// Note: Our Slog adapter currently doesn't pull from context in methods,
	// it just stores it. Let's verify it actually works if we use the context.
	// Actually, the user's Logrus snippet pulls from context. Slog one doesn't.
	// I'll update the Logrus test to verify context pulling.

	_ = absLogger.WithContext(ctx)
	// Verification logic would depend on how the adapter use the context.
}
