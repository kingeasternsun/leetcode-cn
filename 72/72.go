package main

import "fmt"

func main() {

	fmt.Println(minDistance("intention", "execution"))
}

func min(args ...int) int {
	v := args[0]
	for i := range args {
		if args[i] < v {
			v = args[i]
		}
	}
	return v
}

func minDistance(word1 string, word2 string) int {

	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}

	if len2 == 0 {
		return len1
	}

	dp := make([][]int, 2)

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}

	for i := 0; i < len2+1; i++ {
		dp[0][i] = i
	}

	rowID := 1
	preRowID := 0

	for i := 1; i < len1+1; i++ {

		// rowID := i % 2
		// preRowID := (rowID + 1) % 2

		//关键，如果优化存储，只使用2行的话，那么这里要记得更新
		dp[rowID][0] = i
		for j := 1; j < len2+1; j++ {
			// fmt.Println(i, j)

			if word1[i-1] == word2[j-1] {
				dp[rowID][j] = dp[preRowID][j-1]
			} else {
				dp[rowID][j] = min(dp[preRowID][j-1]+1, dp[rowID][j-1]+1, dp[preRowID][j]+1)

			}

		}

		tmp := rowID
		rowID = preRowID
		preRowID = tmp

	}

	fmt.Println(dp)

	return dp[len1%2][len2]
}

func minDistance_slow(word1 string, word2 string) int {

	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}

	if len2 == 0 {
		return len1
	}

	dp := make([][]int, 2)

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}

	for i := 0; i < len2+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			// fmt.Println(i, j)
			rowID := i % 2
			preRowID := (rowID + 1) % 2
			if word1[i-1] == word2[j-1] {
				dp[rowID][j] = dp[preRowID][j-1]
			} else {
				dp[rowID][j] = min(dp[preRowID][j-1]+1, dp[rowID][j-1]+1, dp[preRowID][j]+1)

			}

		}
	}

	fmt.Println(dp)

	return dp[len1%2][len2]
}
