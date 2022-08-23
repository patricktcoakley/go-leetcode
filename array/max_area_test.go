package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
	}

	for _, tt := range tests {
		got := maxArea(tt.height)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
