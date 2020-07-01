//  115. 不同的子序列 https://leetcode-cn.com/problems/distinct-subsequences/

package main

// 相比numDistinct2 优化空间
func numDistinct(s string, t string) int {
	tLen := len(t)
	if tLen == 0 {
		return 1
	}

	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	dp := make([][]int, 2) //dp【m】【n】表示 s的前m个字符串 包含 t的前n个字符串的个数
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, tLen+1)
		dp[i][0] = 1
	}

	// dp[0][0] = 1

	for i := 1; i < sLen+1; i++ {
		for j := 1; j < tLen+1; j++ {
			if s[i-1] != t[j-1] { //如果s的第i个字符和t的第j个字符不同，那么s删去第i个字符
				dp[i&1][j] = dp[(i+1)&1][j]
			} else {
				// 如果s的第i个字符和t的第j个字符相同，那么就要考虑 s删去第i个字符的情况，s和t都删去这个字符的情况
				dp[i&1][j] = dp[(i+1)&1][j-1] + dp[(i+1)&1][j]
			}
		}
	}

	return dp[sLen&1][tLen]

}

func numDistinct2(s string, t string) int {
	tLen := len(t)
	if tLen == 0 {
		return 1
	}

	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	dp := make([][]int, sLen+1) //dp【m】【n】表示 s的前m个字符串 包含 t的前n个字符串的个数
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]int, tLen+1)
		dp[i][0] = 1
	}

	// dp[0][0] = 1

	for i := 1; i < sLen+1; i++ {
		for j := 1; j < tLen+1; j++ {
			if s[i-1] != t[j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[sLen][tLen]

}

// timeout
func numDistinct1(s string, t string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	tLen := len(t)
	if tLen == 0 {
		return 1
	}

	subStr := make([]byte, 0)
	var result int
	dfs(s, sLen, 0, subStr, 0, t, tLen, &result)
	return result

}

// s的前id个 已经生成了 subStr time out
func dfs(s string, sLen int, id int, subStr []byte, subLen int, t string, tLen int, result *int) {
	if subLen == tLen || id == sLen {
		if string(subStr) == t {
			*result = *result + 1
		}
		return
	}

	// for i := id; i < sLen; i++ {
	// 第id个不放进去
	dfs(s, sLen, id+1, subStr, subLen, t, tLen, result)
	//第id个放进去
	if s[id] == t[subLen] {
		dfs(s, sLen, id+1, append(subStr, s[id]), subLen+1, t, tLen, result)
	}
	// }

	return

}
