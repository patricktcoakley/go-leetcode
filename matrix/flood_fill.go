package matrix

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}

	m := len(image)
	n := len(image[0])
	var fill func(int, int)
	fill = func(row, col int) {
		if row >= 0 && row < m && col >= 0 && col < n && image[row][col] == oldColor {
			image[row][col] = newColor
			fill(row+1, col)
			fill(row-1, col)
			fill(row, col+1)
			fill(row, col-1)
		}
	}

	fill(sr, sc)
	return image
}
