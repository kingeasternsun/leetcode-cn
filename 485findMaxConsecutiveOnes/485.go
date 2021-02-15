/*
 * @Description: 485
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-15 22:48:14
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 22:52:16
 * @FilePath: /485findMaxConsecutiveOnes/485.go
 */

package leetcode

func findMaxConsecutiveOnes(nums []int) int {

	res, tmp := 0, 0
	for _, v := range nums {
		if v == 1 {
			tmp++
			if tmp > res {
				res = tmp
			}
		} else {
			tmp = 0
		}
	}

	return res
}
