package main

func repeatedSubstringPattern(s string) bool {

	count := len(s)
	if count <= 1 {
		return false
	}

	for i := 0; i < count/2; i++ {

		left := 0
		right := i + 1
		for right < count && s[left] == s[right] {
			left++
			right++
		}

		if right == count && (count%(i+1) == 0) {
			return true
		}

	}

	return false

}
