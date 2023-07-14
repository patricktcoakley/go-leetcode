package string

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "Hello World", want: 5},
		{s: "   fly me   to   the moon  ", want: 4},
		{s: "luffy is still joyboy", want: 6},
	}
	for _, tt := range tests {
		got := lengthOfLastWord(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
