/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-01-24 16:17:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-28 13:42:20
 * @FilePath: /395longestSubstring/longestSubstring.go
 */
package main

// 暴力 滑动窗口
func longestSubstring(s string, k int) int {
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
