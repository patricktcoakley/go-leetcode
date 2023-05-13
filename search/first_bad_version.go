package search

// Can't test at this time
func isBadVersion(_ int) bool {
	return true
}

func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
