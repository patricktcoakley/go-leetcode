package backtracking

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", want: []string{}},
		{digits: "2", want: []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		got := letterCombinations(tt.digits)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %s, got %s, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
