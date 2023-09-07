package mindays

// 108ms 8.7mb
func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	check := func(day int) bool {
		bloomNum := 0
		flowerNum := 0
		for _, b := range bloomDay {
			if b > day {
				flowerNum = 0
				continue
			}

			flowerNum += 1
			// 可以做一束花了
			if flowerNum == k {
				bloomNum += 1
				flowerNum = 0

				if bloomNum == m {
					return true
				}
			}

		}

		return false
	}

	maxDay := 1
	for _, b := range bloomDay {
		if b > maxDay {
			maxDay = b
		}
	}

	left := 1
	right := maxDay
	for left < right {
		mid := (right-left)/2 + left
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}

// 1760
// 120ms 8.86mb
func minimumSize(nums []int, maxOperations int) int {

	check := func(mid int) bool {
		op := 0
		for _, n := range nums {
			op += (n - 1) / mid

			if op > maxOperations {
				return false
			}
		}

		return true
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	left := 1
	right := max
	for left < right {
		mid := (right-left)/2 + left
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}
