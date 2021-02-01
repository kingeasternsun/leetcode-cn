package leetcode

func canEat(candiesCount []int, queries [][]int) []bool {
	//前缀和
	preSum := make([]int64, len(candiesCount)+1) //记录第i类糖果之前的所有可吃糖果之和 也就是所有类型小于i的糖果之和
	for i := 1; i < len(candiesCount)+1; i++ {

		preSum[i] = int64(candiesCount[i-1]) + preSum[i-1]
	}

	res := make([]bool, len(queries))
	for j, q := range queries {
		days := q[1] + 1 //

		//到第 q[1] 天，至少能吃 q[1]个糖果
		//到第 q[1] 天，最多能吃 q[1]*q[2]个糖果
		if int64(days) > preSum[q[0]+1] {
			//即使每天吃一个，吃到了第 q[0]+1 类糖果，吃超过了
			res[j] = false
			continue
		}

		if int64(days*q[2]) <= preSum[q[0]] {
			//即使每天吃q[2]个，也只能吃到 q[0]-1类糖果，吃不到
			res[j] = false
			continue

		}

		res[j] = true

	}

	return res
}
