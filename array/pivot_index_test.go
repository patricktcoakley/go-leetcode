package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{nums: []int{1, 2, 3}, want: -1},
		{nums: []int{2, 1, -1}, want: 0},
	}

	for _, tt := range tests {
		got := pivotIndex(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
