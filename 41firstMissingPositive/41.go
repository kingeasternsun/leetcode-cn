/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-15 19:16:51
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 20:40:07
 * @FilePath: /41firstMissingPositive/41.go
 */
package leetcode

/*
讲nums中的数字按照value放入到tmp中，如果超过了tmp所能放的下标范围，就抛弃
假设一个极端情况，nums中的每个数刚好填满 tmp （不包含0下标）；那么结果就是 len(tmp)+1
剩下的情况就是 nums至少有一个数字 要么 <=0 要么 >len(tmp) ；那么 tmp中就肯定有个缺口，第一个缺口的地方就是要找的数据
*/
func firstMissingPositive1(nums []int) int {
	tmp := make([]bool, len(nums)+1)
	for _, v := range nums {
		if v > 0 && v < len(nums)+1 {
			tmp[v] = true
		}
	}

	i := 1
	for ; i < len(nums)+1; i++ {
		if !tmp[i] {
			break

		}
	}

	return i

}

/*
 原地替换 类似 442 448
*/
func firstMissingPositive(nums []int) int {

	i := 0
	for i < len(nums) {
		rightPos := nums[i] - 1
		if rightPos < 0 || rightPos >= len(nums) {
			i++
			continue
		}

		if nums[i] == nums[rightPos] {
			i++
			continue
		}

		nums[i], nums[rightPos] = nums[rightPos], nums[i]
	}

	i = 0
	for ; i < len(nums); i++ {
		if nums[i] != i+1 {
			break
		}
	}

	return i + 1

}
