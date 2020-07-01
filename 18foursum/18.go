package main

import "sort"

//4ms 2.8MB
func fourSum(nums []int, target int) [][]int {

	count := len(nums)
	if count < 4 {
		return nil
	}
	result := make([][]int, 0)

	sort.Ints(nums)
	for j := range nums {

		//提前返回
		if j > count-4 {
			break
		}

		//去重
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}

		//最小的四个都大于target ，那么j，k  往后移动肯定都大于 target
		if nums[j]+nums[j+1]+nums[j+2]+nums[j+3] > target {
			break
		}
		// 如果 j 和后面最大的3个还小于 target ，那么 k 再怎么移动都不会等于target ,所以直接跳过k的移动，考虑下一个 j
		if nums[j]+nums[count-1]+nums[count-2]+nums[count-1] < target {
			continue
		}

		//加入上面两个优化 变为16ms

		for k := j + 1; k < count-2; k++ {

			//最小的四个都大于target ，那么j  往后移动k 或 beg ,end 肯定都大于 target
			if nums[j]+nums[k]+nums[k+1]+nums[k+2] > target {
				break
			}
			// 如果 j 和 k 以及后面最大的2个还小于 target ，
			// 那么 再怎么移动beg ,end 都不会等于target ,所以直接跳过beg ,end 的移动，考虑下一个 k
			if nums[j]+nums[k]+nums[count-2]+nums[count-1] < target {
				continue
			}

			//去重
			if k > j+1 && nums[k] == nums[k-1] {
				continue
			}

			tmpResult := twosum(nums, count, target, j, k)
			result = append(result, tmpResult...)

		}

	}

	return result

}

func twosum(nums []int, count int, target, j, k int) [][]int {
	result := make([][]int, 0)
	beg := k + 1
	end := count - 1

	for beg < end {

		//跳过目标v对应的数组项 ,
		if beg > k+1 && nums[beg] == nums[beg-1] { //去重 beg=k+1的时候，可以和nums[k]一样
			beg++
			continue
		}

		if end < count-1 && nums[end] == nums[end+1] {
			end--
			continue
		}

		if nums[beg]+nums[end]+nums[j]+nums[k] == target {
			tmpResult := []int{nums[j], nums[k], nums[beg], nums[end]}
			result = append(result, tmpResult)
			beg++
			end--
			continue
		}

		if nums[beg]+nums[end]+nums[j]+nums[k] > target {
			end--
			continue
		}

		beg++
		// continue

	}

	return result

}
