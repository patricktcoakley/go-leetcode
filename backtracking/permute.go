package backtracking

func permute(nums []int) [][]int {
	var perms [][]int
	n := len(nums)

	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			t := make([]int, n)
			copy(t, nums)
			perms = append(perms, t)
			return
		}

		for j := i; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			backtrack(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	backtrack(0)

	return perms
}
