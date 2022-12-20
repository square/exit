package exit

import (
	"syscall"
	"testing"
)

func TestFromSignal(t *testing.T) {
	scenarios := []struct {
		signal   syscall.Signal
		expected Code
	}{
		{syscall.SIGINT, Code(130)},
	}

	for _, s := range scenarios {
		if got := FromSignal(s.signal); got != s.expected {
			t.Errorf("Expected FromSignal(%d) to be %d; got %d", s.signal, s.expected, got)
		}
	}
}
