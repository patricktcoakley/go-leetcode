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
		// TODO: Add tests
	}{}

	for _, tt := range tests {
		got := rangeSumBST(tt.root, tt.low, tt.high)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
