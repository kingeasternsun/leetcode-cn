/*
 * @Description:5673
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-07 14:37:42
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-07 14:47:21
 * @FilePath: \5673maximumScore\5673.go
 */
package leetcode

import "sort"

/*
贪心 ，尽量让留在地上的石头堆数最多
*/
func maximumScore(a int, b int, c int) int {

	stones := []int{a, b, c}
	store := 0
	for {

		sort.Ints(stones)
		if stones[0] == 0 && stones[1] == 0 {
			return store
		}

		// a + b <= c , (a,c)(b,c)这样两两的去拿
		if stones[0]+stones[1] <= stones[2] {
			return store + stones[0] + stones[1]
		}

		//选次大的和最大的
		stones[1], stones[2] = stones[1]-1, stones[2]-1
		store++

	}
	return store
}

//
