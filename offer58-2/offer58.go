package leetcode

func reverseLeftWords(s string, n int) string {
	chars := []byte(s)
	reverse(chars[:n])
	reverse(chars[n:])
	reverse(chars)
	return string(chars)

}

func reverse(chars []byte) {
	if len(chars) == 0 {
		return
	}

	for i := 0; i < len(chars)/2; i++ {

		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}

	return
}
