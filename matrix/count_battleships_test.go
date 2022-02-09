package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountBattleships(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{matrix: [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}, want: 2},
		{matrix: [][]byte{{'.'}}, want: 0},
	}

	for _, tt := range tests {
		got := countBattleships(tt.matrix)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
