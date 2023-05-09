package backtracking

func letterCombinations(digits string) (combinations []string) {
	n := len(digits)

	if n == 0 {
		return combinations
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func([]byte, int)
	backtrack = func(curr []byte, i int) {
		if i == n {
			combinations = append(combinations, string(curr))
			return
		}

		letters := m[digits[i]]

		for _, v := range letters {
			curr = append(curr, byte(v))
			backtrack(curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack([]byte{}, 0)
	return combinations
}
