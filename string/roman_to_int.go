package string

func romanToInt(s string) int {
	total := numeralToInt(s[len(s)-1])
	for curr := 1; curr < len(s); curr++ {
		prevVal, currVal := numeralToInt(s[curr-1]), numeralToInt(s[curr])
		if prevVal >= currVal {
			total += prevVal
		} else {
			total -= prevVal
		}
	}

	return total
}

func numeralToInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return -1
}
