package string

import "unicode"

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		sr, er := rune(s[start]), rune(s[end])

		for start < end && !(unicode.IsLetter(sr) || unicode.IsDigit(sr)) {
			start += 1
			sr = rune(s[start])
		}

		for start < end && !(unicode.IsLetter(er) || unicode.IsDigit(er)) {
			end -= 1
			er = rune(s[end])
		}

		if unicode.ToUpper(sr) != unicode.ToUpper(er) {
			return false
		}

		start += 1
		end -= 1
	}

	return true
}
