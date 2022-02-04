package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if res, ok := m[diff]; ok {
			return []int{i, res}
		}

		m[v] = i
	}

	return []int{0, 0}
}
