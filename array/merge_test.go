package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
	}
	for _, tt := range tests {
		got := tt.nums1
		merge(got, tt.m, tt.nums2, tt.n)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
