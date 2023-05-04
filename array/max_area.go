package array

func maxArea(height []int) int {
	maxHeight := 0
	left, right := 0, len(height)-1

	for left < right {
		maxHeight = max(maxHeight, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return maxHeight
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
