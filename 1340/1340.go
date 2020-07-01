//1340. 跳跃游戏 V

package main

import (
	"fmt"
	"sort"
)

func main() {

}

type Node struct {
	V, ID int
}

func maxJumps(arr []int, d int) int {
	aLen := len(arr)
	if aLen == 0 || aLen == 1 {
		return aLen
	}
	if d > aLen-1 {
		d = aLen - 1
	}
	sortArr := make([]Node, aLen)
	for i := range arr {
		sortArr[i].ID = i
		sortArr[i].V = arr[i]
	}

	sort.Slice(sortArr, func(i, j int) bool {
		if sortArr[i].V < sortArr[j].V {
			return true
		}

		return false
	})

	fmt.Println(sortArr)
	dp := make([]int, aLen)
	for i := range sortArr {

		// 分别向两边dp

		tmpID := sortArr[i].ID
		dp[tmpID] = 1
		if tmpID > 0 {
			// fmt.Println("left", tmpID-d, tmpID-1)
			// 向左边dp,最远找到d个
			for j := tmpID - 1; j >= (tmpID-d) && j >= 0; j-- {

				//遇到 大于等于当前高度 的就停止搜索
				if arr[j] >= arr[tmpID] {
					break
				}

				//这里 arr[j] < arr[i] ，所以 dp[j]肯定在dp[i]之前计算过了
				if dp[j]+1 > dp[tmpID] {
					dp[tmpID] = dp[j] + 1
				}

			}

		}

		if tmpID < aLen-1 {
			// fmt.Println("right", tmpID+1, tmpID+1)

			// 向右边dp,最远找到d个
			for j := tmpID + 1; j <= (tmpID+d) && j < aLen; j++ {

				//遇到 大于等于当前高度 的就停止搜索
				if arr[j] >= arr[tmpID] {
					break
				}

				//这里 arr[j] < arr[i] ，所以 dp[j]肯定在dp[i]之前计算过了
				if dp[j]+1 > dp[tmpID] {
					dp[tmpID] = dp[j] + 1
				}

			}
		}

	}

	// fmt.Println(dp)
	maxJump := dp[0]
	for i := range dp {
		if dp[i] > maxJump {
			maxJump = dp[i]
		}
	}

	return maxJump
}
