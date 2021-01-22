package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 {

		return 1
	}
	cache := make(map[int]float64, 0)

	return help(x, n, cache)
}

func help(x float64, n int, cache map[int]float64) float64 {
	if n == 1 {
		return x
	}

	if n == -1 {
		return 1.0 / x
	}

	//已经算计算过 直接返回
	if v, ok := cache[n]; ok {
		return v
	}

	res := help(x, n/2, cache)
	res *= res
	if n&1 == 1 {

		if n > 0 {
			res *= x

		} else {
			res /= x
		}
	}

	return res

}
