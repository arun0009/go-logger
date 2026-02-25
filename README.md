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

## Quick Start

### 1. Installation
```bash
go get github.com/arun0009/go-logger
```

### 2. Implementation (using Slog)
```go
package main

import (
	"log/slog"
	"os"

	"github.com/arun0009/go-logger/pkg/logger"
)

func main() {
	// Initialize your favorite backend
	handler := slog.NewJSONHandler(os.Stdout, nil)
	l := logger.NewSlogLogger(slog.New(handler))
	
	// Set as global (optional)
	logger.ReplaceGlobals(l)

	// Log with structure
	logger.L().Info("modern logging enabled", 
		"version", "Go 1.26",
		"env", "production",
	)
}
```

---

## Supported Adapters

| Backend | Performance | Package | Recommended Use |
| :--- | :--- | :--- | :--- |
| **Zerolog** |  Extreme | `github.com/rs/zerolog` | Low-latency microservices |
| **Zap** |  High | `go.uber.org/zap` | Enterprise-grade systems |
| **Slog** |  Standard | `log/slog` (Stdlib) | Default Go projects |
| **Logrus** |  Legacy | `github.com/sirupsen/logrus` | Maintaining older codebases |

### Example: Using Zerolog
```go
zl := zerolog.New(os.Stdout)
l := logger.NewZerologLogger(zl)
logger.ReplaceGlobals(l)

logger.L().Debug("zerolog is active", "speed", "fast")
```

---

## Context Support

Easily propagate request IDs or trace metadata:

```go
func HandleRequest(ctx context.Context) {
    // Attach context-specific fields
    l := logger.L().WithContext(ctx)
    
    l.Info("processing request", "path", "/api/data")
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