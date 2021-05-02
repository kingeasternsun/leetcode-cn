/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-01 19:47:12
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-01 20:35:53
 * @FilePath: /leetcode-cn/offer10-1/10.go
 */
package leetcode

func fib(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 1, 1
	for i := 2; i <= n; i++ {
		r = (p + q) % (1e9 + 7)
		p = q
		q = r
	}

	return r
}
