package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{prices: []int{1, 2, 3, 4, 5}, want: 4},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}

	for _, tt := range tests {
		got := maxProfit2(tt.prices)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %d, got %d, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
