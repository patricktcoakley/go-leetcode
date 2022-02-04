package tree

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high) + root.Val
}
