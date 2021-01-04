package leetcode

func fib(n int) int {

	cache := make(map[int]int, 0)
	return fibWithCache(n, cache)
}

func fibWithCache(n int, cache map[int]int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if v, ok := cache[n]; ok {
		return v
	}

	return fibWithCache(n-1, cache) + fibWithCache(n-2, cache)

}
