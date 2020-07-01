//https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// 1326. 灌溉花园的最少水龙头数目
package main

/*

相似问题
45. 跳跃游戏 II https://leetcode-cn.com/problems/jump-game-ii/
1024. 视频拼接	https://leetcode-cn.com/problems/video-stitching/
*/

func minTaps(n int, ranges []int) int {

	if n == 0 {
		return 0
	}

	nums := make([]int, n+1)
	for i := range ranges {
		ID := i - ranges[i]
		if ID < 0 {
			ID = 0
		}
		nums[ID] = i + ranges[i]
	}

	return minSteps(nums)

}

func minSteps(nums []int) int {
	inLen := len(nums)
	if inLen <= 1 {
		return 0
	}

	var maxLen, preLen int
	steps := 0
	for i := range nums {
		if maxLen < i { //中间断掉了
			return -1
		}

		//提前返回
		if maxLen >= inLen-1 {
			steps++
			return steps
		}

		// 当前已经能够覆盖到maxLen,但是 i能覆盖的却小于 maxLen，肯定不考虑当前i
		if nums[i] <= maxLen {
			continue
		}

		//如果i<=preLen 那么就表明i可以替换i-1，并且可以负责更广的区域
		if i > preLen {
			preLen = maxLen
			steps++
		}
		maxLen = nums[i]

	}
	return steps
}
