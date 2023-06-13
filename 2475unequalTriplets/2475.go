package unequaltriplets

func unequalTriplets(nums []int) int {

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	a := 0
	ret := 0

	for _, v := range m {
		b := len(nums) - a - v
		ret += a * v * b
		a += v

	}

	return ret

}
