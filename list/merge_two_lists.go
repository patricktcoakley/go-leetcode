package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := &ListNode{}
	curr := list3

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list2 != nil {
		curr.Next = list2
	}

	if list1 != nil {
		curr.Next = list1
	}

	return list3.Next
}
