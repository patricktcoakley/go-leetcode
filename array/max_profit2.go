package array

func maxProfit2(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if curr, prev := prices[i], prices[i-1]; curr > prev {
			profit += curr - prev
		}
	}

	return profit
}
