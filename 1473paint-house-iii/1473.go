package main

import (
	"math"
)

/*

	前 i 个房子中 最后一个房子刷成 j 颜色， 构成 k 个社区，可能是由
	1. 前 i-1 个房子，最后一个房子颜色还是j颜色，构成	k	个社区
	2. 前 i-1 个房子，最后一个房子颜色不是j颜色，构成	k-1	个社区

	重点在于两个初始化
	1. 前 0   个房子，无论涂成生命颜色，都不可能组成 k>0 个社区 所以初始化为 maxint32
	2. 前 i>0 个房子，无论涂成生命颜色，都不可能组成   0 个社区 所以初始化为 maxint32


*/

func makeMaz(n, target int) [][]int {

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, target+1)
	}
	return a
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {

	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = makeMaz(n, target)
	}

	for j := 0; j < n; j++ {
		for k := 1; k < target+1; k++ {
			dp[0][j][k] = math.MaxInt32
		}
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][0] = math.MaxInt32
		}
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			for k := 1; k <= target; k++ {

				curCost := 0
				//如果房子已经图上颜色了
				if houses[i-1] > 0 {

					//和dp的颜色不一样
					if houses[i-1]-1 != j {
						dp[i][j][k] = math.MaxInt32
						continue
					} else {
						curCost = 0
					}

				} else {
					curCost = cost[i-1][j]
				}

				//前i-1房子的最后一个房子颜色也是j
				dp[i][j][k] = dp[i-1][j][k]
				// 2. 前 i-1 个房子，最后一个房子颜色不是j颜色，构成	k-1	个社区
				for c := 0; c < n; c++ {
					if (c != j) && (dp[i-1][c][k-1]) < dp[i][j][k] {
						dp[i][j][k] = dp[i-1][c][k-1]
					}
				}

				dp[i][j][k] = dp[i][j][k] + curCost
				// fmt.Println(i, j+1, k, dp[i][j][k])

			}

		}
	}

	minV := math.MaxInt32

	for j := 0; j < n; j++ {
		if dp[m][j][target] < minV {
			minV = dp[m][j][target]
		}
	}

	if minV == math.MaxInt32 {
		minV = -1
	}

	return minV
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
