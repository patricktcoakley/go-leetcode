package matrix

func numIslands(grid [][]byte) int {
	count := 0
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count += 1
				islandHelper(grid, rows, cols, row, col)
			}
		}
	}

	return count
}

func islandHelper(grid [][]byte, m, n, row, col int) {
	if row < 0 || row >= m || col < 0 || col >= n {
		return
	}

	if grid[row][col] == '1' {
		grid[row][col] = '0'
		islandHelper(grid, m, n, row+1, col)
		islandHelper(grid, m, n, row-1, col)
		islandHelper(grid, m, n, row, col+1)
		islandHelper(grid, m, n, row, col-1)
	}
}

func numIslandsIterative(grid [][]byte) int {
	count := 0
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count += 1
				q := [][]int{{row, col}}
				for len(q) > 0 {
					top := q[len(q)-1]
					q = q[:len(q)-1]
					tRow, tCol := top[0], top[1]
					if tRow >= 0 && tRow < rows && tCol >= 0 && tCol < cols && grid[tRow][tCol] == '1' {
						grid[tRow][tCol] = '0'
						next := [][]int{{tRow + 1, tCol}, {tRow - 1, tCol}, {tRow, tCol + 1}, {tRow, tCol - 1}}
						q = append(q, next...)
					}
				}
			}
		}
	}

	return count
}
