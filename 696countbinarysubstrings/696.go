package main

func countBinarySubstrings(s string) int {

	left := 0
	right := 0
	sum := 0

	count := len(s)
	if count == 0 {
		return 0
	}

	leftbyte := s[0]
	left = 1

	i := 1
	for i < count && s[i] == leftbyte {
		i++
		left++
	}

	if i == count {
		return 0
	}

	//here  : s[i]!=leftbyte
	rightbyte := s[i]
	right = 1
	i++

	for i < count {

		for i < count && s[i] == rightbyte {
			i++
			right++
		}

		if i < count {
			sum = sum + min(left, right)

			leftbyte, rightbyte = rightbyte, leftbyte
			left = right
			right = 1
			i++
		}

	}
	sum = sum + min(left, right)
	return sum

}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
