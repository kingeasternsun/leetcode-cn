package main

func main() {

	longestPalindrome("abccccdd")
}

/*

409. 最长回文串
https://leetcode-cn.com/problems/longest-palindrome/
1. 计算各个字符的个数，
2. 如果是偶数，就直接加上
3. 如果是奇数，就减一相加，同时标记多了一个奇数
*/

func longestPalindrome(s string) int {
	m := make(map[byte]int, 0)
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		m[s[i]]++
	}

	maxLen := 0
	has := false
	for _, v := range m {
		if v&1 == 1 { //奇数 ，减一相加
			maxLen = maxLen + v - 1
			has = true
		} else {
			maxLen = maxLen + v
		}

	}
	if has {
		maxLen++
	}

	return maxLen

}
