package logger

import (
	"log/slog"
)

// Discard returns a logger that discards all log records.
// It utilizes slog.DiscardHandler introduced in Go 1.24.
func Discard() Logger {
	return NewSlogLogger(slog.New(slog.DiscardHandler))
}
