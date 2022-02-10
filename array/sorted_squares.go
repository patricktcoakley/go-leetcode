package array

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	curr := n - 1

	for l <= r {
		lVal, rVal := nums[l]*nums[l], nums[r]*nums[r]
		if lVal > rVal {
			res[curr] = lVal
			l += 1
		} else {
			res[curr] = rVal
			r -= 1
		}

		curr -= 1
	}
	return res
}
