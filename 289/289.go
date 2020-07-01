package main

func main() {

}

/*

289. 生命游戏 https://leetcode-cn.com/problems/game-of-life/
0ms 2.1MB
*/
func gameOfLife(board [][]int) {

	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])
	if cols == 0 {
		return
	}

	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {
			judge(board, rows, cols, i, j)
		}

	}

	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {
			board[i][j] = board[i][j] >> 1
		}

	}

	return

}

func judge(board [][]int, rows, cols, i, j int) {
	liveCnt := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {

			if x == 0 && y == 0 {
				continue
			}

			if 0 <= i+x && i+x < rows && 0 <= j+y && j+y < cols {
				if (board[i+x][j+y] & 1) > 0 {
					liveCnt++
				}
			}

		}
	}

	//如果当前是活细胞
	if (board[i][j] & 1) > 0 {
		//对于活细胞，只有如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
		if liveCnt == 2 || liveCnt == 3 {
			board[i][j] = board[i][j] | 2
		}
	} else {
		//如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
		if liveCnt == 3 {
			board[i][j] = board[i][j] | 2
		}

	}
	return

}
