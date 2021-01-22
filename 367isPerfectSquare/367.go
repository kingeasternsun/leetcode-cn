package leetcode

func isPerfectSquare(num int) bool {

	if num <= 0 {
		return false
	}

	beg := 1
	end := num

	for beg < end {
		mid := (end-beg)/2 + beg
		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}

	return beg*beg == num

}
