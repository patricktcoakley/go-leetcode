package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intPtr(i int) *int {
	return &i
}

func createTree(vals []*int) *TreeNode {
	n := len(vals)
	if n == 0 {
		return nil
	}

	i := 1
	root := &TreeNode{Val: *vals[0]}
	q := []*TreeNode{root}

	for len(q) > 0 && i < n {
		top := q[0]
		q = q[1:]

		if top != nil {
			if vals[i] != nil {
				top.Left = &TreeNode{Val: *vals[i]}
			} else {
				top.Left = nil
			}

			q = append(q, top.Left)
			i += 1
		}

		if i >= n {
			break
		}

		if vals[i] != nil {
			top.Right = &TreeNode{Val: *vals[i]}
		} else {
			top.Right = nil
		}

		q = append(q, top.Right)
		i += 1
	}

	return root
}
