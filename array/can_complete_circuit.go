package array

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	tank := 0
	total := 0

	for i := 0; i < len(gas); i++ {
		curr := gas[i] - cost[i]
		tank += curr
		total += curr
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
