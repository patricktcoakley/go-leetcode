package tree

func inorderTraversal(root *TreeNode) []int {
	vals := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		vals = append(vals, node.Val)
		traverse(node.Right)
	}

	traverse(root)

	return vals
}

func inorderTraversalIterative(root *TreeNode) []int {
	vals := []int{}
	if root == nil {
		return vals
	}

	var s []*TreeNode
	curr := root
	for curr != nil || len(s) > 0 {
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}
		curr = s[len(s)-1]
		s = s[:len(s)-1]
		vals = append(vals, curr.Val)
		curr = curr.Right
	}

	return vals
}
