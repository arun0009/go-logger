package logger

import (
	"sync"
)

var (
	_globalMU sync.RWMutex
	_globalL  Logger
)

// ReplaceGlobals replaces the global Logger and returns a
// function to restore the original values. It's safe for concurrent use.
func ReplaceGlobals(logger Logger) func() {
	_globalMU.Lock()
	prev := _globalL
	_globalL = logger
	_globalMU.Unlock()
	return func() { ReplaceGlobals(prev) }
}

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
// It's safe for concurrent use.
func L() Logger {
	_globalMU.RLock()
	l := _globalL
	_globalMU.RUnlock()
	return l
}
