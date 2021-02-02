package leetcode

/*
二分法
https://github.com/kingeasternsun/leetcode-cn
*/
func longestSubarray(nums []int) int {

	start := 0
	end := len(nums)
	for start < end {
		if start == end-1 {
			if judege(nums, end) {
				return end - 1
			}
			if judege(nums, start) {
				return start - 1
			}
			return start - 1
		}

		mid := (end-start)/2 + start
		if judege(nums, mid) {
			start = mid
		} else {
			end = mid - 1
		}

	}
	if start > 0 {
		return start - 1
	}
	return start

}

func judege(nums []int, winLen int) bool {
	numCnt := [2]int{}
	for i := 0; i < winLen; i++ {
		numCnt[nums[i]]++
	}

	if numCnt[0] <= 1 {
		return true
	}

	for i := winLen; i < len(nums); i++ {
		numCnt[nums[i-winLen]]--
		numCnt[nums[i]]++
		if numCnt[0] <= 1 {
			return true
		}

	}

	return false
}
