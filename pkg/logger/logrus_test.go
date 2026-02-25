package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogrusInfoLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)

	L().Info("direct", "foo", "bar")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "info", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestLogrusWarnLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)

	L().Warn("direct", "foo", "bar", "log", "logrus")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "warning", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "logrus", fields["log"])
}

func TestLogrusErrorLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)

	L().Error("Error creating account", "acctNumber", 7899, "log", "logrus")

	var fields map[string]any
	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account", fields["msg"])
	assert.Equal(t, "error", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "logrus", fields["log"])
}

func TestLogrusNoOutputLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.InfoLevel)
	absLogger := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)

	L().Debug("direct", "foo", "bar")

	assert.Equal(t, "", buffer.String())
}
