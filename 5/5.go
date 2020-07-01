package main

//https://leetcode-cn.com/problems/longest-palindromic-substring/
//5. 最长回文子串
func longestPalindrome(s string) string {

	sLen := len(s)
	if sLen == 0 || sLen == 1 {
		return s
	}

	var maxLen, maxBeg, maxEnd int
	for i := range s {

		right := i //记录连续等于当前s[I]的右边界
		for right < sLen && s[right] == s[i] {
			right++
		}

		//连续s[i]的左边
		left := i - 1
		// fmt.Println(left, right)

		//分别取两边看是否相同
		for left >= 0 && right < sLen && s[left] == s[right] {
			left--
			right++
		}

		// fmt.Println("x", left, right, maxLen)

		if right-left-1 >= maxLen {
			maxLen = right - left - 1
			maxBeg = left + 1
			maxEnd = right
			// fmt.Println("x2", maxBeg, maxEnd)
		}

	}

	return s[maxBeg:maxEnd]
}
