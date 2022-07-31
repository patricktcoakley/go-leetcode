package string

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	sCurr := 0
	tCurr := 0

	for tCurr < len(t) {
		if s[sCurr] == t[tCurr] {
			sCurr += 1
		}
		if sCurr == len(s) {
			return true
		}
		tCurr += 1
	}
	return false
}
