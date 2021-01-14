package leetcode

func prefixesDivBy5(A []int) []bool {

	res := make([]bool, len(A))

	num := 0
	for i, v := range A {

		num = (num << 1) + v
		num = num % 5
		if num == 0 {
			res[i] = true
		} else {
			res[i] = false
		}

	}
	return res

}
