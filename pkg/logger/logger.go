package logger

import "context"

// Logger is the minimal interface your modules will depend on
type Logger interface {
	Debug(msg string, keyvals ...any)
	Info(msg string, keyvals ...any)
	Warn(msg string, keyvals ...any)
	Error(msg string, keyvals ...any)

	// WithContext attaches context for trace IDs, request IDs, etc.
	WithContext(ctx context.Context) Logger
}
