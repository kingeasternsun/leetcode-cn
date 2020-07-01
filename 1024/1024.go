// 1024. 视频拼接	https://leetcode-cn.com/problems/video-stitching/

package main

/*

 可以转为 最小跳转数的那个问题
 //https://leetcode-cn.com/problems/jump-game-ii/
 //45. 跳跃游戏 II

*/

func videoStitching(clips [][]int, T int) int {
	nums := make([]int, T+1)
	var ID, v int
	for i := range clips {
		ID, v = clips[i][0], clips[i][1]
		if ID <= T && v > nums[ID] { //if ID > T,所以肯定拼不成，忽略
			nums[ID] = v
		}
	}

	return minSteps(nums, T)
}

func minSteps(nums []int, T int) int {
	inLen := len(nums)
	if inLen == 0 {
		return 0
	}

	var maxLen, preLen int
	steps := 1
	for i := range nums {
		if maxLen < i { //中间断掉了
			return -1
		}

		//提前返回
		if maxLen >= T {
			return steps
		}

		// 当前已经能够覆盖到maxLen,但是 i能覆盖的却小于 maxLen，肯定不考虑当前i
		if nums[i] <= maxLen {
			continue
		}

		//如果i<=preLen 那么就表明i可以替换i-1，并且可以负责更广的区域，所以只在i > prelen加一
		if i > preLen {
			// fmt.Println(i, nums[i])
			steps++
			preLen = maxLen
		} else {

			// fmt.Println(i, nums[i])

		}
		// preLen = maxLen
		maxLen = nums[i]
		// fmt.Println("t", i, nums[i], preLen, maxLen, steps)

	}
	return steps
}
