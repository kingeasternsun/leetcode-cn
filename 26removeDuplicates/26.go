/*
 * @Description: 26
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:17:27
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:21:14
 * @FilePath: \26removeDuplicates\26.go
 */
package leetcode

//双指针
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i, j := 0, 1

	for j < len(nums) {
		if nums[j] == nums[j-1] {
			j++
			continue
		}

		i++
		if i != j {
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1
}
