package graph

func findJudge(n int, trust [][]int) int {
	trustLevel := make([]int, n+1)

	for _, v := range trust {
		trustLevel[v[0]] -= 1
		trustLevel[v[1]] += 1
	}

	for i := 1; i < n+1; i++ {
		if trustLevel[i] == n-1 {
			return i
		}
	}

	return -1
}
