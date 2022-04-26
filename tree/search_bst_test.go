package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSearchBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			root: createTree([]*int{intPtr(4), intPtr(2), intPtr(7), intPtr(1), intPtr(3)}),
			val:  2,
			want: createTree([]*int{intPtr(2), intPtr(1), intPtr(3)}),
		},
		{
			root: createTree([]*int{intPtr(4), intPtr(2), intPtr(7), intPtr(1), intPtr(3)}),
			val:  5,
			want: nil,
		},
	}

	for _, tt := range tests {
		got := searchBST(tt.root, tt.val)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
