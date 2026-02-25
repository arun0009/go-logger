package logger

import "sync"

var (
	_globalMU sync.RWMutex
	_globalL  Logger
)

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
// It's safe for concurrent use.
func L() Logger {
	_globalMU.RLock()
	l := _globalL
	_globalMU.RUnlock()
	return l
}

// ReplaceGlobals replaces the global Logger. It's safe for concurrent use.
func ReplaceGlobals(l Logger) {
	_globalMU.Lock()
	_globalL = l
	_globalMU.Unlock()
}
