package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		{head: createList([]int{1, 2, 2, 1}), want: true},
		{head: createList([]int{1, 2}), want: false},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.head)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
