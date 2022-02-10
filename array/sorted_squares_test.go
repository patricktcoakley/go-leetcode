package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nums: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
		{nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		got := sortedSquares(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
