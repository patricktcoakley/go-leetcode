package array

func maxProfit(prices []int) int {
	profit := 0
	lowest := prices[0]
	for _, v := range prices[1:] {
		if v < lowest {
			lowest = v
		}

		if v-lowest > profit {
			profit = v - lowest
		}

	}

	return profit
}
