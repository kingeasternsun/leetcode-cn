package main

func isSubsequence(s string, t string) bool {

	sLen := len(s)
	tLen := len(t)

	i := 0
	j := 0
	for i < sLen {

		for j < tLen && t[j] != s[i] {
			j++
		}

		if j == tLen {
			return false
		}

		i++
		j++

	}

	return true
}
