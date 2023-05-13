package search

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
	}
	for _, tt := range tests {
		got := search(tt.nums, tt.target)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
