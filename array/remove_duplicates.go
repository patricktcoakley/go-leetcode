package array

func removeDuplicates(nums []int) int {
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[start] {
			nums[start+1] = nums[i]
			start += 1
		}
	}

	return start + 1
}
