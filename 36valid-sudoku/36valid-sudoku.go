package main

func isValid(board [][]byte, i, j int, c byte) bool {

	// 所在行
	for k := 0; k < 9; k++ {
		if board[i][k] == c {
			return false
		}

		if board[k][j] == c {
			return false
		}

	}

	// 所在3*3方块
	for row := i / 3 * 3; row < (i/3*3 + 3); row++ {
		for col := j / 3 * 3; col < (j/3*3 + 3); col++ {
			if board[row][col] == c {
				return false
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				c := board[i][j]
				board[i][j] = '.'
				if !isValid(board, i, j, c) {
					return false
				}
				board[i][j] = c
			}
		}
	}
	return true

}
