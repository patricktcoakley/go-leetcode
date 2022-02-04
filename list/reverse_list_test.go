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
		{list: CreateList([]int{1, 2, 3, 4, 5}), want: CreateList([]int{5, 4, 3, 2, 1})},
		{list: &ListNode{}, want: &ListNode{}},
	}

	for _, tt := range tests {
		want := tt.want
		got := reverseList(tt.list)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %v", want, got, cmp.Diff(want, got))
		}
	}
}
