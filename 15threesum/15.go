package main

import "sort"

/*
滑窗
记得去重，去重的时候还要注意 beg > i+1 && nums[beg] == nums[beg-1]
提前返回
*/
func threeSum(nums []int) [][]int {

	count := len(nums)
	if count < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for i, v := range nums {

		//去重
		if i > 0 && v == nums[i-1] {
			continue
		}

		//剩余两个要从后面找，所以和肯定要大于0
		if v > 0 {
			break
		}

		//提前判断返回
		if nums[0]+nums[1] > -v || nums[count-1]+nums[count-2] < -v {
			continue
		}

		beg := i + 1
		end := count - 1

		for beg < end {

			//跳过目标v对应的数组项 ,
			if beg > i+1 && nums[beg] == nums[beg-1] { //去重 beg=i+1的时候，可以一样
				beg++
				continue
			}

			if end < count-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			if nums[beg]+nums[end] == -v {
				tmpResult := []int{v, nums[beg], nums[end]}
				result = append(result, tmpResult)
				beg++
				end--
				continue
			}

			if nums[beg]+nums[end] > -v {
				end--
				continue
			}

			if nums[beg]+nums[end] < -v {
				beg++
				// continue
			}

		}

	}

	return result

}
