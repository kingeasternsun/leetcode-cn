package main

import "strings"

/*
跟139类似
1. 139中dp记录的是前i个字符是否可以分割
2. 140中dp记录的是前i个字符可以构成的句子集合
*/
func wordBreak(s string, wordDict []string) []string {

	count := len(s)
	if count == 0 {
		return nil
	}

	if !checkwordBreak(s, wordDict) {
		return nil
	}

	dp := make([][]string, len(s)+1) //记录前i个字符，可以组成的列表
	for i := 0; i < count; i++ {

		for _, word := range wordDict {

			if strings.HasSuffix(s[:i+1], word) { //139里面这里就可以break，但是这里不能够

				preID := i + 1 - len(word)
				if preID == 0 { //直接到开始
					dp[i+1] = append(dp[i+1], word)
					continue
				}

				//前preID不可以拆分为句子
				if len(dp[preID]) == 0 {
					continue
				}

				//前preID可以拆分，拼接word后 加入到 dp[i+1]
				for _, preS := range dp[preID] {
					dp[i+1] = append(dp[i+1], preS+" "+word)
				}

			}
		}

	}
	return dp[count]

}

func checkwordBreak(s string, wordDict []string) bool {
	count := len(s)
	if count == 0 {
		return false
	}
	dps := make([]bool, count+1)
	dps[0] = true

	for i := 1; i <= count; i++ {

		for j := range wordDict {
			if strings.HasSuffix(s[:i], wordDict[j]) {
				if dps[i-len(wordDict[j])] {
					dps[i] = true
					break
				}
			}
		}
	}

	return dps[count]
}
