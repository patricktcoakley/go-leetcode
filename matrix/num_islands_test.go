package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}}, want: 1},
		{grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}}, want: 3},
	}

	for _, tt := range tests {
		got := numIslands(tt.grid)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
