//go:build !windows

package glog

// shouldRegisterStderrSink determines whether we should register a log sink that writes to stderr.
// Today, this always returns true on non-Windows platforms, as it specifically checks for a
// condition that is only present on Windows.
func shouldRegisterStderrSink() bool { _ = "STUB: not implemented"; return false }

func lookupUser() string { _ = "STUB: not implemented"; return "" }
