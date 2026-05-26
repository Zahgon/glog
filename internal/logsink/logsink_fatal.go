package logsink

import (
	"unsafe"
)

func fatalMessageStore(e savedEntry) {
	_ = "STUB: not implemented"
	// Only put a new one in if we haven't assigned before.
	return
}

var fatalMessage unsafe.Pointer // savedEntry stored with CompareAndSwapPointer

// FatalMessage returns the Meta and message contents of the first message
// logged with Fatal severity, or false if none has occurred.
func FatalMessage() (*Meta, []byte, bool) { _ = "STUB: not implemented"; return nil, nil, false }

// DoNotUseRacyFatalMessage is FatalMessage, but worse.
//
//go:norace
//go:nosplit
func DoNotUseRacyFatalMessage() (*Meta, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}
