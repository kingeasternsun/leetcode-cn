package main

//  169. 多数元素 https://leetcode-cn.com/problems/majority-element/
/*
 把每个数字当成一个人，从左到右让互相厮杀，如果一样就加血
*/
func majorityElement(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}

	liveNum := nums[0]
	liveCount := 1 //几滴血
	i := 1
	for i < count {

		//对手已经没血了
		if liveCount == 0 {
			liveNum = nums[i]
			liveCount = 1
			i++
			continue
		}

		//不一样，就互相砍一刀
		if nums[i] != liveNum {
			liveCount--
		} else {
			liveCount++
		}

		i++
	}

	return liveNum
}
