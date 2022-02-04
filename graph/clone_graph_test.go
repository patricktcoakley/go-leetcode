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
		{node: &Node{}, want: &Node{}},
	}

	for _, tt := range tests {
		want := tt.want
		got := cloneGraph(tt.node)
		if !cmp.Equal(want, got, cmpopts.SortSlices(func(x, y int) bool { return x < y })) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}
}
