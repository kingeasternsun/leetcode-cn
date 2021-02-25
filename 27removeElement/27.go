/*
 * @Description: 27
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:24:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:28:51
 * @FilePath: \27removeElement\27.go
 */
package leetcode

//双指针
func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	//这里trick的地方在于i的初始化值是 -1
	i, j := -1, 0
	for j < len(nums) {
		if nums[j] == val {
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
