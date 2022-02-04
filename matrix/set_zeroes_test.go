package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, want: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for _, tt := range tests {
		want := tt.want
		got := tt.matrix
		setZeroes(got)

		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %v", want, got, cmp.Diff(want, got))
		}
	}
}
