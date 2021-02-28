/*
 * @Description: 896
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-28 13:45:31
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-28 13:56:40
 * @FilePath: /896isMonotonic/896.go
 */

package leetcode

func isMonotonic(A []int) bool {

	if len(A) <= 2 {
		return true
	}

	// 先假设 是单调递增 也是单调递减；
	// 如果遇到 A[i]>A[i+1] 说明不是单调递增，如果遇到 A[i]<A[i+1]说明不是单调递减
	incre, decre := true, true
	for i := 1; i < len(A) && (incre || decre); i++ {
		if A[i-1] > A[i] {
			incre = false
		}

		if A[i-1] < A[i] {
			decre = false
		}
	}

	return incre || decre

}
