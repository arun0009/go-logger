package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestZerologInfoLogger(t *testing.T) {
	var buffer bytes.Buffer
	logger := zerolog.New(&buffer)
	absLogger := NewZerologLogger(logger)
	ReplaceGlobals(absLogger)

	L().Info("direct", "foo", "bar")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["message"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestZerologWarnLogger(t *testing.T) {
	var buffer bytes.Buffer
	logger := zerolog.New(&buffer)
	absLogger := NewZerologLogger(logger)
	ReplaceGlobals(absLogger)

	L().Warn("direct", "foo", "bar", "log", "zerolog")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["message"])
	assert.Equal(t, "warn", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "zerolog", fields["log"])
}

func TestZerologErrorLogger(t *testing.T) {
	var buffer bytes.Buffer
	logger := zerolog.New(&buffer)
	absLogger := NewZerologLogger(logger)
	ReplaceGlobals(absLogger)

	L().Error("Error creating account", "acctNumber", 7899, "log", "zerolog")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account", fields["message"])
	assert.Equal(t, "error", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "zerolog", fields["log"])
}
