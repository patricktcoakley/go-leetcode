package string

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{ransomNote: "a", magazine: "b", want: false},
		{ransomNote: "aa", magazine: "ab", want: false},
		{ransomNote: "aa", magazine: "aab", want: true},
	}
	for _, tt := range tests {
		got := canConstruct(tt.ransomNote, tt.magazine)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
