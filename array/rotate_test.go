package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		rotate(tt.nums, tt.k)
		got := tt.nums
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
