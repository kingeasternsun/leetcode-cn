package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "b"
	a = a[1:]
	fmt.Println(a == "")
}

/*
139. 单词拆分 https://leetcode-cn.com/problems/word-break/
*/

func wordBreakdfs(s string, wordDict []string) bool {
	// sort.Strings(wordDict)
	return dfs(s, wordDict)
}

//time out
func dfs(s string, wordDict []string) bool {

	if s == "" {
		return true
	}

	for i := range wordDict {
		if strings.HasPrefix(s, wordDict[i]) {
			if dfs(s[len(wordDict[i]):], wordDict) {
				return true
			}
		}
	}

	return false

}

// 0ms 2.1MB
func wordBreak(s string, wordDict []string) bool {
	// sort.Strings(wordDict)
	return dp(s, wordDict)
}
func dp(s string, wordDict []string) bool {
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
