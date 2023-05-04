package dp

func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 1

	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}

	return c
}
