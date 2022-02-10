package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidPath(t *testing.T) {
	tests := []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{n: 3, edges: [][]int{{0, 1}, {1, 2}, {2, 0}}, source: 0, destination: 2, want: true},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, source: 0, destination: 5, want: false},
	}

	for _, tt := range tests {
		want := tt.want
		got := validPath(tt.n, tt.edges, tt.source, tt.destination)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}
}
