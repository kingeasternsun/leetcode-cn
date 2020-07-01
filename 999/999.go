package main

func main() {

}

/*
https://leetcode-cn.com/problems/available-captures-for-rook/  一开始被题目给吓到了
*/
func numRookCaptures(board [][]byte) int {
	rows := len(board)
	if rows == 0 {
		return 0
	}
	cols := len(board[0])
	if cols == 0 {
		return 0
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'R' {
				return move(i, j, 0, -1, board) + move(i, j, 0, 1, board) + move(i, j, -1, 0, board) + move(i, j, 1, 0, board)
			}
		}
	}

	return 0

}

func move(x, y int, dx, dy int, board [][]byte) int {
	for x >= 0 && x < 8 && y >= 0 && y < 8 {

		if board[x][y] == 'B' {
			return 0
		}
		if board[x][y] == 'p' {
			return 1
		}
		x += dx
		y += dy

	}

	return 0
}
