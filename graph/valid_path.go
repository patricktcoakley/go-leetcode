package graph

func validPath(_ int, edges [][]int, source int, destination int) bool {
	nodes := make(map[int][]int)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		nodes[node1] = append(nodes[node1], node2)
		nodes[node2] = append(nodes[node2], node1)
	}

	q := []int{source}
	seen := make(map[int]bool)

	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		if front == destination {
			return true
		}

		seen[front] = true
		for _, neighbor := range nodes[front] {
			if _, ok := seen[neighbor]; !ok {
				q = append(q, neighbor)
			}
		}
	}

	return false
}
