package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		head.Left, head.Right = head.Right, head.Left
		if head.Left != nil {
			q = append(q, head.Left)
		}

		if head.Right != nil {
			q = append(q, head.Right)
		}
	}
	return root
}
