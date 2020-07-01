package main

// https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
// 1312. 让字符串成为回文串的最少插入次数
/*
   dp[i][j] 跟 dp[i+1][j-1] ,dp[i][j-1],dp[i+1][j] 有关，最终终止到i>=j
   - 二位数组中，dp[0][0]到dp[sLen-1][sLen-1]的对角线以及左下方的都是0
   - 每次和对角线平行右移更新
*/

//4ms,6.6MB
func minInsertions1(s string) int {
	sLen := len(s)

	if sLen == 0 {
		return 0
	}

	// return dp(s, 0, sLen-1)

	arr := make([][]int, sLen)
	for i := 0; i < sLen; i++ {
		arr[i] = make([]int, sLen)
	}

	// 对角线右边的斜线
	for c := 1; c < sLen; c++ {
		i := 0 // arr[0][i]
		j := c
		for j < sLen {

			if s[i] == s[j] {
				arr[i][j] = arr[i+1][j-1]
			} else {
				arr[i][j] = min(arr[i+1][j]+1, arr[i][j-1]+1)
			}
			// return min(dp(s, i+1, j)+1, dp(s, i, j-1)+1)

			i++
			j++

		}
	}

	return arr[0][sLen-1]

}

/*
  - 从头到尾，每次计算步长为c的子传的最小修改数目
  - 从第i个字符，步长为step的字串
			if s[i] == s[i+c] {
				curArr[i] = prePreArr[i+1]
			} else {
				curArr[i] = min(preArr[i]+1, preArr[i+1]+1)
			}
  -  使用三个数组，一个保存上上步长的结果，一个保存上个步长的结果，一个保存当前的，每次迭代更新
*/
// 4ms 2.3MB
func minInsertions(s string) int {
	sLen := len(s)

	if sLen == 0 {
		return 0
	}

	// return dp(s, 0, sLen-1)

	prePreArr := make([]int, sLen)
	preArr := make([]int, sLen)
	curArr := make([]int, sLen)

	// 对角线右边的第c个平行斜线
	for c := 1; c < sLen; c++ {
		// i := 0 // arr[0][i]
		i := 0 //行数
		for i < sLen-c {

			// 想象一个虚拟的列 i+c
			if s[i] == s[i+c] {
				curArr[i] = prePreArr[i+1]
			} else {
				curArr[i] = min(preArr[i]+1, preArr[i+1]+1)
			}
			i++

		}

		tmp := prePreArr
		prePreArr = preArr
		preArr = curArr
		curArr = tmp

	}

	return preArr[0]

}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// 递归超时
func dp(s string, i, j int) int {
	if i >= j {
		return 0
	}

	if s[i] == s[j] {
		return dp(s, i+1, j-1)
	}

	return min(dp(s, i+1, j)+1, dp(s, i, j-1)+1)

}
