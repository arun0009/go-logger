package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	logger zerolog.Logger
	ctx    context.Context
}

func NewZerologLogger(logger zerolog.Logger) *ZerologLogger {
	return &ZerologLogger{logger: logger}
}

func (z *ZerologLogger) Debug(msg string, keyvals ...any) { z.log(zerolog.DebugLevel, msg, keyvals...) }
func (z *ZerologLogger) Info(msg string, keyvals ...any)  { z.log(zerolog.InfoLevel, msg, keyvals...) }
func (z *ZerologLogger) Warn(msg string, keyvals ...any)  { z.log(zerolog.WarnLevel, msg, keyvals...) }
func (z *ZerologLogger) Error(msg string, keyvals ...any) { z.log(zerolog.ErrorLevel, msg, keyvals...) }

func (z *ZerologLogger) WithContext(ctx context.Context) Logger {
	return &ZerologLogger{logger: z.logger, ctx: ctx}
}

func (z *ZerologLogger) log(level zerolog.Level, msg string, keyvals ...any) {
	e := z.logger.WithLevel(level)

	ctxFields := extractCtxFields(z.ctx)
	for i := 0; i < len(ctxFields); i += 2 {
		if i+1 < len(ctxFields) {
			if key, ok := ctxFields[i].(string); ok {
				e = e.Interface(key, ctxFields[i+1])
			}
		}
	}

	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			key, ok := keyvals[i].(string)
			if !ok {
				continue
			}
			e = e.Interface(key, keyvals[i+1])
		}
	}
	e.Msg(msg)
}
