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
	absLogger.WithFields(Fields{
		"foo": "bar",
	}).Infof("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

//set logger to info and see that it doesn't print debug statements
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
	absLogger.WithFields(Fields{
		"foo": "bar",
	}).Debugf("direct")
	assert.Equal(t, "", string(buffer.Bytes()))
}
