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

func solveSudoku(board [][]byte) {

	if len(board) != 9 || len(board[0]) != 9 {
		return
	}

	dp(board)

}

func dp(board [][]byte) bool {
	s := "123456789"
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := range s {

					if isValid(board, i, j, s[k]) {
						board[i][j] = s[k]
						if dp(board) {
							return true
						}
						board[i][j] = '.'
					}

				}

				return false
			}
		}
	}

	//  如果都填满了，并且都有效，肯定返回true
	return true
}
