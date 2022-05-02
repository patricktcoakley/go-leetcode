package backtracking

func letterCombinations(digits string) []string {
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

	res := []string{}

	if len(digits) == 0 {
		return res
	}

	var backtrack func(int, string)
	backtrack = func(i int, curr string) {
		if i == len(digits) {
			res = append(res, curr)
			return
		}

		letters := m[digits[i]]
		for c := range letters {
			curr += string(letters[c])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, "")

	return res
}
