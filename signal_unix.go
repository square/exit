//go:build linux || darwin
// +build linux darwin

package exit

import "syscall"

// FromSignal returns the exit code that corresponds to when a program
// exits in response to a signal.
func FromSignal(signal syscall.Signal) Code {
	// According to https://tldp.org/LDP/abs/html/exitcodes.html, it's standard
	// for a unix process to exit with 128 + n where n is a fatal signal.
	return Code(128 + int(signal))
}
