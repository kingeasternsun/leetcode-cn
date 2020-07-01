package main

func main() {

}

/*
62. 不同路径 https://leetcode-cn.com/problems/unique-paths/
*/
// 0ms 1.9mB
func uniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	if m < n {
		m, n = n, m //保证 n为最小
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)

	//模拟 m行 n 列的dp，最后一行是 m-1 行
	for j := 0; j < n-1; j++ {
		dp[(m-1)&1][j] = 1
	}

	//从下往上 逐行计算
	row := m - 2
	for row >= 0 {

		dp[row&1][n-1] = 1
		// 每一行 从右边到左边 计算
		for j := n - 2; j >= 0; j-- {
			dp[row&1][j] = dp[(row+1)&1][j] + dp[row&1][j+1] // 往下 dp[(row-1)&1][j] 或 往右 dp[row][j+1] 的和
		}

		row--
	}

	return dp[0][0]
}
