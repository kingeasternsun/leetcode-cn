package main

func majorityElement(nums []int) []int {

	cx, cy := 0, 0
	x, y := 0, 0

	for _, n := range nums {

		if (cx == 0 || n == x) && n != y {

			cx++
			x = n
		} else if cy == 0 || n == y {
			cy++
			y = n
		} else {
			cx--
			cy--
		}

	}

	res := make([]int, 0)
	cx = 0
	cy = 0
	for _, n := range nums {
		if n == x {
			cx++
		}

		if n == y {
			cy++
		}
	}

	if cx > len(nums)/3 {
		res = append(res, x)
	}

	if y != x && cy > len(nums)/3 {
		res = append(res, y)
	}

	return res

}
