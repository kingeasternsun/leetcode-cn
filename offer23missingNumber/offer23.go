package leetcode

func missingNumber(nums []int) int {

	//二分法

	beg := 0
	end := len(nums) //这里要注意，用 len(nums)而不要用len(nums)-1

	for beg < end {

		mid := (beg + end) / 2
		if nums[mid] == mid {
			beg = mid + 1
		} else {
			end = mid
		}

	}
	return beg

}
