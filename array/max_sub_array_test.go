package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{nums: []int{1}, want: 1},
		{nums: []int{5, 4, -1, 7, 8}, want: 23},
	}

	for _, tt := range tests {
		got := maxSubArray(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
