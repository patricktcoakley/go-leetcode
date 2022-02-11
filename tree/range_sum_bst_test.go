package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRangeSumBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		low  int
		high int
		want int
	}{
		{root: createTree([]*int{intPtr(10), intPtr(5), intPtr(15), intPtr(3), intPtr(7), nil, intPtr(18)}), low: 7, high: 15, want: 32},
		{root: createTree([]*int{intPtr(10), intPtr(5), intPtr(15), intPtr(3), intPtr(7), intPtr(13), intPtr(18), intPtr(1), nil, intPtr(6)}), low: 6, high: 10, want: 23},
	}

	for _, tt := range tests {
		got := rangeSumBST(tt.root, tt.low, tt.high)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
