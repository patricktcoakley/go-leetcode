package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateBST(t *testing.T) {
	tests := []struct {
		vals []*int
		want *TreeNode
	}{
		{vals: []*int{}, want: nil},
		{vals: []*int{intPtr(1)}, want: &TreeNode{Val: 1}},
		{vals: []*int{intPtr(2), intPtr(1), intPtr(3)}, want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{vals: []*int{intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6)}, want: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}},
		{vals: []*int{intPtr(5), intPtr(3), intPtr(6), intPtr(2), intPtr(4), nil, nil, intPtr(1)}, want: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}},
		{vals: []*int{intPtr(1), intPtr(2), intPtr(2), nil, intPtr(3), nil, intPtr(3)}, want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}},
	}

	for _, tt := range tests {
		got := createTree(tt.vals)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
