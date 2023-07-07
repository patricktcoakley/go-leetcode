package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{3, 2, 3}, want: 3},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}

	for _, tt := range tests {
		got := majorityElement(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
