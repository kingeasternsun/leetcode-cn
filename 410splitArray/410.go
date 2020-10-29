package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}

func splitArray(nums []int, m int) int {

	cnt := len(nums)
	if m > cnt {
		m = cnt
	}

	dp := make([][]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		dp[i] = make([]int, m+1)

		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0

	for i := 1; i < cnt+1; i++ { //表示第i个位置，而不是下标

		tmpm := m
		if i < tmpm {
			tmpm = i //前i个元素最多只能分为i个
		}

		for j := 1; j < tmpm+1; j++ { //dp[i][j]

			v := math.MaxInt32
			rightSum := 0 //表示切割右边的和

			for k := i; k > 0; k-- { //在 k个 位置分割, 前k-1个 分割为j-1, 后面单独作为一个连续子数组

				if k < j { //提前返回
					break
				}

				rightSum = rightSum + nums[k-1]
				tmpMax := max(dp[k-1][j-1], rightSum)
				if tmpMax < v {
					v = tmpMax
				}

			}

			dp[i][j] = v

		}

	}

	return dp[cnt][m]

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

//二分法
func splitArrayBi(nums []int, m int) int {

	return 0
}

func canSplit(nums []int, m int, threshold int) int {

	curSum := 0
	curCnt := 0

	for i := 0; i < len(nums); i++ {

		if curSum+nums[i] <= threshold {
			curSum += nums[i]
			continue
		}

		curCnt++
		curSum = 0
		i = i - 1

		//现在已经有 curCnt 个大于 threshold 的分组了
		if curCnt > m {
			return curCnt
		}

	}

	return curCnt

}
