package array

func maxSubArray(nums []int) int {
	max := nums[0]
	curr := nums[0]

	for _, v := range nums[1:] {
		if v > curr+v {
			curr = v
		} else {
			curr += v
		}

		if curr > max {
			max = curr
		}
	}

	return max
}
