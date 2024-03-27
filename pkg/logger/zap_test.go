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
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"foo": "bar",
	}).Info("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestZapInfofLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"ping": "pong",
	}).Infof("received %s balls", "ping pong")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "received ping pong balls", fields["msg"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "pong", fields["ping"])
}

func TestZapWarnLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"foo": "bar",
		"log": "zap",
	}).Warn("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "warn", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "zap", fields["log"])
}

func TestZapWarnfLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.DebugLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"ping": "pong",
		"log":  "zap",
	}).Warnf("received %s balls", "table tennis")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "received table tennis balls", fields["msg"])
	assert.Equal(t, "warn", fields["level"])
	assert.Equal(t, "pong", fields["ping"])
	assert.Equal(t, "zap", fields["log"])
}

func TestZapPanicLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.ErrorLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		err := json.Unmarshal(buffer.Bytes(), &fields)
		assert.Nil(t, err)
		assert.Equal(t, "db not found", fields["msg"])
		assert.Equal(t, "panic", fields["level"])
		assert.Equal(t, "dataDB", fields["db"])
		assert.Equal(t, "zap", fields["log"])
	}()
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"db":  "dataDB",
		"log": "zap",
	}).Panic("db not found")
}

func TestZapErrorLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	writerSync := zapcore.AddSync(&buffer)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		writerSync,
		zap.InfoLevel,
	)
	z := zap.New(core)
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"acctNumber": 7899,
		"log":        "zap",
	}).Errorf("Error creating account %s", "testAccount")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account testAccount", fields["msg"])
	assert.Equal(t, "error", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "zap", fields["log"])
}

// set logger to info and see that it doesn't print debug statements
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
	absLogger, _ := NewZapLogger(z)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"foo": "bar",
	}).Debugf("direct")
	assert.Equal(t, "", string(buffer.Bytes()))
}
