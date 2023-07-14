package list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	var s []int
	length := 1
	curr := head

	for curr != nil {
		if length >= left && length <= right {
			s = append(s, curr.Val)
		}

		length++
		curr = curr.Next
	}

	curr = head

	for i := 1; i < length; i++ {
		if i >= left && i <= right {
			curr.Val = s[len(s)-1]
			s = s[:len(s)-1]
		}

		curr = curr.Next
	}

	return head
}
