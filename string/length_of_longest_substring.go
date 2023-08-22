package string

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	longest := 0
	start := 0

	for i, v := range s {
		if val, ok := m[v]; ok && val > start {
			start = m[v]
		}

		if i-start+1 > longest {
			longest = i - start + 1
		}

		m[v] = i + 1
	}

	return longest
}
