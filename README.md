## Logger interface with implementations (Zap and Logrus)

[![Build Status](https://api.travis-ci.com/arun0009/go-logger.svg?branch=master)](https://travis-ci.com/arun0009/go-logger)

You can use `logger` as an interface (example below) and set actual implementation to `ReplaceGlobals`, this allows 
you to change log library without changing your application code.

Also, when we create go libraries in general we shouldn't be logging but at times we do have to log, debug what the 
library is doing or trace the log. 

We cannot implement a library with one log library and expect applications to use the same log library. We use two 
of the popular log libraries [logrus](https://github.com/sirupsen/logrus) and [zap](https://github.com/uber-go/zap)
and this `go-logger` library allows you to use either one by using an interface. 

You can add your implementation if you want to add more log libraries (e.g. zerolog).

## Installation

go get -u github.com/arun0009/go-logger

## Quick Start

[logrus](https://github.com/sirupsen/logrus) example

```go
package main

import (
	"os"

	"github.com/arun0009/go-logger/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusLog := logrus.New()
	logrusLog.SetFormatter(&logrus.JSONFormatter{})
	logrusLog.SetOutput(os.Stdout)
	logrusLog.SetLevel(logrus.DebugLevel)
	log, _ := logger.NewLogrusLogger(logrusLog)
	logger.ReplaceGlobals(log)
        //anywhere in your code you can now use logger.L() as its globally set
	logger.L().WithFields(logger.Fields{
		"foo": "bar",
	}).Info("direct")
}
```

[zap](https://github.com/uber-go/zap) example

```go
package main

import (
	"os"

	"github.com/arun0009/go-logger/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	consoleEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder,
		zapcore.Lock(zapcore.AddSync(os.Stderr)),
		zapcore.DebugLevel)
	zapLogger := zap.New(core)
	log, _ := logger.NewZapLogger(zapLogger)
 	logger.ReplaceGlobals(log)
        //anywhere in your code you can now use logger.L() as its globally set
	logger.L().WithFields(logger.Fields{
		"foo": "bar",
	}).Info("direct")
}
```