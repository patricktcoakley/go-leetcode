package tree

func kthSmallest(root *TreeNode, k int) int {
	s := []*TreeNode{}
	for {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}

		root = s[len(s)-1]
		s = s[:len(s)-1]
		k -= 1
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

	return -1
}
