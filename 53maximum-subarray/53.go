package main

func maxSubArray(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0

	}

	maxV := nums[0]
	tmpV := nums[0] //以i为结尾的子数组的最大和
	for i := 1; i < count; i++ {

		if tmpV > 0 {
			tmpV = tmpV + nums[i]
		} else {
			tmpV = nums[i]
		}

		if tmpV > maxV {
			maxV = tmpV
		}
	}
	return maxV

}
