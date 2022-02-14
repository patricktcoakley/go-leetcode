package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		list *ListNode
		want *ListNode
	}{
		{list: createList([]int{1, 2, 3, 4, 5}), want: createList([]int{5, 4, 3, 2, 1})},
		{list: &ListNode{}, want: &ListNode{}},
	}

	for _, tt := range tests {
		got := reverseList(tt.list)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
