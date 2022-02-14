package hash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		jewels string
		stones string
		want   int
	}{
		{jewels: "aA", stones: "aAAbbbb", want: 3},
		{jewels: "z", stones: "ZZ", want: 0},
	}

	for _, tt := range tests {
		got := numJewelsInStones(tt.jewels, tt.stones)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
