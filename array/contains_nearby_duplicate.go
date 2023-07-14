package array

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if val, ok := m[v]; ok && abs(val-i) <= k {
			return true
		}

		m[v] = i
	}

	return false
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
