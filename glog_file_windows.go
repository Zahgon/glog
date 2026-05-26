//go:build windows

package glog

// shouldRegisterStderrSink determines whether we should register a log sink that writes to stderr.
// Today, this checks if stderr is "valid", in that it maps to a non-NULL Handle.
// Windows Services are spawned without Stdout and Stderr, so any attempt to use them equates to
// referencing an invalid file Handle.
// os.Stderr's FD is derived from a call to `syscall.GetStdHandle(syscall.STD_ERROR_HANDLE)`.
// Documentation[1] for the GetStdHandle function indicates the return value may be NULL if the
// application lacks the standard handle, so consider Stderr valid if its FD is non-NULL.
// [1]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
func shouldRegisterStderrSink() bool { _ = "STUB: not implemented"; return false }

// This follows the logic in the standard library's user.Current() function, except
// that it leaves out the potentially expensive calls required to look up the user's
// display name in Active Directory.
func lookupUser() string { _ = "STUB: not implemented"; return "" }
