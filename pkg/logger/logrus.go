package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

func NewLogrusLogger(logger *logrus.Logger) *LogrusLogger {
	return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Debug(msg string, keyvals ...any) { l.entry(keyvals...).Debug(msg) }
func (l *LogrusLogger) Info(msg string, keyvals ...any)  { l.entry(keyvals...).Info(msg) }
func (l *LogrusLogger) Warn(msg string, keyvals ...any)  { l.entry(keyvals...).Warn(msg) }
func (l *LogrusLogger) Error(msg string, keyvals ...any) { l.entry(keyvals...).Error(msg) }

func (l *LogrusLogger) WithContext(ctx context.Context) Logger {
	return &LogrusLogger{logger: l.logger, ctx: ctx}
}

func (l *LogrusLogger) entry(keyvals ...any) *logrus.Entry {
	fields := logrus.Fields{}

	ctxFields := extractCtxFields(l.ctx)
	for i := 0; i < len(ctxFields); i += 2 {
		if i+1 < len(ctxFields) {
			if key, ok := ctxFields[i].(string); ok {
				fields[key] = ctxFields[i+1]
			}
		}
	}

	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			if key, ok := keyvals[i].(string); ok {
				fields[key] = keyvals[i+1]
			}
		}
	}

	if l.ctx != nil {
		return l.logger.WithContext(l.ctx).WithFields(fields)
	}
	return l.logger.WithFields(fields)
}
