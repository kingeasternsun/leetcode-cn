/*
 * @Description: 866
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-16 11:24:32
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 14:21:30
 * @FilePath: /866primePalindrome/866.go
 */
package leetcode

import (
	"math"
)

/*
1. 构建回文数字
2. 然后判断是否是素数
*/
func primePalindrome(N int) int {

	numLen := getIntLength(N)
	if numLen == 1 {
		i := N
		for !isP(i) {
			i++
		}
		return i
	}

	//最小回文半径
	radMin := numLen / 2

	//回文半径的长度的范围
	for rad := radMin; rad < 5; rad++ {

		//长度为 rad 的数字最大值
		numMax := int(math.Pow(10, float64(rad)))

		//回文长度是偶数
		for leftHalfNum := int(math.Pow(10, float64(rad-1))); leftHalfNum < numMax; leftHalfNum++ {
			otherHalfNum := getReverseNum(leftHalfNum)
			n := leftHalfNum*numMax + otherHalfNum
			if n >= N && isP(n) {
				return n
			}
		}

		//回文长度是奇数
		for leftHalfNum := int(math.Pow(10, float64(rad-1))); leftHalfNum < numMax; leftHalfNum++ {
			otherHalfNum := getReverseNum(leftHalfNum)
			for i := 0; i < 10; i++ {
				n := leftHalfNum*numMax*10 + i*numMax + otherHalfNum
				if n >= N && isP(n) {
					return n
				}
			}
		}

	}

	return 0

}

//求出反向数字
func getReverseNum(n int) int {
	res := 0
	for n > 0 {
		res = res*10 + n%10
		n = n / 10
	}
	return res
}

func getIntLength(N int) int {

	res := 0
	for N > 0 {
		N = N / 10
		res++
	}
	return res
}

func isP(N int) bool {
	if N < 2 {
		return false
	}
	//快速判断
	if N&1 == 0 && N > 9 {
		return false
	}

	r := int(math.Sqrt(float64(N)))
	for d := 2; d <= r; d++ {
		if N%d == 0 {
			return false
		}
	}

	return true

}
