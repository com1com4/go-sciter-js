package window

import (
	"runtime"

	"github.com/com1com4/go-sciter-js"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/com1com4/go-sciter-js/issues/201
func init() {
	runtime.LockOSThread()
}
