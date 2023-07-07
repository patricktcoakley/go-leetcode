package array

func majorityElement(nums []int) int {
	count := 0
	x := 0

	for _, num := range nums {
		if count == 0 {
			x = num
		}

		if num == x {
			count += 1
		} else {
			count -= 1
		}
	}

	return x
}
