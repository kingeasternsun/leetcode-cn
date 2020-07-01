package main

//24ms 6.1MB
func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == n {
		return m
	}

	mLen := getBitLen(m)
	nLen := getBitLen(n)
	if nLen > mLen {
		return 0
	}

	//二进制位数相同，从左到右，直到遇到不一样的bit

	var result int
	for i := mLen; i >= 1; i-- {
		pos := (1 << uint(i-1))

		//第 i 个位置都是1
		if n&pos > 0 && m&pos > 0 {
			result += pos
			continue
		}

		//第 i 个位置都是0
		if n&pos == 0 && m&pos == 0 {
			continue
		}

		// 到这里 只会是 n的第 i 个位上是 1， m的第 i 个位上是 0
		return result

	}

	return result
}

func getBitLen(m int) int {

	bCount := 0
	for m > 0 {
		bCount++
		m = (m >> 1)
	}

	return bCount
}
