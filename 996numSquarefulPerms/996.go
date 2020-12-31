package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(numSquarefulPerms([]int{1, 17, 8}) == 2)
	// fmt.Println(numSquarefulPerms([]int{2, 2, 2}) == 1)
	// fmt.Println(numSquarefulPerms([]int{5, 11, 5, 4, 5}) == 2)
	// fmt.Println(numSquarefulPerms([]int{80, 1, 80, 1, 3, 6, 3}) == 3)

	fmt.Println(numSquarefulPerms1([]int{1, 17, 8}) == 2)
	fmt.Println(numSquarefulPerms1([]int{2, 2, 2}) == 1)
	fmt.Println(numSquarefulPerms1([]int{5, 11, 5, 4, 5}) == 2)
	fmt.Println(numSquarefulPerms1([]int{80, 1, 80, 1, 3, 6, 3}) == 3)

}

// 4, 5,5,5,11
// 5, 11,5,4,5
// 5,4,5,11,5
// 类似题目 943 980 996 1125 ,
// 重复的需要去重，先不考虑重复，最终统一处理 8ms 4.5MB
func numSquarefulPerms1(A []int) int {

	n := uint(len(A))

	if n == 1 {
		return 1
	}

	dp := make([][]int, n) //  二维 dp[j][mask]  ；mask 表示哪几个数字被选中 ， j表示最后一个是第j个数字
	var i uint = 0

	for ; i < n; i++ {
		dp[i] = make([]int, int(1<<n))
		dp[i][1<<i] = 1 //选择第 i 个
	}

	res := 0
	var mask uint = 0
	for ; mask < (1 << n); mask++ {

		//分别以j结尾
		var j uint = 0
		for ; j < n; j++ {

			//j不在mask里面就直接跳过
			if (1<<j)&mask == 0 {
				continue
			}

			//不满足条件的
			if dp[j][mask] == 0 {
				continue
			}

			// 下一个选中 k
			var k uint = 0
			for ; k < n; k++ {

				//已经选择过k了，就跳过
				if (1<<k)&mask > 0 {
					continue
				}

				if isPass(A[j] + A[k]) {
					newMask := (1 << k) | mask
					dp[k][newMask] = dp[k][newMask] + dp[j][mask] //这里是重点。因为这个状态可能重复出现

				}

			}

		}
	}

	// res := 0
	i = 0
	for ; i < n; i++ {

		res = res + dp[i][(1<<n)-1]

	}

	//计算阶乘
	factorial := make([]int, 20)
	factorial[0] = 1
	for i := 1; i < 20; i++ {
		factorial[i] = i * factorial[i-1]
	}

	m := make(map[int]int, 0) //用于去重
	for _, a := range A {
		m[a] = m[a] + 1
	}

	for _, cnt := range m {
		res = res / factorial[cnt]
	}

	return res
}

func isPass(a int) bool {

	b := int(math.Sqrt(float64(a)))
	return b*b == a
}

//dfs
func numSquarefulPerms(A []int) int {

	n := (len(A))

	if n == 1 {
		return 1
	}

	// sort.Ints(A)

	m := make(map[int]int, 0)
	for _, a := range A {
		m[a] = m[a] + 1
	}

	res := 0
	dfs(m, -1, 0, n, &res)
	return res
}

func dfs(m map[int]int, preValue, i, n int, res *int) {

	if i == n {
		*res = *res + 1
		return
	}

	for k, cnt := range m {
		if cnt == 0 {
			continue
		}

		if preValue < 0 || isPass(k+preValue) {

			m[k] = m[k] - 1

			dfs(m, k, i+1, n, res)

			m[k] = m[k] + 1

		}

	}

	return

}
