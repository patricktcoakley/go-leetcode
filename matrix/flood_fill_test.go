package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{
			image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			newColor: 2,
			want:     [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image:    [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:       0,
			sc:       0,
			newColor: 2,
			want:     [][]int{{2, 2, 2}, {2, 2, 2}},
		},
	}

	for _, tt := range tests {
		got := floodFill(tt.image, tt.sr, tt.sc, tt.newColor)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}

}
