package main

func countSubstrings(s string) int {

	count := len(s)
	if count <= 1 {
		return count
	}

	ans := count
	//奇数回文
	for i := 1; i < count; i++ {

		left := i - 1
		right := i + 1
		width := 0

		for left >= 0 && right < count {

			if s[left] == s[right] {
				width++
				left--
				right++
			} else {
				break
			}
		}

		ans += width

	}

	//偶数回文
	for i := 1; i < count; i++ {

		left := i - 1
		right := i
		width := 0

		for left >= 0 && right < count {

			if s[left] == s[right] {
				width++
				left--
				right++
			} else {
				break
			}
		}

		ans += width

	}

	return ans

}
