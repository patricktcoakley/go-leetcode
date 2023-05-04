package stack

func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	var stack []rune

	for _, v := range s {
		if curr, ok := m[v]; ok {

			n := len(stack)
			if n == 0 || curr != stack[n-1] {
				return false
			}

			stack = stack[:n-1]

		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}
