package string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abc", t: "ahbgdc", want: true},
		{s: "axc", t: "ahbgdc", want: false},
	}
	for _, tt := range tests {
		got := isSubsequence(tt.s, tt.t)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
