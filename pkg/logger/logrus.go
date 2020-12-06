package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

//NewLogrusLogger create new logger using logrus logger
func NewLogrusLogger(logger *logrus.Logger) (Logger, error) {
	return &logrusLogger{
		logger: logger,
	}, nil
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *logrusLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *logrusLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *logrusLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *logrusLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *logrusLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *logrusLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *logrusLogEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *logrusLogEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *logrusLogEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *logrusLogEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *logrusLogEntry) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *logrusLogEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *logrusLogEntry) Debugf(template string, args ...interface{}) {
	l.entry.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *logrusLogEntry) Infof(template string, args ...interface{}) {
	l.entry.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *logrusLogEntry) Warnf(template string, args ...interface{}) {
	l.entry.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *logrusLogEntry) Errorf(template string, args ...interface{}) {
	l.entry.Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *logrusLogEntry) Panicf(template string, args ...interface{}) {
	l.entry.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *logrusLogEntry) Fatalf(template string, args ...interface{}) {
	l.entry.Fatalf(template, args...)
}

// WithFields adds fields to the logging context
func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

// convertToLogrusFields converts Fields to logrus.Fields
func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
