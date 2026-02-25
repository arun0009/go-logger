package logger

import "context"

type contextKey string

const logFieldsKey contextKey = "log_fields"

// WithFields returns a new context with the given keyvals.
func WithFields(ctx context.Context, keyvals ...any) context.Context {
	return context.WithValue(ctx, logFieldsKey, keyvals)
}

func extractCtxFields(ctx context.Context) []any {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(logFieldsKey)
	if kv, ok := v.([]any); ok {
		return kv
	}
	return nil
}
