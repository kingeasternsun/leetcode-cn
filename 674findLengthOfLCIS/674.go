package leetcode

func findLengthOfLCIS(nums []int) int {

	res := 0

	tmpLen := 0
	for i, v := range nums {
		if i == 0 {
			tmpLen = 1
			res = 1
			continue
		}

		if v > nums[i-1] {
			tmpLen++
		} else {
			if tmpLen > res {
				res = tmpLen
			}

			tmpLen = 1
		}

	}

	if tmpLen > res {
		res = tmpLen
	}

	return res

}
