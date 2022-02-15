package hash

func longestConsecutive(nums []int) int {
	longest := 0
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		m[num] = true
	}

	for num := range m {
		if !m[num-1] {
			curr := num
			streak := 1

			for m[curr+1] {
				curr += 1
				streak += 1
			}

			if longest < streak {
				longest = streak
			}
		}
	}

	return longest
}
