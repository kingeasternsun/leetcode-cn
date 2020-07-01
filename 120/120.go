// https://leetcode-cn.com/problems/triangle/submissions/ 120. 三角形最小路径和

package main

func min(i, j int) int {
	if j < i {
		return j
	}

	return i
}

func minimumTotal(triangle [][]int) int {

	if triangle == nil {
		return 0
	}

	count := len(triangle)
	if count == 1 {
		return triangle[0][0]
	}

	dp := make([]int, count)
	copy(dp, triangle[count-1])

	for i := count - 1; i > 0; i-- { //第i层,每一层包含i个

		for j := 0; j < i; j++ {
			// a, b, c := getID(i+1, j), getID(i+1, j+1), triangle[i-1][j]
			// fmt.Println(a, b, c)
			dp[j] = min(dp[j], dp[j+1]) + triangle[i-1][j] //第i层 第j下标的 跟 第i+1层，第j，j+1的相关
		}

	}

	return dp[0]

}
func minimumTotal2(triangle [][]int) int {

	if triangle == nil {
		return 0
	}

	count := len(triangle)
	if count == 1 {
		return triangle[0][0]
	}

	dp := make([]int, count*(1+count)/2)

	for j := 0; j < count; j++ {
		dp[getID(count, j)] = triangle[count-1][j] //第count层，第 j 个先赋值
	}

	for i := count - 1; i > 0; i-- { //第i层,每一层包含i个

		for j := 0; j < i; j++ {
			// a, b, c := getID(i+1, j), getID(i+1, j+1), triangle[i-1][j]
			// fmt.Println(a, b, c)
			dp[getID(i, j)] = min(dp[getID(i+1, j)], dp[getID(i+1, j+1)]) + triangle[i-1][j] //第i层 第j下标的 跟 第i+1层，第j，j+1的相关
		}

	}

	return dp[0]

}

//第 i 层，索引是j
func getID(i, j int) int {
	return (i)*(i-1)/2 + j
}
