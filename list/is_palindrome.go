package list

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverseList(slow)
	fast = head

	for slow.Next != nil {
		if slow.Val != fast.Val {
			return false
		}

		slow = slow.Next
		fast = fast.Next
	}

	return true
}
