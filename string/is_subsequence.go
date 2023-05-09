package string

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	left, right := 0, 0

	for left < sLen && right < tLen {
		if s[left] == t[right] {
			left += 1
		}

		right += 1
	}

	return left == sLen
}
