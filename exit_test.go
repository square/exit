package exit

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestFromError(t *testing.T) {
	// Test cases for the FromError function
	tests := []struct {
		name string
		err  error
		want int
	}{
		{
			name: "nil",
			err:  nil,
			want: 0,
		},
		{
			name: "exit error",
			err:  ErrInternalError,
			want: 100,
		},
		{
			name: "error",
			err:  fmt.Errorf("wrapped error"),
			want: 1,
		},
		{
			name: "wrapped error",
			err:  Wrap(fmt.Errorf("wrapped error"), UsageError),
			want: 80,
		},
		{
			name: "error wrapped more than once",
			err:  Wrap(Wrap(fmt.Errorf("wrapped error"), UsageError), Unavailable),
			want: 101,
		},
		{
			name: "*exec.ExitError",
			err:  exec.Command("sh", "-c", "exit 3").Run(),
			want: 3,
		},
		{
			name: "wrapped *exec.ExitError",
			err:  Wrap(exec.Command("exit 3").Run(), NotOK),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromError(tt.err); got != tt.want {
				t.Errorf("FromError(%+v): got %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}
