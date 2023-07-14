package list

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{head: createList([]int{1, 2, 3, 4, 5}), k: 2, want: createList([]int{4, 5, 1, 2, 3})},
		{head: createList([]int{0, 1, 2}), k: 4, want: createList([]int{2, 0, 1})},
	}

	for _, tt := range tests {
		got := rotateRight(tt.head, tt.k)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want %v, got %v, diff %v", tt.want, got, cmp.Diff(tt.want, got))
		}
	}
}
