package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindJudge(t *testing.T) {
	tests := []struct {
		n     int
		trust [][]int
		want  int
	}{
		{n: 2, trust: [][]int{{1, 2}}, want: 2},
		{n: 3, trust: [][]int{{1, 3}, {2, 3}}, want: 3},
		{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}, want: -1},
	}

	for _, tt := range tests {
		want := tt.want
		got := findJudge(tt.n, tt.trust)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}
}
