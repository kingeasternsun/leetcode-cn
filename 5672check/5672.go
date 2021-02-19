/*
 * @Description:5672
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-07 14:17:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-07 14:52:15
 * @FilePath: \5672check\5672.go
 */
package leetcode

func check(nums []int) bool {

	if len(nums) <= 1 {
		return true
	}
	//找到中断点
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}
	if i == len(nums) {
		return true
	}
	if nums[i] > nums[0] {
		return false
	}

	for i = i + 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			return false
		}
		if nums[i] < nums[i-1] {
			return false
		}

	}

	return true
}
