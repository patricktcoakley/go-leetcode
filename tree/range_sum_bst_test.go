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
		want := tt.want
		got := rangeSumBST(tt.root, tt.low, tt.high)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}

}
