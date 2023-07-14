package string

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{words: []string{"flower", "flow", "flight"}, want: "fl"},
		{words: []string{"dog", "racecar", "car"}, want: ""},
	}
	for _, tt := range tests {
		got := longestCommonPrefix(tt.words)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
