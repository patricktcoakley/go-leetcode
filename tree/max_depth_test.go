package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{root: createTree([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}), want: 3},
		{root: createTree([]*int{intPtr(1), nil, intPtr(2)}), want: 2},
	}

	for _, tt := range tests {
		got := maxDepth(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
