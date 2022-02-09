package stack

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(}", want: false},
	}

	for _, tt := range tests {
		got := isValid(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %s", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
