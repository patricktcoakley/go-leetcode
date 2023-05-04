package array

func nextPermutation(nums []int) {
	n := len(nums)
	l := n - 2
	for l >= 0 && nums[l+1] <= nums[l] {
		l -= 1
	}

	if l >= 0 {
		r := n - 1
		for nums[r] <= nums[l] {
			r -= 1
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	l += 1
	r := n - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l += 1
		r -= 1
	}
}
