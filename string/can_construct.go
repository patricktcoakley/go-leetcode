package string

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, v := range magazine {
		m[v-'a'] += 1
	}

	for _, v := range ransomNote {
		if m[v-'a'] == 0 {
			return false
		}

		m[v-'a'] -= 1
	}

	return true
}
