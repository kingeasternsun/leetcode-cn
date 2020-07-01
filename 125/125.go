package main

import "strings"

func main() {
	isPalindrome("race a car")
}

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
