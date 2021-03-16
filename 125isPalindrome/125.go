/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-07-01 11:11:03
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 14:29:24
 * @FilePath: /125isPalindrome/125.go
 */
package leetcode

import "strings"

func isPalindrome(s string) bool {

	sLen := len(s)
	if sLen == 0 {
		return true
	}

	s = strings.ToLower(s)

	i := 0
	j := sLen - 1

	for i <= j {
		if !(('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9')) {
			i++
			continue
		}
		if !(('a' <= s[j] && s[j] <= 'z') || ('0' <= s[j] && s[j] <= '9')) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}

	return true
}
