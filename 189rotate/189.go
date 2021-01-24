package leetcode

func rotate(nums []int, k int) {
	if len(nums) == 0 && len(nums) == 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)

}

func reverse(nums []int) {

	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums)/2; i++ {

		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return
}
