package graph

func cloneGraph(node *Node) *Node {
	var m = make(map[*Node]*Node)
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
		clone.Neighbors = append(clone.Neighbors, helper(neighbor, m))
	}

	return clone

}
