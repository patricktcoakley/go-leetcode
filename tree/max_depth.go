package tree

func maxDepth(root *TreeNode) int {
	return maxDepthTail(root, 0)
}

func maxDepthTail(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	left := maxDepthTail(root.Left, depth+1)
	right := maxDepthTail(root.Right, depth+1)

	if left > right {
		return left
	}

	return right
}
