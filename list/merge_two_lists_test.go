package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{list1: createList([]int{1, 2, 4}), list2: createList([]int{1, 3, 4}), want: createList([]int{1, 1, 2, 3, 4, 4})},
		{list1: nil, list2: nil, want: nil},
		{list1: nil, list2: &ListNode{}, want: &ListNode{}},
	}

	for _, tt := range tests {
		got := mergeTwoLists(tt.list1, tt.list2)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
