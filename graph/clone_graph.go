package graph

func cloneGraph(node *Node) *Node {
	m := map[*Node]*Node{}

	var cloneNode func(*Node, map[*Node]*Node) *Node
	cloneNode = func(node *Node, m map[*Node]*Node) *Node {
		if node == nil {
			return nil
		}

		if n, ok := m[node]; ok {
			return n
		}

		clone := &Node{Val: node.Val}
		m[node] = clone

		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, cloneNode(neighbor, m))
		}

		return clone
	}

	return cloneNode(node, m)
}
