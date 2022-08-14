package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		node *Node
		want *Node
	}{
		// TODO: add tests
		{node: nil, want: nil},
	}

	for _, tt := range tests {
		got := cloneGraph(tt.node)
		if !cmp.Equal(tt.want, got, cmpopts.SortSlices(func(x, y int) bool { return x < y })) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
