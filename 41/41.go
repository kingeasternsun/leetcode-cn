package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/first-missing-positive/
//41. 缺失的第一个正数

/*
- 数组长度为1单独考虑
- 如果数组中的数值都大于 nums的长度sLen ，那么结果肯定是 1。也就是数值都大于2，结果肯定为1.
- 如果数组中的数值都小于等于 0 ，那么结果肯定是 1
- 如果有一部分数在(0,sLen),一部分在别的范围，也就是说在[1,sLen]的数的个数小于sLen，无法填满长度为sLen的桶，必然有空隙，这个空隙就是

*/

func main() {
	a := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(a))
	// m := make(map[int]string, 0)
	// m[0] = "aa"
	// m[1] = "bb"
	// data, _ := json.Marshal(m)
	// fmt.Println(string(data))
	// ta := `
	// :{"did":"24dfa7d316dc0000c410010004000000","sf_cap":"{\"total\":60,\"left\":{\"24dfa7d316dc0000c610010001000000\":30,\"24dfa7d316dc0000c610010002000000\":30,\"24dfa7d316dc0000c610010003000000\":30,\"24dfa7d316dc0000c610010004000000\":30,\"24dfa7d316dc0000c610010005000000\":30,\"24dfa7d316dc0000c610010006000000\":30,\"24dfa7d316dc0000c610010007000000\":30,\"24dfa7d316dc0000c610010008000000\":30,\"24dfa7d316dc0000c610010009000000\":30,\"24dfa7d316dc0000c61001000a000000\":30,\"24dfa7d316dc0000c61001000b000000\":30,\"24dfa7d316dc0000c61001000c000000\":30,\"2767926320bb0000c610010002000000\":30,\"2767926320bb0000c610010003000000\":30,\"2767926320bb0000c610010004000000\":30,\"2767926320bb0000c610010005000000\":30,\"2767926320bb0000c610010006000000\":30,\"2767926320bb0000c610010007000000\":30,\"2767926320bb0000c610010008000000\":30,\"2767926320bb0000c610010009000000\":30,\"2767926320bb0000c61001000a000000\":30,\"2767926320bb0000c61001000b000000\":30,\"b72311d930f50000c61001000a000000\":30}}","sf_result":0,"sf_week":"2020-02-17","uuid":"5c567951-5317-11ea-91fc-00163e02ae23"}

	// `

	// fmt.Println(len(ta))
}

//0ms,2.2MB
func firstMissingPositive(nums []int) int {

	sLen := len(nums)
	if sLen == 0 {
		return 1
	}

	if sLen == 1 {
		if nums[0] == 1 {
			return 2
		}
		return 1
	}

	i := 0
	for i < sLen {

		for nums[i] >= 1 && nums[i] <= sLen && nums[i]-1 > i {
			tmpID := nums[i] - 1

			//如果给要交换的地方的数值一样，就跳出来，不然这里就会一直死循环了
			if nums[i] == nums[tmpID] {
				break
			}
			nums[i], nums[tmpID] = nums[tmpID], nums[i] //这里交换几次，整个程序 剩余可交换次数就少几次 ；所以整体还是O(n)
			// i++
			// fmt.Println(nums)
		}

		if nums[i] < 1 || nums[i] > sLen {
			i++
			continue
		}

		// 到达这里的时候，
		// 要么 nums[i]-1 <= i;
		// 要么 nums[nums[i]-1]=num[i] ;n所以 调用下面的 nums[nums[i]-1] = nums[i]也就没有问题了
		nums[nums[i]-1] = nums[i]
		i++

	}

	// fmt.Println(nums)

	for i := range nums {
		if i != nums[i]-1 {
			return i + 1
		}
	}

	return nums[sLen-1] + 1
}
