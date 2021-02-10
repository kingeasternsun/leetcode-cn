/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-09 17:58:29
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 13:08:30
 * @FilePath: /992subarraysWithKDistinct/992.go
 */
package leetcode

/*
类似 求和为特定value的子数组个数；滑窗加hash

思路
先求包含k个不同字符的最长区间
然后求固定右边界，包含k个不同字符的子数组个数

*/
func subarraysWithKDistinct(A []int, K int) int {
	res := 0
	cntMap := make(map[int]int, 0)

	//其实这一步可以省略 也不影响最终的结果
	for i := 0; i < K; i++ {
		cntMap[A[i]]++
	}
	// 左闭 右闭 区间 [left, right]
	left := 0
	right := K - 1
	for left <= right {
		if len(cntMap) <= K {
			if len(cntMap) == K {
				res += help(A, left, right, cntMap, K)
			}
			if right+1 == len(A) {
				return res
			}
			//向右移动
			right++
			cntMap[A[right]]++
		} else if len(cntMap) > K {
			cntMap[A[left]]--
			if cntMap[A[left]] == 0 {
				delete(cntMap, A[left])
			}
			left++
		}
	}

	return res
}

//　计算以right为结尾的满足条件的子数组有多少个
func help(A []int, left, right int, cntMap map[int]int, K int) int {
	tmpMap := make(map[int]int, 0) //记录移除的个数，这样就不会改变原有的cntMap

	cnt := 0
	for left <= right {
		cnt++
		//移动left
		tmpMap[A[left]]++
		if tmpMap[A[left]] == cntMap[A[left]] { //A【left】在窗口中要全部删除掉了
			break
		}

		left++
	}

	return cnt
}
