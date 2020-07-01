package main

//https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
//44ms 3.7MB
func findLength(A []int, B []int) int {

	//dp 最长公共子序列
	aLen := len(A)
	if aLen == 0 {
		return 0
	}

	bLen := len(B)
	if bLen == 0 {
		return 0
	}

	if aLen < bLen {
		aLen, bLen = bLen, aLen
		A, B = B, A
	}

	dp := make([][]int, 2) //以 i，j为结尾的连续子数组长度
	dp[0] = make([]int, bLen+1)
	dp[1] = make([]int, bLen+1)
	maxLen := 0

	for i := 1; i < aLen+1; i++ {
		preRow := (i + 1) & 1
		for j := 1; j < bLen+1; j++ {
			if A[i-1] == B[j-1] { // i，j 是从1开始，所以要-1 得到真实的索引
				dp[i&1][j] = dp[preRow][j-1] + 1
				if dp[i&1][j] > maxLen {
					maxLen = dp[i&1][j]
				}
				continue
			}

			// dp[i&1][j] = max(dp[i&1][j-1], dp[preRow][j])
			dp[i&1][j] = 0 //因为是连续子数组 所以这里要置为0

		}
	}

	return maxLen

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
