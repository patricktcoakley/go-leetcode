package string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abcabcbb", want: 3},
		{s: "bbbbb", want: 1},
		{s: "pwwkew", want: 3},
	}
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
