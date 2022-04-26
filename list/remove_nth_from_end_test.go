package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		list *ListNode
		n    int
		want *ListNode
	}{
		{list: createList([]int{1, 2, 3, 4, 5}), n: 2, want: createList([]int{1, 2, 3, 5})},
		{list: createList([]int{1}), n: 1, want: createList([]int{})},
		{list: createList([]int{1, 2}), n: 1, want: createList([]int{1})},
	}

	for _, tt := range tests {
		got := removeNthFromEnd(tt.list, tt.n)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
