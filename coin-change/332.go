package main

import (
	"math"
	"sort"
)

func main() {
	coinChange([]int{1, 2, 3}, 9)
}

func coinChangedfs(coins []int, amount int) int {
	cLen := len(coins)
	if cLen == 0 {
		if amount == 0 {
			return 0
		}
		return -1
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	// fmt.Println(coins)
	result := math.MaxInt32

	dfs(coins, cLen, amount, 0, 0, &result)
	if result == math.MaxInt32 {
		return -1
	}

	return result

}

// coins组成 target
/*
https://leetcode-cn.com/problems/coin-change/solution/322-by-ikaruga/
*/
// 4ms 2.1MB
func dfs(coins []int, cLen int, target int, cIndex int, step int, result *int) {

	if target == 0 {

		if step < *result {
			*result = step
		}

		return
	}

	// 这里一定要判断
	if cIndex >= cLen {
		return
	}

	if target < 0 {
		return
	}

	// k+step < *resut 提前剪枝，如果步数已经比之前的大了，就没有必要再进行了
	for k := target / coins[cIndex]; k >= 0 && k+step < *result; k-- {
		// fmt.Println(cIndex)
		dfs(coins, cLen, target-k*coins[cIndex], cIndex+1, step+k, result)
	}

	return
}

// 4ms 2.1MB

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	cLen := len(coins)
	if cLen == 0 {
		return -1
	}
	result := dp(coins, cLen, amount)
	if result == math.MaxInt32 {
		return -1
	}

	return result
}
func min(i, j int) int {
	if i > j {
		return j
	}

	return i
}
func dp(coins []int, cLen int, target int) (result int) {

	if target == 0 {
		return 0
	}

	stepDp := make([]int, target+1) //stepDp[i]表示组成i所需要的最小隐蔽数目
	for i := 1; i < target+1; i++ {
		stepDp[i] = math.MaxInt32
	}

	for i := 1; i < target+1; i++ {

		for j := 0; j < cLen; j++ { // 最后一个硬币为coins【j】的情况
			if i-coins[j] >= 0 {
				stepDp[i] = min(stepDp[i], stepDp[i-coins[j]]+1)
			}

		}
	}

	return stepDp[target]
}
