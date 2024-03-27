package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogrusInfoLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger, _ := NewLogrusLogger(logrus)
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

func TestLogrusInfofLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger, _ := NewLogrusLogger(logrus)
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

func TestLogrusWarnLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"foo": "bar",
		"log": "logrus",
	}).Warn("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "warning", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "logrus", fields["log"])
}

func TestLogrusWarnfLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"ping": "pong",
		"log":  "logrus",
	}).Warnf("received %s balls", "table tennis")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "received table tennis balls", fields["msg"])
	assert.Equal(t, "warning", fields["level"])
	assert.Equal(t, "pong", fields["ping"])
	assert.Equal(t, "logrus", fields["log"])
}

func TestLogrusPanicLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.ErrorLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		err := json.Unmarshal(buffer.Bytes(), &fields)
		assert.Nil(t, err)
		assert.Equal(t, "db not found", fields["msg"])
		assert.Equal(t, "panic", fields["level"])
		assert.Equal(t, "dataDB", fields["db"])
		assert.Equal(t, "logrus", fields["log"])
	}()
	L().WithFields(Fields{
		"db":  "dataDB",
		"log": "logrus",
	}).Panic("db not found")
}

func TestLogursErrorLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.DebugLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"acctNumber": 7899,
		"log":        "logrus",
	}).Errorf("Error creating account %s", "testAccount")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account testAccount", fields["msg"])
	assert.Equal(t, "error", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "logrus", fields["log"])
}

// set logger to info and see that it doesn't print debug statements
func TestLogrusNoOutputLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.InfoLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	ReplaceGlobals(absLogger)
	L().WithFields(Fields{
		"foo": "bar",
	}).Debugf("direct")

	assert.Equal(t, "", string(buffer.Bytes()))
}
