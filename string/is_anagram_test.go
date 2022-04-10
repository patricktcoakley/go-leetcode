package string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{s: "anagram", t: "nagaram", want: true},
		{s: "rat", t: "car", want: false},
	}
	for _, tt := range tests {
		got := isAnagram(tt.s, tt.t)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
