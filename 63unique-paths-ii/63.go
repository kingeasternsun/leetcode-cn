package main

/*

63. 不同路径 II https://leetcode-cn.com/problems/unique-paths-ii/
0ms 2.6MB
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	if rows == 0 {
		return 0
	}

	cols := len(obstacleGrid[0])
	if cols == 0 {
		return 0
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	if obstacleGrid[rows-1][cols-1] == 1 {
		return 0
	}

	if rows == 1 {

		for i := 0; i < cols; i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}

	}

	if cols == 1 {
		for i := 0; i < rows; i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
	}

	//模拟 rows行 cols 列的dp
	dp := make([][]int, 2)
	dp[0] = make([]int, cols)
	dp[1] = make([]int, cols)

	//最后一行是 rows-1 行
	for j := cols - 1; j >= 0; j-- {

		if obstacleGrid[rows-1][j] == 1 { //有障碍物 肯定到不了
			dp[(rows-1)&1][j] = 0
		} else {
			// 没有障碍物 那就跟右边的一样，因为右边的可能因为障碍物无法到达，所以当前即使没障碍物，右边的有障碍物，也无法达到
			// 因为只能向右移动，所以 dp[(rows-1)&1][j] = dp[(rows-1)&1][j+1]
			if j == cols-1 {
				dp[(rows-1)&1][j] = 1
			} else {
				dp[(rows-1)&1][j] = dp[(rows-1)&1][j+1]
			}
		}

	}

	//从下往上 逐行计算
	row := rows - 2
	for row >= 0 {

		//初始化最右边的，只能向下移动
		if obstacleGrid[row][cols-1] == 1 { //有障碍物 肯定到不了
			dp[row&1][cols-1] = 0
		} else {
			//没有障碍物 那就跟下边的一样，因为下边的可能因为障碍物无法到达，所以当前即使没障碍物，下边的有障碍物，也无法达到
			// 因为只能向下移动，所以 dp[(rows-1)&1][j] = dp[(rows-1)&1][j-1]
			dp[row&1][cols-1] = dp[(row-1)&1][cols-1]

		}

		// 每一行 从右边到左边 计算
		for j := cols - 2; j >= 0; j-- {

			if obstacleGrid[row][j] == 1 { //有障碍物 肯定到不了
				dp[row&1][j] = 0
				continue
			}

			dp[row&1][j] = dp[(row+1)&1][j] + dp[row&1][j+1] // 往下 dp[(row-1)&1][j] 或 往右 dp[row][j+1] 的和
		}

		row--
	}

	return dp[0][0]

}
