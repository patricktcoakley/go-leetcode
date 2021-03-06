package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}

	for _, tt := range tests {
		got := tt.matrix
		rotate(got)

		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
