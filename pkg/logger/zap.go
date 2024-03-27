package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// NewZapLogger create new logger using zap logger
func NewZapLogger(logger *zap.Logger) (Logger, error) {
	sugaredLogger := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	return &zapLogger{
		sugaredLogger: sugaredLogger,
	}, nil
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.sugaredLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}

// WithFields adds fields to the logging context
func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
