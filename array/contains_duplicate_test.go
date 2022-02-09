package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 2, 3, 1}, want: true},
		{nums: []int{1, 2, 3, 4}, want: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for _, tt := range tests {
		got := containsDuplicate(tt.nums)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %t, got %t, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
