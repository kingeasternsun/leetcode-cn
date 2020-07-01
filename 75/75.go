package main

import "fmt"

func main() {
	nums := []int{1}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2}
	sortColors(nums)
	fmt.Println(nums)
	nums = []int{0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors1(nums []int) {

	sLen := len(nums)
	if sLen == 0 {
		return
	}

	a := make([]int, 3)

	for i := range nums {
		a[nums[i]]++
	}

	beg := 0
	end := 0
	for i := range a {
		end = beg + a[i]
		for j := beg; j < end; j++ {
			nums[j] = i
		}

		beg = end

	}
}

func sortColors(nums []int) {
	sLen := len(nums)
	if sLen == 0 {
		return
	}

	//left左边[0,left-1]的都是0， right 右边[right +1:]的都是2
	//[left:cur-1]的都是1
	left, cur := 0, 0 //
	right := sLen - 1

	for cur <= right && cur < sLen && right >= 0 {

		// 和 left互换，从而将当前的0移动到前面，将left所在1换到当前
		if nums[cur] == 0 {
			if cur != left {
				nums[left], nums[cur] = nums[cur], nums[left]
			}

			// 替换后left上就变为0，cur变为1，所以都向右移动
			left++
			cur++
			continue
		}

		if nums[cur] == 1 {
			cur++
			continue
		}

		// 如果当前位置等于2，就和right替换，
		nums[right], nums[cur] = nums[cur], nums[right]
		// 替换后right上就变为2，cur还不一定,所以right向左移动，cur不动
		right--

	}

	return
}
