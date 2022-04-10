package string

func isAnagram(s string, t string) bool {
	size := len(s)
	if size != len(t) {
		return false
	}

	m := map[byte]int{}
	for i := 0; i < size; i++ {
		m[s[i]] += 1
		m[t[i]] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
