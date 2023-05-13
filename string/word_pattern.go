package string

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	m := make(map[string]rune, len(pattern))
	used := make([]bool, 26)

	for i, letter := range pattern {
		word := words[i]
		if val, ok := m[word]; ok {
			if val != letter {
				return false
			}
		} else {
			if used[letter-'a'] {
				return false
			}
			m[word] = letter
			used[letter-'a'] = true
		}
	}

	return true
}
