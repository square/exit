/*
Package exit defines semantic exit codes which may be used by Command Line tools
to aid in debugging and instrumentation.

This package defines exit codes in two ranges. Exit Codes 80-99 indicate a user error
of some sort. Exit Codes 100-119 indicate software or system error of some sort.
*/
package exit

import (
	"errors"
	"fmt"
)

// Code is the exit code that is passed to the system call `exit`
// when the program terminates. Conventionally, the value zero indicates
// success and all other values (1-255) indicate failure.
type Code = int

const (
	// OK indicates that the program exited successfully.
	OK Code = 0

	// NotOK indicates that the program exited unsuccessfully
	// but gives no extra context as to what the failure was.
	NotOK Code = 1
)

// Exit Codes 80-99 are reserved for user errors.
const (
	// UsageError indicates that the program exited unsuccessfully
	// because it was used incorrectly.
	//
	// Examples: a required argument was omitted or an invalid value
	// was supplied for a flag.
	UsageError Code = 80

	// UnknownSubcommand indicates that the program exited unsuccessfully
	// because an unrecognized subcommand was invoked.
	//
	// This is intended for CLI multi-tools. When you run a command that
	// doesn't exist from the shell, the shell exits 127. This is distinct
	// from that value in that the command itself exists but the subcommand
	// does not (e.g. `git nope` could exit 81).
	UnknownSubcommand Code = 81

	// RequirementNotMet indicates that the program exited unsuccessfully
	// because a precondition wasn't satisfied.
	//
	// Examples: the user must be on a VPN before using the program or have
	// a minimum version of some other software installed.
	RequirementNotMet Code = 82

	// Forbidden indicates that the program exited unsuccessfully
	// because the user isn't authorized to perform the requested action.
	Forbidden Code = 83

	// MovedPermanently indicates that the program exited unsuccessfully
	// because it has been migrated to a new location.
	MovedPermanently Code = 84
)

// Exit Codes 100-119 are reserved for software or system errors.
const (
	// InternalError indicates that the program exited unsuccessfully
	// because of a problem in its own code.
	//
	// Used instead of 1 when the problem is known to be with the program's
	// code or dependencies.
	InternalError Code = 100

	// Unavailable indicates that the program exited unsuccessfully
	// because a service it depends on was not available.
	//
	// Examples: A local daemon or remote service did not respond, a connection
	// was closed unexpectedly, an HTTP service responded with 503.
	Unavailable Code = 101
)

// IsUserError reports whether an exit code is a user error.
// It returns true if the code is in the range 80-99 and false if not.
func IsUserError(code Code) bool {
	return code >= 80 && code <= 99
}

// IsSoftwareError reports whether an exit code is a software error.
// It returns true if the code is in the range 100-119 and false if not.
func IsSoftwareError(code Code) bool {
	return code >= 100 && code <= 119
}

// IsSignal reports whether an exit code is derived from a signal.
// It returns true if the code is in the range 128-255 and false if not.
// It also returns true if the code is -1 because ProcessState.ExitCode()
// may return -1 if the process was terminated by a signal.
func IsSignal(code Code) bool {
	// -1 is not a valid exit code, but ProcessState.ExitCode()
	// may return -1 if the process was terminated by a signal.
	//
	// https://pkg.go.dev/os#ProcessState.ExitCode
	return code == -1 || code > 128 && code < 255
}

var (
	ErrNotOK             = Error{Code: NotOK}
	ErrUsageError        = Error{Code: UsageError}
	ErrUnknownSubcommand = Error{Code: UnknownSubcommand}
	ErrRequirementNotMet = Error{Code: RequirementNotMet}
	ErrForbidden         = Error{Code: Forbidden}
	ErrMovedPermanently  = Error{Code: MovedPermanently}
	ErrInternalError     = Error{Code: InternalError}
	ErrUnavailable       = Error{Code: Unavailable}
)

func FromError(err error) Code {
	var e Error
	if errors.As(err, &e) {
		return e.Code
	} else if err == nil {
		return OK
	} else {
		return NotOK
	}
}

func Wrap(err error, code Code) error {
	return Error{Code: code, Cause: err}
}

type Error struct {
	Code  Code
	Cause error
}

func (e Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	} else {
		return fmt.Sprintf("exit %d", e.Code)
	}
}

func (e Error) Unwrap() error {
	return e.Cause
}
