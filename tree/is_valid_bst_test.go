package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{root: createTree([]*int{intPtr(2), intPtr(1), intPtr(3)}), want: true},
		{root: createTree([]*int{intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6)}), want: false},
	}

	for _, tt := range tests {
		got := isValidBST(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
