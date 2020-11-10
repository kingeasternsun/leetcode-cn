package main

import (
	"fmt"
)

func main() {
	//															 [1,1,3,3,2,3,7,9,9,8,8,3]
	nums := []int{5, 4, 7, 4, 3, 2} //     ->  [1,1,3,3,2,3,8,3,7,8,9,9]
	nextPermutation(nums)           //我的答案  [1 1 3 3 2 3 8 3 7 8 9 9]
	fmt.Println(nums)

}

/*
 1. 如果最后两个是从小到大，直接换以下就可以
 2. 向右移动，直到找到当前数值小于下一个数值的 i
 3. 如果找不到，说明当前是完全从高到低的，反转就可以
 4. 在 i 右边的逆序（包含i+1的递减序列）里面,找到大于 nums[i]的的最小值所在的位置 j
 5. 更换 i 和 j, 更换后，nums 从i+1 开始的子序列 依然是递减的
 6. 从 i+1反转就可以
*/

func nextPermutation(nums []int) {

	if len(nums) <= 1 {
		return
	}

	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	var i int
	for i = len(nums) - 2; i >= 0; i-- {

		if nums[i] < nums[i+1] {
			break
		}

	}

	//不存在下一个更大的排列
	if i < 0 {
		swap(nums)
		return
	}

	//如果最后两个是从小到大
	if i == len(nums)-2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		return
	}

	//在 nums[i+1..]中 查询 大于 num[i]的最小值 替换
	j := biSearch(i+1, nums, nums[i]) // 2 4 3 1  ,找到3              3 4  2 1

	nums[i], nums[j] = nums[j], nums[i] // 2 4 3 1 变为 3 4 2 1        4 3  2  1

	swap(nums[i+1:])
	return

}

func swap(nums []int) {

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {

		nums[i], nums[j] = nums[j], nums[i]

	}

}

//nums 是一个递减数列
func biSearch(beg int, nums []int, val int) int {
	end := len(nums) - 1
	for beg < end {

		//只剩下两个了 这里要特别处理
		if beg == end-1 {
			if nums[end] > val {
				return end
			}
			return beg
		}

		mid := (beg + end) / 2

		if nums[mid] > val {

			beg = mid
		} else {
			end = mid - 1
		}

	}
	return beg

}
