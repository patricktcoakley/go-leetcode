package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head.Next

	for i := 0; i < n; i++ {
		if fast == nil {
			return head.Next
		}

		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
