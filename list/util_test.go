package list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateList(t *testing.T) {
	got := CreateList([]int{1, 2, 3})
	want := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	if !cmp.Equal(want, got, cmpopts.SortSlices(func(x, y int) bool { return x < y })) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}
}

func TestCreateCycledList(t *testing.T) {
	got := CreateCycledList([]int{1, 2, 3}, 1, 0)
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2, Next: head}
	want := head

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}

	got = CreateCycledList([]int{1}, 0, 0)
	head = &ListNode{Val: 1}
	head.Next = head
	want = head

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v, diff %s", want, got, cmp.Diff(want, got))
	}
}
