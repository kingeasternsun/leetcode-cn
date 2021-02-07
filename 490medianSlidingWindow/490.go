/*
 * @Author: kingeasternsun
 * @Date: 2021-02-03 09:02:46
 * @LastEditTime: 2021-02-03 09:40:16
 * @LastEditors: kingeasternsun
 * @Description: In User Settings Edit
 * @FilePath: \490medianSlidingWindow\490.go
 */
package leetcode

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	if k > len(nums) {
		return nil
	}

	if k == 0 {
		return nil
	}
	res := []float64{}

	if k == 1 {
		for _, v := range nums {
			res = append(res, float64(v))
		}
		return res
	}

	winNums := make([]int, k)
	copy(winNums, nums[:k])
	sort.Ints(winNums)

	res = append(res, getMedian(winNums, k))
	for i := k; i < len(nums); i++ {
		id := sort.SearchInts(winNums, nums[i-k]) //移除i-k这个数字
		winNums[id] = nums[i]                     //加入 i上的数字

		//调整winNum[id]的数字位置，让winNums保持有序
		for id < k-1 && winNums[id] > winNums[id+1] {
			winNums[id], winNums[id+1] = winNums[id+1], winNums[id]
			id++
		}

		for id > 0 && winNums[id] < winNums[id-1] {
			winNums[id], winNums[id-1] = winNums[id-1], winNums[id]
			id--
		}

		res = append(res, getMedian(winNums, k))
	}
	return res
}

func getMedian(winNums []int, k int) float64 {
	return (float64(winNums[k/2]) + float64(winNums[(k-1)/2])) / 2
}
