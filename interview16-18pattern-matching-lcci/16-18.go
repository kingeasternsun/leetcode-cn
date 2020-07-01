package main

func checkMatch(pattern string, value string, aStr, bStr string, aLen, bLen int) bool {

	vStartID := 0
	vEndID := 0
	for _, p := range pattern {

		if p == 'a' {
			vEndID = vStartID + aLen
			if value[vStartID:vEndID] != aStr {
				return false
			}
		} else {
			vEndID = vStartID + bLen
			if value[vStartID:vEndID] != bStr {
				return false
			}
		}

		vStartID = vEndID
	}

	return true
}

func getFirstPos(pattern string, a rune) int {

	for i, p := range pattern {
		if p == a {
			return i
		}
	}

	return -1
}

/*

1. 根据a对应字符串的长度可以计算除b对应字符串的长度，并且根据patern可以得到b的开始位置。反向同理。
2. 要特别处理 a或b其中一个数量为0的情况。
3. 特殊处理 value 为空，这个时候 如果只有 a或b的情况

*/
func patternMatching(pattern string, value string) bool {

	pLen := len(pattern)
	vLen := len(value)
	if pLen == 0 && vLen == 0 {
		return true
	}
	if pLen == 0 && vLen > 0 {
		return false
	}

	aCount := 0
	bCount := 0

	for i := 0; i < pLen; i++ {
		if pattern[i] == 'a' {
			aCount++
		} else {
			bCount++
		}
	}

	if aCount == 0 && bCount == 0 {
		return false
	}

	//空字符串, 模式里面要么只有a或只有b
	if vLen == 0 {

		if aCount > 0 && bCount > 0 {
			return false
		}
		return true
	}

	if aCount == 0 {

		if bCount == 1 {
			return true
		}

		if vLen%bCount > 0 {
			return false
		}

		bLen := vLen / bCount
		return checkMatch(pattern, value, "", value[:bLen], 0, bLen)

	}

	if bCount == 0 {

		if aCount == 1 {
			return true
		}
		if vLen%aCount > 0 {
			return false
		}

		aLen := vLen / aCount
		return checkMatch(pattern, value, value[:aLen], "", aLen, 0)
	}

	maxALen := vLen / aCount // max length of 'a' resprent substr
	maxBLen := vLen / bCount // max length of 'b' resprent substr

	for i := 0; i <= maxALen; i++ {

		aStr := value[:i] //substr of 'a' resprent

		//快速剪枝
		if i > 0 && pattern[0] == pattern[pLen-1] && (value[vLen-i:] != aStr) {
			continue
		}

		if (vLen-i*aCount)%bCount > 0 {
			continue
		}

		j := (vLen - i*aCount) / bCount

		pos := getFirstPos(pattern, 'b')
		bStr := value[pos*i : pos*i+j]
		if i == j && aStr == bStr {
			continue
		}

		if checkMatch(pattern, value, aStr, bStr, i, j) {
			return true
		}

	}

	//b在前面
	for i := 0; i <= maxBLen; i++ {

		bStr := value[:i] //substr of 'b' resprent
		//快速剪枝
		if i > 0 && pattern[0] == pattern[pLen-1] && (value[vLen-i:] != bStr) {
			continue
		}

		if (vLen-i*bCount)%aCount > 0 {
			continue
		}

		j := (vLen - i*bCount) / aCount

		pos := getFirstPos(pattern, 'a')
		aStr := value[pos*i : pos*i+j]
		// aStr := value[i : i+j]
		if i == j && aStr == bStr {
			continue
		}

		if checkMatch(pattern, value, aStr, bStr, j, i) {
			return true
		}

	}

	return false
}
