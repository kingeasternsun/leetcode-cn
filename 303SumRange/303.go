/*
 * @Description: 303
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-01 09:20:33
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-01 09:25:02
 * @FilePath: /303SumRange/303.go
 */
package leetcode

type NumArray struct {
	PreSum []int //前缀和
}

func Constructor(nums []int) NumArray {

	preSum := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			preSum[i] = nums[i]
		} else {
			preSum[i] = nums[i] + preSum[i-1]
		}
	}
	return NumArray{PreSum: preSum}
}

func (this *NumArray) SumRange(i int, j int) int {

	if i > j {
		return 0
	}

	if i == 0 {
		return this.PreSum[j]
	}
	return this.PreSum[j] - this.PreSum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
