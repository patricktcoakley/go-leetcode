package hash

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]bool)
	total := 0

	for _, c := range jewels {
		m[c] = true
	}

	for _, c := range stones {
		if _, ok := m[c]; ok {
			total += 1
		}
	}

	return total
}
