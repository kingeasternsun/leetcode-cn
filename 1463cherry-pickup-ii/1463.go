package main

// dp ，记录在 第i行，机器人分别在x y 位置的最大值
func newMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}

func maxPre(x, y, col int, dp [][]int) int {

	xs := make([]int, 0)
	xs = append(xs, x)
	if x > 0 {
		xs = append(xs, x-1)
	}
	if x < col-1 {
		xs = append(xs, x+1)
	}

	ys := make([]int, 0)
	ys = append(ys, y)
	if y > 0 {
		ys = append(ys, y-1)
	}
	if y < col-1 {
		ys = append(ys, y+1)
	}

	max := 0
	for _, x1 := range xs {
		for _, y1 := range ys {

			if dp[x1][y1] > max {
				max = dp[x1][y1]
			}
		}
	}

	return max
}

/*

 关键点
 1. 初始化的时候要加1，考虑到 初始状态的两个地方都可能没有樱桃
 2. 在迭代的时候，如果上一个最大值是0，表示不可达，那么当前也不可达
 76ms 6.6MB
*/
func cherryPickup1(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}

	dp := make([][][]int, 2)
	dp[0] = newMatrix(col)
	dp[1] = newMatrix(col)

	// 一开始两个机器人分别在 0 col-1 位置
	// 记录在 第0行，机器人分别在 0 col-1 位置的最大值
	dp[0][0][col-1] = grid[0][0] + grid[0][col-1] + 1 //考虑为0的情况 所以加1

	//
	for i := 1; i < row; i++ {

		for x := 0; x < col; x++ {
			for y := x; y < col; y++ {

				dp[i&1][x][y] = maxPre(x, y, col, dp[(i+1)&1])
				if dp[i&1][x][y] == 0 { //为0，表示这个状态不可能达到，所以就直接跳过
					continue
				}

				if x == y {
					dp[i&1][x][y] += grid[i][x]
				} else {
					dp[i&1][x][y] += (grid[i][x] + grid[i][y])
				}

			}
		}
	}

	max := 0
	for x := 0; x < col; x++ {
		for y := x; y < col; y++ {
			if dp[(row-1)&1][x][y] > max {
				max = dp[(row-1)&1][x][y]
			}

		}
	}

	return max - 1
}

/*
从上往下扩散 迭代 12ms 3.4MB
*/
func cherryPickup(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}

	dp := make([][][]int, 2)
	dp[0] = newMatrix(col)
	dp[1] = newMatrix(col)

	// 一开始两个机器人分别在 0 col-1 位置
	// 记录在 第0行，机器人分别在 0 col-1 位置的最大值
	dp[0][0][col-1] = grid[0][0] + grid[0][col-1] + 1 //考虑为0的情况 所以加1

	//
	for i := 0; i < row-1; i++ {

		for x := 0; x < col; x++ {
			for y := x; y < col; y++ {

				if dp[i&1][x][y] == 0 { //为0，表示这个状态不可能达到，所以就直接跳过
					continue
				}

				//计算下一层
				for x1 := maxF(0, x-1); x1 <= minF(x+1, col-1); x1++ {
					for y1 := maxF(0, y-1); y1 <= minF(y+1, col-1); y1++ {
						tmp := dp[i&1][x][y] //从当前i层的x，y 跳到i+1层的x1，y1

						if x1 == y1 {
							tmp += grid[i+1][x1]
						} else {
							tmp += (grid[i+1][x1] + grid[i+1][y1])
						}

						if tmp > dp[(i+1)&1][x1][y1] {
							dp[(i+1)&1][x1][y1] = tmp
						}

					}
				}

			}
		}
	}

	max := 0
	for x := 0; x < col; x++ {
		for y := x; y < col; y++ {
			if dp[(row-1)&1][x][y] > max {
				max = dp[(row-1)&1][x][y]
			}

		}
	}

	return max - 1
}

func maxF(x, y int) int {
	if x > y {
		return x
	}

	return y
}
func minF(x, y int) int {
	if x > y {
		return y
	}

	return x
}
