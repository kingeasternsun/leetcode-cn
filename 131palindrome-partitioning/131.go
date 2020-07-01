package main

func main() {

}

// 131. 分割回文串 https://leetcode-cn.com/problems/palindrome-partitioning/
func partition(s string) [][]string {

	curList := make([]string, 0)
	result := make([][]string, 0)
	dfs(curList, s, &result)
	return result

}

/*
	dfs 没想到没有超时 8ms 5.3MB
*/
func dfs(curList []string, left string, result *[][]string) {

	if left == "" {
		// tmp := make([]string, len(curList))
		// copy(tmp, curList)
		curList = append(curList[:0:0], curList...) //https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
		*result = append(*result, curList)
		return
	}

	for i := range left {

		if ispalindrome(left[:i+1]) {
			// golang的代码优势就在 这里 curList通过append传入后，传入的是一个新的slice，不会影响当前的 curList，无需回溯
			dfs(append(curList, left[:i+1]), left[i+1:], result)
		}

	}

	return
}

func ispalindrome(s string) bool {
	count := len(s)
	if count == 1 {
		return true
	}
	for i := 0; i < count/2; i++ {
		if s[i] != s[count-i-1] {
			return false
		}
	}

	return true
}
