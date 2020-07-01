package main

//55. 跳跃游戏 https://leetcode-cn.com/problems/jump-game/

//8ms 4.2MB
func canJump(nums []int) bool {

	var maxID int //表示当前可以到达的最大下标
	for i := range nums {

		if i > maxID {
			return false
		}

		if i+nums[i] > maxID {
			maxID = i + nums[i]
		}

	}

	return true

}
