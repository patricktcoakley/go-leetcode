package string

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		want    bool
	}{
		{pattern: "abba", s: "dog cat cat dog", want: true},
		{pattern: "abba", s: "dog cat cat fish", want: false},
		{pattern: "aaaa", s: "dog cat cat dog", want: false},
	}

	for _, tt := range tests {
		got := wordPattern(tt.pattern, tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
