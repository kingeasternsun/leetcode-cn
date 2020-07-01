package main

import "math"

func longestCommonPrefix(strs []string) string {

	count := len(strs)
	if count == 0 {
		return ""
	}

	if count == 1 {
		return strs[0]
	}

	minLen := math.MaxInt32

	for i := 0; i < count; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	for i := 0; i < minLen; i++ {

		a := strs[0][i]

		for j := 1; j < count; j++ {
			if strs[j][i] != a {
				return strs[0][0:i]
			}
		}
	}

	return strs[0][0:minLen]
}
