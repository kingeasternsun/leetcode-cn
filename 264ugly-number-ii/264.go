/*
 * @Author: your name
 * @Date: 2020-03-15 17:12:20
 * @LastEditTime: 2020-03-16 09:32:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\264\264.go
 */
package main

func main() {

	nthUglyNumber(3)
}

func min(i, j, k int) int {

	if j < i {
		i = j
	}

	if k < i {
		i = k
	}

	return i

}

func nthUglyNumber(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	u := make([]int, n+1)
	u[1] = 1

	u2 := 1
	u3 := 1
	u5 := 1

	i := 2
	for ; i <= n; i++ {

		u[i] = min(u[u2]*2, u[u3]*3, u[u5]*5)

		for u[u2]*2 <= u[i] && u2 <= n {
			u2++
		}

		for u[u3]*3 <= u[i] && u3 <= n {
			u3++
		}

		for u[u5]*5 <= u[i] && u5 <= n {
			u5++
		}

	}
	return u[n]

}
func nthUglyNumber2(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	// if n == 2 {
	// 	return 2
	// }

	u := make([]int, n+1)
	u[1] = 1

	for i := 2; i <= n; i++ {

		id2, v2 := help(u, i-1, u[i-1], 2)
		u[i] = v2
		id3, v3 := help(u, id2, u[i-1], 3)
		if v3 < u[i] {
			u[i] = v3
		}

		_, v5 := help(u, id3, u[i-1], 5)
		if v5 < u[i] {
			u[i] = v5
		}

	}

	return u[n]

}

//  从 curID 查询 乘以mg后 大于 value的最小值newVluae
func help(u []int, curID, value, mg int) (id int, newValue int) {
	for i := curID; i > 0; i-- {
		if u[i]*mg <= value {
			return i + 1, u[i+1] * mg
		}
	}

	return 1, u[1] * mg
}
