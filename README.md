# ⚡ go-logger

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Performance](https://img.shields.io/badge/Performance-High-blueviolet?style=for-the-badge)](README.md#benchmarks)

A modernized, high-performance, and context-aware structured logging interface for Go. Designed to be modular and pluggable, allowing you to swap backends without changing a single line of application code.

---

## ✨ Features

- **Lightning Fast**: Blazing performance with first-class **Zerolog** and **Zap** support.
- **Structured by Default**: Uses modern `keyvals ...any` patterns (compatible with `slog`).
- **Context Aware**: Native `WithContext` support for request tracing and spans.
- **Pluggable Backends**: Out-of-the-box adapters for **Zap**, **Logrus**, **Slog**, and **Zerolog**.
- **Go 1.26 Native**: Utilizes `slog.NewMultiHandler` and `slog.DiscardHandler` for built-in multi-logging and silencing.
- **Developer Friendly**: Support for both Dependency Injection and Global Singleton patterns.

---

## Installation
```bash
go get github.com/arun0009/go-logger
```

---

## Supported Adapters

| Backend | Performance | Recommended Use |
| :--- | :--- | :--- |
| **Zerolog** | Extreme | Low-latency microservices |
| **Zap** | High | Enterprise-grade systems |
| **Slog** | Standard | Default Go projects |
| **Logrus** | Legacy | Maintaining older codebases |

### Backend Examples

#### Zerolog
```go
import (
    "os"
    "github.com/rs/zerolog"
    "github.com/arun0009/go-logger/pkg/logger"
)

zl := zerolog.New(os.Stdout)
l := logger.NewZerologLogger(zl)
logger.ReplaceGlobals(l)

// Usage
logger.L().Info("using zerolog", "speed", "extreme")
```

#### Zap
```go
import (
    "go.uber.org/zap"
    "github.com/arun0009/go-logger/pkg/logger"
)

z, _ := zap.NewProduction()
l := logger.NewZapLogger(z)
logger.ReplaceGlobals(l)

// Usage
logger.L().Info("using zap", "type", "structured")
```

#### Slog (Go 1.21+)
```go
import (
    "log/slog"
    "os"
    "github.com/arun0009/go-logger/pkg/logger"
)

sl := slog.New(slog.NewJSONHandler(os.Stdout, nil))
l := logger.NewSlogLogger(sl)
logger.ReplaceGlobals(l)

// Usage
logger.L().Info("using slog", "source", "standard library")
```

#### Logrus
```go
import (
    "github.com/sirupsen/logrus"
    "github.com/arun0009/go-logger/pkg/logger"
)

lr := logrus.New()
l := logger.NewLogrusLogger(lr)
logger.ReplaceGlobals(l)

// Usage
logger.L().Info("using logrus", "mode", "compatibility")
```

---

## Context & Metadata

`go-logger` supports seamless metadata propagation (e.g., Trace IDs, Request IDs) via the standard `context.Context`.

### 1. Attaching Metadata
Use `logger.WithFields` to inject metadata into a context.

```go
ctx := logger.WithFields(context.Background(), 
    "trace_id", "12345", 
    "user_id", 42,
)
```

### 2. Logging with Context
The logger automatically extracts and includes these fields in every log entry.

```go
func HandleRequest(ctx context.Context) {
    // Automatically includes trace_id, user_id, etc.
    logger.L().WithContext(ctx).Info("processing request")
}
```

---

## Benchmarks

*Results from Apple M4 (Go 1.26)*

| Adapter | ns/op | Operations/s |
| :--- | :--- | :--- |
| **Zerolog** | **188.2** | 6.3M |
| **Zap** | 283.2 | 3.9M |
| **Slog** | 388.5 | 3.1M |
| **Logrus** | 1057.0 | 1.0M |

---

##  License
MIT License. See [LICENSE](LICENSE) for details.