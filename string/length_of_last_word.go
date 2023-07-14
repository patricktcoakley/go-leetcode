package string

import "unicode"

func lengthOfLastWord(s string) int {
	var length int

	if len(s) < 1 {
		return length
	}

	for curr := len(s) - 1; curr >= 0; curr-- {
		if !unicode.IsLetter(rune(s[curr])) && length > 0 {
			break
		} else if unicode.IsLetter(rune(s[curr])) {
			length++
		}
	}

	return length
}
