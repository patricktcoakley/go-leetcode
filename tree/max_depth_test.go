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
		want := tt.want
		got := maxDepth(tt.root)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}

}
