package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}

	for _, tt := range tests {
		want := tt.want
		got := maxProfit(tt.prices)
		if !cmp.Equal(want, got) {
			t.Errorf("want %d, got %d, diff %s", want, got, cmp.Diff(want, got))
		}
	}
}
