package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{}
	curr := head
	for _, v := range arr {
		curr.Next = &ListNode{Val: v, Next: nil}
		curr = curr.Next
	}

	return head.Next
}

func CreateCycledList(arr []int, cycleFrom, cycleTo int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{}
	curr := head
	var to, from *ListNode
	for i, v := range arr {

		curr.Next = &ListNode{Val: v, Next: nil}
		curr = curr.Next

		if i == cycleTo {
			to = curr
		}

		if i == cycleFrom {
			from = curr
		}
	}

	curr = head.Next

	for curr != nil {
		if curr.Val == from.Val {
			from.Next = to
			break
		}
		curr = curr.Next
	}

	return head.Next
}
