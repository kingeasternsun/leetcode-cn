package main

import (
	"math"
	// "golang.org/x/net/idna"
)



func divide(dividend int, divisor int) int {

	var result int
	if dividend == math.MinInt64 && divisor == -1 {
		return math.MaxInt64
	}

	flag := 1

	if dividend > 0 && divisor < 0 {
		flag = -1
	}

	if dividend < 0 && divisor > 0 {
		flag = -1
	}

	n1 := int64(dividend)
	if n1 < 0 {
		n1 = -n1
	}

	n2 := int64(divisor)
	if n2 < 0 {
		n2 = -n2
	}

	for n2 <= n1 {

		var shift uint
		for n1 >= (n2 << shift) {
			shift++
		}
		n1 -= n2 << (shift - 1)
		result += int(1 << (shift - 1))

	}

	return result * flag

}
