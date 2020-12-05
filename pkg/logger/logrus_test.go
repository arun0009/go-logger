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
func TestLogrusNoOutputLogger(t *testing.T) {
	var buffer bytes.Buffer
	logrus := log.New()
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(&buffer)
	logrus.SetLevel(log.InfoLevel)
	absLogger, _ := NewLogrusLogger(logrus)
	absLogger.WithFields(Fields{
		"foo": "bar",
	}).Debugf("direct")

	assert.Equal(t, "", string(buffer.Bytes()))
}
