/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-01-24 16:17:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 15:27:18
 * @FilePath: \395longestSubstring\longestSubstring.go
 */
package main

// 暴力 滑动窗口
func longestSubstring1(s string, k int) int {
	if k <= 1 {
		return len(s)
	}

	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[(s[i]-'a')]++
	}

	for winLen := len(s); winLen > 0; winLen-- {
		if match([]byte(s), cnt, k, winLen) {
			return winLen
		}

		//避免重复计算
		cnt[(s[winLen-1]-'a')]--

	}

	return 0
}

func match(s []byte, cnt [26]int, k int, winLen int) bool {

	if windowMatch(cnt[:], k) {
		return true
	}

	for i := winLen; i < len(s); i++ {
		cnt[(s[i]-'a')]++
		cnt[(s[i-winLen]-'a')]--

		if windowMatch(cnt[:], k) {
			return true
		}

	}

	return false
}

func windowMatch(cnt []int, k int) bool {
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && cnt[i] < k {
			return false
		}
	}
	return true
}

/*
 如果窗口中 某个字符的总频率都小于k，那么期望的窗口肯定不包含这个字符
*/
func longestSubstring(s string, k int) int {
	if k <= 1 {
		return len(s)
	}

	return longestSubStringInWindow([]byte(s), k)
}

func longestSubStringInWindow(s []byte, k int) int {
	if len(s) < k {
		return 0
	}

	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[(s[i]-'a')]++
	}

	splitID := -1
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] < k { //以这个字符进行分割
			splitID = i
			break
		}
	}

	//表示这个里面的频率都大于等于k，满足条件了
	if splitID == -1 {
		return len(s)
	}

	left := longestSubStringInWindow(s[:splitID], k)
	right := longestSubStringInWindow(s[splitID+1:], k)
	if left > right {
		return left
	}
	return right

}
