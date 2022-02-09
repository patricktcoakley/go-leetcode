package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
		// TODO: Add tests
	}{}

	for _, tt := range tests {
		got := maxDepth(tt.root)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}