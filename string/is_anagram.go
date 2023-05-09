package string

func isAnagram(s string, t string) bool {
	n := len(s)
	if n != len(t) {
		return false
	}

	m := make(map[byte]int, n)
	for i := 0; i < n; i++ {
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
