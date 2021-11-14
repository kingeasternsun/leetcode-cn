/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-03 21:47:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-03 21:52:49
 * @FilePath: /7reverse/7.go
 */
package leetcode

import "math"

func reverse(x int) int {
	res := 0

	for x != 0 {
		m := x % 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && m > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && m < -8) {
			return 0
		}

		res = res*10 + m
		x = x / 10

	}

	return res

}
