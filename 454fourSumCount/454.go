package main

func main() {

}

func fourSumCount_timeout(A []int, B []int, C []int, D []int) int {

	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	for _, a := range A {
		for _, b := range B {
			m1[a+b] = m1[a+b] + 1
		}
	}

	for _, a := range C {
		for _, b := range D {
			m2[a+b] = m2[a+b] + 2
		}
	}

	res := 0
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {

			if k1+k2 == 0 {
				res += v1 * v2
			}
		}
	}

	return res

}
func fourSumCount(A []int, B []int, C []int, D []int) int {

	m1 := make(map[int]int, 0)
	// m2 := make(map[int]int, 0)
	for _, a := range A {
		for _, b := range B {
			m1[a+b] = m1[a+b] + 1
		}
	}

	res := 0
	for _, a := range C {
		for _, b := range D {
			res += m1[-a-b]
		}
	}

	return res

}
