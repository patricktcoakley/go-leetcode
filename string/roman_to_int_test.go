package string

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "III", want: 3},
		{s: "LVIII", want: 58},
		{s: "MCMXCIV", want: 1994},
	}

	for _, tt := range tests {
		got := romanToInt(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
