package matrix

func countBattleships(board [][]byte) int {
	count := 0
	rows, cols := len(board), len(board[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'X' {
				count += 1
				battleshipHelper(board, rows, cols, row, col)
			}
		}
	}

	return count
}

func battleshipHelper(board [][]byte, rows, cols, row, col int) {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return
	}

	if board[row][col] == 'X' {
		board[row][col] = '.'
		battleshipHelper(board, rows, cols, row+1, col)
		battleshipHelper(board, rows, cols, row-1, col)
		battleshipHelper(board, rows, cols, row, col+1)
		battleshipHelper(board, rows, cols, row, col-1)
	}

}
