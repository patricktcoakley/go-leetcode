package number

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 4, want: 4},
		{n: 25, want: 1389537},
	}

	for _, tt := range tests {
		got := tribonacci(tt.n)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
