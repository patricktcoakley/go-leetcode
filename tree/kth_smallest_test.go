package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want int
	}{
		{root: createTree([]*int{intPtr(3), intPtr(1), intPtr(4), nil, intPtr(2)}), k: 1, want: 1},
		{root: createTree([]*int{intPtr(5), intPtr(3), intPtr(6), intPtr(2), intPtr(4), nil, nil, intPtr(1)}), k: 3, want: 3},
	}
	for _, tt := range tests {
		got := kthSmallest(tt.root, tt.k)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
