package logger

import (
	"log/slog"
)

// NewMultiLogger creates a logger that fans out log records to multiple loggers.
// It utilizes slog.NewMultiHandler introduced in Go 1.26.
// Note: It expects SlogLogger instances to extract their handlers.
func NewMultiLogger(loggers ...Logger) Logger {
	handlers := make([]slog.Handler, 0, len(loggers))
	for _, l := range loggers {
		if sl, ok := l.(*SlogLogger); ok {
			handlers = append(handlers, sl.logger.Handler())
		}
		// For other adapters, we might need a way to get their handler or wrap them.
		// For now, we focus on the slog-native multi-handler.
	}

	if len(handlers) == 0 {
		return Discard()
	}

	return NewSlogLogger(slog.New(slog.NewMultiHandler(handlers...)))
}
