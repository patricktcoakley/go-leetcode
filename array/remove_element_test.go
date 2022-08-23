package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, want: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5},
	}

	for _, tt := range tests {
		got := removeElement(tt.nums, tt.val)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
