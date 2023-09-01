package shortestseq

// 双指针
func shortestSeq(big []int, small []int) []int {
	if len(big) == 0 {
		return nil
	}

	smallMap := make(map[int]struct{})
	for _, n := range small {
		smallMap[n] = struct{}{}
	}

	curMap := make(map[int]int, 0)

	left := 0
	right := 0

	min := len(big) + 1
	min_left := -1
	min_right := -1

	for left <= right && right <= len(big) {
		if len(curMap) < len(smallMap) {
			if right == len(big) {
				break
			}
			if _, ok := smallMap[big[right]]; ok {
				curMap[big[right]] += 1
			}
			right++
			continue
		}

		for left < right && len(curMap) == len(smallMap) {
			if right-left < min {
				min = right - left
				min_left = left
				min_right = right - 1
			}

			if curMap[big[left]] == 1 {
				delete(curMap, big[left])
			}
			left++
		}

	}

	if min == len(big)+1 {
		return nil
	}

	return []int{min_left, min_right}
}
