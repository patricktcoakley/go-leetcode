package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: createTree([]*int{intPtr(1), nil, intPtr(2), intPtr(3)}),
			want: []int{1, 3, 2},
		},
		{
			root: createTree(nil),
			want: nil,
		},
		{
			root: createTree([]*int{intPtr(1)}),
			want: []int{1},
		},
	}

	for _, tt := range tests {
		got := inorderTraversal(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}

func TestInorderTraversalIterative(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: createTree([]*int{intPtr(1), nil, intPtr(2), intPtr(3)}),
			want: []int{1, 3, 2},
		},
		{
			root: nil,
			want: nil,
		},
		{
			root: createTree([]*int{intPtr(1)}),
			want: []int{1},
		},
	}

	for _, tt := range tests {
		got := inorderTraversalIterative(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
