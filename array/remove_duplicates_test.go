package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 2}, want: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}

	for _, tt := range tests {
		got := removeDuplicates(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
