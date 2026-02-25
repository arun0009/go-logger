package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapInfoLogger(t *testing.T) {
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger := NewZapLogger(z)
	ReplaceGlobals(absLogger)

	L().Info("direct", "foo", "bar")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestZapWarnLogger(t *testing.T) {
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger := NewZapLogger(z)
	ReplaceGlobals(absLogger)

	L().Warn("direct", "foo", "bar", "log", "zap")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "warn", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "zap", fields["log"])
}

func TestZapErrorLogger(t *testing.T) {
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.InfoLevel,
	)
	z := zap.New(core)
	absLogger := NewZapLogger(z)
	ReplaceGlobals(absLogger)

	L().Error("Error creating account", "acctNumber", 7899, "log", "zap")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account", fields["msg"])
	assert.Equal(t, "error", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "zap", fields["log"])
}

func TestZapNoOutputLogger(t *testing.T) {
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.InfoLevel,
	)
	z := zap.New(core)
	absLogger := NewZapLogger(z)
	ReplaceGlobals(absLogger)

	L().Debug("direct", "foo", "bar")
	assert.Equal(t, "", buffer.String())
}
