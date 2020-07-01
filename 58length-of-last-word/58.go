package main

import "strings"

//https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {

	s = strings.TrimRight(s, " ")
	pos := strings.LastIndex(s, " ")
	if pos == -1 {
		return len(s)
	}

	return len(s[pos+1:])
}
