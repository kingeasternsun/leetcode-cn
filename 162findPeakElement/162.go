/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 11:30:48
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 11:55:33
 * @FilePath: \162findPeakElement\162.go
 */
package leetcode

//https://talkgo.org/t/topic/1667/18 三分法
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	beg := 0
	end := len(nums)

	for beg < end-1 {

		lmid := (end-beg)/2 + beg
		rmid := (end-lmid)/2 + lmid

		if nums[lmid] >= nums[rmid] {
			end = rmid
		} else {
			beg = lmid
		}
	}
	if nums[beg] > nums[end] {
		return beg
	}

	return end

}
