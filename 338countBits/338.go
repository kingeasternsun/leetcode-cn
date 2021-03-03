/*
 * @Description: 338
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-03 11:00:27
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 11:07:40
 * @FilePath: \338countBits\338.go
 */
package leetcode

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}
