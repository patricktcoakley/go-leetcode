package string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefangIPaddr(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "1.1.1.1", want: "1[.]1[.]1[.]1"},
		{s: "255.100.50.0", want: "255[.]100[.]50[.]0"},
	}
	for _, tt := range tests {
		got := defangIPaddr(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
