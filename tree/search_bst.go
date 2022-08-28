package tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}

	return root
}

func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
