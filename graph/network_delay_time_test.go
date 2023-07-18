package graph

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		times [][]int
		n     int
		k     int
		want  int
	}{
		{times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, n: 4, k: 2, want: 2},
		{times: [][]int{{1, 2, 1}}, n: 2, k: 1, want: 1},
		{times: [][]int{{1, 2, 1}}, n: 2, k: 2, want: -1},
	}

	for _, tt := range tests {
		want := tt.want
		got := networkDelayTime(tt.times, tt.n, tt.k)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
		}
	}
}
