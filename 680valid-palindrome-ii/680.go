package main

func validPalindrome(s string) bool {
	count := len(s)
	if count == 0 {
		return true
	}

	beg := 0
	end := count - 1

	for beg < end {
		if s[beg] == s[end] {
			beg++
			end--
			continue
		}

		//遇到不相等的，删除左边或右边
		if isPalindome(s, beg+1, end) {
			return true
		}

		if isPalindome(s, beg, end-1) {
			return true
		}

		return false
	}

	return true

}

func isPalindome(s string, beg int, end int) bool {
	for beg < end {

		if s[beg] != s[end] {
			return false
		}
		beg++
		end--
	}
	return true
}
