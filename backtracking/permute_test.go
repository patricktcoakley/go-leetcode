package backtracking

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
		{nums: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}}},
		{nums: []int{1}, want: [][]int{{1}}},
	}

	for _, tt := range tests {
		got := permute(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
