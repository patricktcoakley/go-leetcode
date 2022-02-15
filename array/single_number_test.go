package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{2, 2, 1}, want: 1},
		{nums: []int{4, 1, 2, 1, 2}, want: 4},
		{nums: []int{1}, want: 1},
	}

	for _, tt := range tests {
		got := singleNumber(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
