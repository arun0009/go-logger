package logger

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscard(t *testing.T) {
	l := Discard()
	assert.NotNil(t, l)
	// Should not panic or error
	l.Info("this should be discarded", "foo", "bar")
}

func TestNewMultiLogger(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	l1 := NewSlogLogger(slog.New(slog.NewJSONHandler(&buf1, nil)))
	l2 := NewSlogLogger(slog.New(slog.NewJSONHandler(&buf2, nil)))

	ml := NewMultiLogger(l1, l2)
	ml.Info("multi-log", "key", "value")

	assert.Contains(t, buf1.String(), "multi-log")
	assert.Contains(t, buf1.String(), "value")
	assert.Contains(t, buf2.String(), "multi-log")
	assert.Contains(t, buf2.String(), "value")
}
