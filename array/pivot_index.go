package array

func pivotIndex(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}

	l := 0
	for i, num := range nums {
		if l == s-l-num {
			return i
		}
		l += num
	}

	return -1
}
