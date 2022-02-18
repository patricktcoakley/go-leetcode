package dp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 2, want: 2},
		{n: 3, want: 3},
	}

	for _, tt := range tests {
		got := climbStairs(tt.n)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
