package main

//https://leetcode-cn.com/problems/jump-game-ii/
//45. 跳跃游戏 II

// func main() {
// 	nums := []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}

// 	t := time.Now()
// 	fmt.Println(jump(nums))
// 	fmt.Println(time.Now().Sub(t))
// }

func jump(nums []int) int {
	for i := range nums {
		nums[i] = i + nums[i]
	}

	return minSteps(nums)
}

func minSteps(nums []int) int {
	inLen := len(nums)
	if inLen == 0 || inLen == 1 {
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
