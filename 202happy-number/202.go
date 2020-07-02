package main

func isHappy(n int) bool {

	eMap := make(map[int]struct{}, 0)

	for {

		if _, ok := eMap[n]; ok {
			return false
		}

		eMap[n] = struct{}{}
		n = getSum(n)
		if n == 1 {
			return true
		}
	}

}

func getSum(n int) int {

	sum := 0
	for n > 0 {
		sum = (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
