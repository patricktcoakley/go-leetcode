package string

import "unicode"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		leftVal, rightVal := rune(s[left]), rune(s[right])

		for left < right && !(unicode.IsLetter(leftVal) || unicode.IsDigit(leftVal)) {
			left += 1
			leftVal = rune(s[left])
		}

		for left < right && !(unicode.IsLetter(rightVal) || unicode.IsDigit(rightVal)) {
			right -= 1
			rightVal = rune(s[right])
		}

		if unicode.ToUpper(leftVal) != unicode.ToUpper(rightVal) {
			return false
		}

		left += 1
		right -= 1
	}

	return true
}
