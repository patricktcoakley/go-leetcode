package array

func permute(nums []int) [][]int {
	perms := [][]int{}

	var backtrack func([]int)
	backtrack = func(curr []int) {
		if len(curr) == len(nums) {
			t := make([]int, len(curr))
			copy(t, curr)
			perms = append(perms, t)
			return
		}

		for _, num := range nums {
			if !contains(curr, num) {
				curr = append(curr, num)
				backtrack(curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack([]int{})

	return perms
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}

	return false
}
