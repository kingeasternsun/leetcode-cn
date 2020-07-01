package main

// 类似的https://leetcode-cn.com/problems/maximal-square/
/*
与求最大正方形一样，
1. 同样求矩阵中每个点作为正方形右下角的最大正方行的边长
2. 把这些边长求和就得到答案

为了求矩阵中每个点作为右下角的正方形的最大边长，
1. 首先计算每个点作为底的最长垂直线长度，对应代码中的 height，由于只跟上一行有关，所以优化为两行存储就可以
2. 每个点作为右下角的最大正方形边长跟 这个点左边点的最大正方形边长preLen 当前点作为底的垂直线长度有关curHeight
3. 详细处理看代码
*/
func countSquares(matrix [][]int) int {

	rows := len(matrix)
	if rows == 0 {
		return 0

	}

	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	//记录 以当前点的垂直线的长度
	height := make([][]int, 2)
	height[0] = make([]int, cols+1)
	height[1] = make([]int, cols+1)

	sum := 0

	for row := 0; row < rows; row++ {

		rowID := row & 1
		preRowID := (row + 1) & 1
		preLen := 0 //当前行中 前一个点的正方形最大边长

		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				height[rowID][col] = 0
				preLen = 0
				continue
			}

			curHeight := height[preRowID][col] + 1
			height[rowID][col] = curHeight

			//当前高度小于或等于左边点的正方形的边长，所以以该节点作为正方形右下角的正方形的最大边长就是 curHeight
			if curHeight <= preLen {
				sum += curHeight
				preLen = curHeight
				continue
			}

			//看前一个正方形上面是否有1
			j := col - preLen
			for ; j < col; j++ {
				if matrix[row-preLen][j] == 0 {
					break
				}
			}

			if j == col {
				sum = sum + preLen + 1
				preLen = preLen + 1
				// continue
			} else {
				sum += preLen
				preLen = preLen
			}

		}

	}

	return sum

}
