package logger

import (
	"context"
	"log/slog"
)

type SlogLogger struct {
	logger *slog.Logger
	ctx    context.Context
}

func NewSlogLogger(logger *slog.Logger) *SlogLogger {
	return &SlogLogger{logger: logger}
}

func (s *SlogLogger) Debug(msg string, keyvals ...any) {
	s.logger.DebugContext(s.ctxOrDefault(), msg, s.toAttrs(keyvals)...)
}
func (s *SlogLogger) Info(msg string, keyvals ...any) {
	s.logger.InfoContext(s.ctxOrDefault(), msg, s.toAttrs(keyvals)...)
}
func (s *SlogLogger) Warn(msg string, keyvals ...any) {
	s.logger.WarnContext(s.ctxOrDefault(), msg, s.toAttrs(keyvals)...)
}
func (s *SlogLogger) Error(msg string, keyvals ...any) {
	s.logger.ErrorContext(s.ctxOrDefault(), msg, s.toAttrs(keyvals)...)
}

func (s *SlogLogger) WithContext(ctx context.Context) Logger {
	return &SlogLogger{logger: s.logger, ctx: ctx}
}

func (s *SlogLogger) ctxOrDefault() context.Context {
	if s.ctx != nil {
		return s.ctx
	}
	return context.Background()
}

func (s *SlogLogger) toAttrs(keyvals []any) []any {
	ctxFields := extractCtxFields(s.ctx)
	totalLen := len(ctxFields) + len(keyvals)
	attrs := make([]any, 0, totalLen/2)

	// Add context fields
	for i := 0; i < len(ctxFields); i += 2 {
		if i+1 < len(ctxFields) {
			if key, ok := ctxFields[i].(string); ok {
				attrs = append(attrs, slog.Any(key, ctxFields[i+1]))
			}
		}
	}

	// Add method fields
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			key, ok := keyvals[i].(string)
			if !ok {
				continue
			}
			attrs = append(attrs, slog.Any(key, keyvals[i+1]))
		}
	}
	return attrs
}
