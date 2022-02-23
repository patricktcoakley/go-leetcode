package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			root: createTree([]*int{intPtr(4), intPtr(2), intPtr(7), intPtr(1), intPtr(3), intPtr(6), intPtr(9)}),
			want: createTree([]*int{intPtr(4), intPtr(7), intPtr(2), intPtr(9), intPtr(6), intPtr(3), intPtr(1)}),
		},
		{
			root: createTree([]*int{intPtr(2), intPtr(1), intPtr(3)}),
			want: createTree([]*int{intPtr(2), intPtr(3), intPtr(1)}),
		},
		{
			root: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		got := invertTree(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
