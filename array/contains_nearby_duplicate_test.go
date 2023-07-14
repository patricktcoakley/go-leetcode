package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{nums: []int{1, 2, 3, 1}, k: 3, want: true},
		{nums: []int{1, 0, 1, 1}, k: 1, want: true},
		{nums: []int{1, 2, 3, 1, 2, 3}, want: false},
	}

	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %t, got %t, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
