package graph

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	m := make(map[int][][]int, len(times))

	for _, time := range times {
		node, neighbor, neighborTime := time[0], time[1], time[2]
		m[node] = append(m[node], []int{neighbor, neighborTime})
	}

	q := [][]int{{k, 0}}
	visited := make(map[int]int, len(times))

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		topNode, topTime := top[0], top[1]
		if _, ok := visited[topNode]; !ok {
			visited[topNode] = math.MaxInt
		}

		if visited[topNode] > topTime {
			visited[topNode] = topTime
			for _, neighbor := range m[topNode] {
				neighborNode, neighborTime := neighbor[0], neighbor[1]
				q = append(q, []int{neighborNode, topTime + neighborTime})
			}
		}
	}

	if len(visited) != n {
		return -1
	}

	var max = 0
	for _, v := range visited {
		if v > max {
			max = v
		}
	}

	return max
}
