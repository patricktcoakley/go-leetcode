package string

import "strings"

func longestCommonPrefix(words []string) string {
	var longest strings.Builder

	shortest := words[0]
	for _, str := range words[1:] {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	for i, v := range shortest {
		for _, word := range words {
			if rune(word[i]) != v {
				return longest.String()
			}
		}

		longest.WriteRune(v)
	}

	return longest.String()
}
