package list

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head  *ListNode
		left  int
		right int
		want  *ListNode
	}{
		{head: createList([]int{1, 2, 3, 4, 5}), left: 2, right: 4, want: createList([]int{1, 4, 3, 2, 5})},
		{head: &ListNode{Val: 5}, left: 1, right: 1, want: &ListNode{Val: 5}},
	}

	for _, tt := range tests {
		got := reverseBetween(tt.head, tt.left, tt.right)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
