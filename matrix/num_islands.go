package matrix

func numIslands(grid [][]byte) int {
	count := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count += 1
				islandHelper(grid, m, n, i, j)
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
