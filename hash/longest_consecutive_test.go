package hash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
	}

	for _, tt := range tests {
		got := longestConsecutive(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
