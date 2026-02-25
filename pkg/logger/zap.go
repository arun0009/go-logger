package logger

import (
	"context"

	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
	ctx    context.Context
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: logger.Sugar()}
}

func (z *ZapLogger) Debug(msg string, keyvals ...any) { z.log().Debugw(msg, keyvals...) }
func (z *ZapLogger) Info(msg string, keyvals ...any)  { z.log().Infow(msg, keyvals...) }
func (z *ZapLogger) Warn(msg string, keyvals ...any)  { z.log().Warnw(msg, keyvals...) }
func (z *ZapLogger) Error(msg string, keyvals ...any) { z.log().Errorw(msg, keyvals...) }

func (z *ZapLogger) WithContext(ctx context.Context) Logger {
	return &ZapLogger{logger: z.logger, ctx: ctx}
}

func (z *ZapLogger) log() *zap.SugaredLogger {
	fields := extractCtxFields(z.ctx)
	if len(fields) == 0 {
		return z.logger
	}
	return z.logger.With(fields...)
}
