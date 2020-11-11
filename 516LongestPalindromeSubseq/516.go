package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab") == 4)
	fmt.Println(longestPalindromeSubseq("cbbd") == 2)
	// fmt.Println(longestPalindromeSubseq("cbbd") == 2)
}

//最长公共子序列的问题
func longestPalindromeSubseq(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	num := len(s)

	dp := make([][]int, num+1)

	for i := 0; i <= num; i++ {
		dp[i] = make([]int, num+1)
	}

	for i := num; i > 0; i-- {
		for j := i; j <= num; j++ {

			if i == j {
				dp[i][j] = 1
				continue
			}

			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}

		}
	}

	return dp[1][num]

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
