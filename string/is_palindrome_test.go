package string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "A man, a plan, a canal: Panama", want: true},
		{s: "race a car", want: false},
		{s: " ", want: true},
	}
	for _, tt := range tests {
		got := isPalindrome(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
