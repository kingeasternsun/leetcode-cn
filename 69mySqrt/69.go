package leetcode

func mySqrt(x int) int {

	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}

	//二分法

	start := 1
	end := x
	for start < end {
		mid := (start-end)/2 + end
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			start = mid
		} else {
			end = mid - 1
		}
	}

	return start

}
