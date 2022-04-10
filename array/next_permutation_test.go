package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3}, want: []int{1, 3, 2}},
		{nums: []int{3, 2, 1}, want: []int{1, 2, 3}},
		{nums: []int{1, 1, 5}, want: []int{1, 5, 1}},
	}
	for _, tt := range tests {
		got := tt.nums
		nextPermutation(got)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
