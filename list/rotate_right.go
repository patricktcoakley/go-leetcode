package list

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	var vals []int
	curr := head

	for curr != nil {
		vals = append(vals, curr.Val)
		curr = curr.Next
	}

	n := len(vals)
	k %= n

	reverse(vals, 0, n-1)
	reverse(vals, 0, k-1)
	reverse(vals, k, n-1)

	curr = head

	for _, val := range vals {
		curr.Val = val
		curr = curr.Next
	}

	return head
}

func reverse(vals []int, left int, right int) {
	for left < right {
		vals[left], vals[right] = vals[right], vals[left]
		left += 1
		right -= 1
	}
}
