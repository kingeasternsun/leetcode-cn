package pivotinteger

import "sort"

func pivotInteger(n int) int {
	sum := (n) * (n + 1) >> 1
	ret := sort.Search(n+1, func(i int) bool {
		cusSum := i * (i + 1) >> 1
		return cusSum >= sum+i-cusSum
	})
	cusSum := ret * (ret + 1) >> 1
	if cusSum == sum+ret-cusSum {
		return ret
	}

	return -1
}
