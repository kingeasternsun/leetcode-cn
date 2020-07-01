package main

/*
 记录以当前数值为子数组最后一个数字的最大值和最小值
*/
func maxProduct(nums []int) int {
	count := len(nums)

	if count == 0 {
		return 0
	}

	maxV := nums[0]
	tmpMax := nums[0]
	tmpMin := nums[0]
	for i := 1; i < count; i++ {
		newMax := max(max(nums[i], nums[i]*tmpMax), nums[i]*tmpMin)
		newMin := min(min(nums[i], nums[i]*tmpMax), nums[i]*tmpMin)
		if newMax > maxV {
			maxV = newMax
		}

		tmpMax, tmpMin = newMax, newMin
	}

	return maxV
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
