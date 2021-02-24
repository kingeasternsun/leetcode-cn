/*
 * @Description:1300
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 17:12:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 10:25:50
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

func findBestValue1(arr []int, target int) int {

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
	for i := 0; i <= arr[len(arr)-1]; i++ {
		id := sort.SearchInts(arr, i+1) //查询比i大的
		sum := (len(arr) - id) * i      //arr[id..]都变为 i
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
 分别找小于 target 或 大于targe的组合
 1. arr排序
 2. 求前缀和
 3. 双重二分查找  12ms 5.5MB
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

	//查询大于等于 target 的最小value
	//这里初始值比较重要，left要从0开始，right的最大值就是arr的最大值，更大是没用的
	left, right := 0, arr[len(arr)-1]
	best, sum := 0, 0
	for left <= right {
		mid := (right-left)/2 + left
		sum = computeSum(arr, preSum, mid)
		if sum == target {
			return mid
		}

		if sum > target {
			best, right = mid, mid-1
		} else {
			best, left = mid, mid+1
		}
	}

	best2 := 0
	if sum > target {
		best2 = best - 1
	} else {
		best2 = best + 1
	}

	sum2 := computeSum(arr, preSum, best2)
	if abs(sum-target) < abs(sum2-target) {
		return best
	}

	if abs(sum-target) > abs(sum2-target) {
		return best2
	}

	if best < best2 {
		return best
	}

	return best2
}

func computeSum(arr, preSum []int, mid int) int {
	id := sort.SearchInts(arr, mid+1) //查询大于mid的第一个位置
	sum := (len(arr) - id) * mid      // 大于mid的都变为mid，也就是 arr[id..]都变为 mid
	if id > 0 {
		sum += preSum[id-1]
	}

	return sum
}
