/*
 * @Description:1300
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 17:12:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 18:09:39
 * @FilePath: \1300findBestValue\1300.go
 */
package leetcode

import "sort"

/*
 分别找小于 target 或 大于targe的组合
 1. arr排序
 2. 求前缀和
 3. 二分查找
*/

func findBestValue(arr []int, target int) int {

	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)
	preSum := make([]int, len(arr))
	for i := range arr {
		preSum[i] = arr[i]
		if i > 0 {
			preSum[i] += preSum[i-1]
		}
	}

	res, diff := 0, target
	for i := arr[0] - 1; i <= arr[len(arr)-1]; i++ {
		id := sort.SearchInts(arr, i+1) //查询比i大的
		sum := (len(arr) - id) * i
		if id > 0 {
			sum += preSum[id-1]
		}

		if sum == target {
			return i
		}

		if abs(sum-target) < diff {
			res = i
			diff = abs(sum - target)
		}

	}

	return res

}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

/*

func findBestValue(arr []int, target int) int {

	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)
	preSum := make([]int, len(arr))
	for i := range arr {
		preSum[i] = arr[i]
		if i > 0 {
			preSum[i] += preSum[i-1]
		}
	}

	//查找和小于target的最大 value
	beg, end, best := 0, arr[len(arr)-1], 0
	for beg < end {
		mid := (end-beg)/2 + beg
		id := sort.SearchInts(arr, mid)
		//将arr中 [id..]变为mid 后 数组和
		sum := (len(arr) - id) * mid
		if id > 0 {
			sum += preSum[id-1]
		}

		if sum == target {
			return mid
		}

		if sum < target {
			best, beg = mid, mid+1
		} else {
			end = mid - 1
		}

	}

}
*/
