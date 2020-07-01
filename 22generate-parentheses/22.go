package main

import "strings"

func generateParenthesis(n int) []string {

	curResult := make([]byte, 0)
	result := make([]string, 0)
	dfs(0, 0, n, curResult, &result)
	return result
}

// (lnum, rnum 分别表示当前左括号 右括号的个数，只要保证左括号数目大于等于右括号就可以
func dfs(lnum, rnum, n int, curResult []byte, result *[]string) {

	if lnum == n && rnum == n {
		// curResult = append(curResult[:0:0], curResult...)
		*result = append(*result, string(curResult))
		return
	}

	//左右括号数目相同，只能放左括号
	if lnum == rnum {
		dfs(lnum+1, rnum, n, append(curResult, '('), result)
		return
	}

	//左括号数目够了，只能放右括号
	if lnum == n {

		suf := strings.Repeat(")", lnum-rnum)

		*result = append(*result, string(curResult)+suf)

		// dfs(lnum+1, rnum, n, append(curResult, '('), result)
		return
	}

	//到这里只会是左括号大于右括号
	dfs(lnum+1, rnum, n, append(curResult, '('), result)
	dfs(lnum, rnum+1, n, append(curResult, ')'), result)
	return
}
