package matrix

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		rowSet, colSet := make(map[byte]struct{}, 9), make(map[byte]struct{}, 9)
		for col := 0; col < 9; col++ {
			if !tryAdd(board[row][col], rowSet) || !tryAdd(board[col][row], colSet) {
				return false
			}
		}
	}

	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			boxSet := make(map[byte]struct{}, 9)
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if !tryAdd(board[i][j], boxSet) {
						return false
					}
				}
			}
		}
	}

	return true
}

func tryAdd(val byte, visited map[byte]struct{}) bool {
	if val == '.' {
		return true
	}

	if _, ok := visited[val]; ok {
		return false
	}

	visited[val] = struct{}{}
	return true
}
