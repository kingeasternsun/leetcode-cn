/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-26 16:25:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-26 16:25:49
 * @FilePath: \35searchInsert\35.go
 */
package leetcode

import "sort"

func searchInsert(nums []int, target int) int {

	return sort.SearchInts(nums, target)
}
