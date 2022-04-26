package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas  []int
		cost []int
		want int
	}{
		{gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}, want: 3},
		{gas: []int{2, 3, 4}, cost: []int{3, 4, 3}, want: -1},
	}

	for _, tt := range tests {
		got := canCompleteCircuit(tt.gas, tt.cost)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
