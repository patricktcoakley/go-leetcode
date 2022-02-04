package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		list *ListNode
		want bool
	}{
		{list: CreateCycledList([]int{3, 2, 0, -4}, 3, 1), want: true},
		{list: CreateCycledList([]int{1, 2}, 1, 0), want: true},
		{list: &ListNode{Val: 1}, want: false},
	}

	for _, tt := range tests {
		want := tt.want
		got := hasCycle(tt.list)
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v, diff %v", want, got, cmp.Diff(want, got))
		}
	}
}
