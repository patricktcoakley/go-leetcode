package graph

func cloneGraph(node *Node) *Node {
	m := make(map[*Node]*Node)

	return helper(node, m)
}

func helper(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := m[node]; ok {
		return n
	}

	clone := &Node{Val: node.Val, Neighbors: nil}

	m[node] = clone

	for _, neighbor := range node.Neighbors {
		m[clone].Neighbors = append(m[clone].Neighbors, helper(m[neighbor], m))
	}

	return m[node]
}

func cloneGraphIterative(node *Node) *Node {
	if node == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	m[node] = &Node{Val: node.Val, Neighbors: []*Node{}}
	q := []*Node{node}

	for len(q) > 0 {
		head := q[0]
		q = q[1:]

		for _, neighbor := range head.Neighbors {
			if _, ok := m[neighbor]; !ok {
				m[neighbor] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				q = append(q, neighbor)
			}

			m[head].Neighbors = append(m[head].Neighbors, m[neighbor])
		}
	}

	return m[node]
}
