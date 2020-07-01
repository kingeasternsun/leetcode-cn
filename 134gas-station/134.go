package main

/*
134. 加油站 https://leetcode-cn.com/problems/gas-station/
1. 贪婪算法，从最开始有剩余油的加油站出发，类似 最大子数组和的问题
2. 先根据 gas 和 cost 计算从这个站离开的到下一个加油站的净剩余油量

看了官方题解 才搞清楚， 不好想到这样的答案

我们将出发站点 Ns 和无法到达站点 k 作为分隔点，将左式分成三个部分

可以这么理解，从Ns出发可以到达k-1，但是无法到达k，由于可以达到k-1，从Ns出发到达K-1时候，油量剩余 >=0; 无法到达 k，说明 gas[k-1]-cost[k-1]肯定小于0，
在执行算法的时候，tmpSum就会重置为0，然后从 k 开始重新累加。由于最终的 tmpSum是从 Ns开始的，所以 sum[k,Ns-1]也肯定小于0，
*/

func canCompleteCircuit(gas []int, cost []int) int {

	count := len(gas)
	if count == 0 {
		return -1
	}

	// left := make([]int, count)
	left := 0
	sum := 0    //计算所有left的总和，如果 sum<0 那么肯定无法走完全程
	tmpSum := 0 //表示从有剩余汽油的加油站出发，目前总剩余油量
	startID := 0

	for i := 0; i < count; i++ {
		left = gas[i] - cost[i]
		tmpSum += left
		sum += left

		// 由于 tmpSum 变为负数的时候就重置为0 ，所以在 执行 tmpSum += left 的时候 tmpSum肯定 >=0
		// 如果 tmpSum + left 变为了负数,那么当前的 left 肯定是负数，而且这个负数还很大，把之前的汽油都给消耗光了，
		// 所以只能放弃之前的，尝试从下一个开始
		if tmpSum < 0 {
			tmpSum = 0
			startID = i + 1
		}
	}

	if sum < 0 {
		return -1
	}

	// 走到这里，说明 sum>=0, 如果sum大于等于0，那么tmpSum肯定大于等于0 ，那么 startID 不会等于 count、
	// 反证法： 如果 tmpSum < 0, 假设这个 tmpSum是从第j个开始累加的，tmpSum = sum[j,count-1],那么 gas[j] - cost[j] 肯定要 >= 0 ，不然 就从j+1开始了
	// 因为 sum[0,j-1]+tmpSum = sum， 而 sum >=0, tmpSum<0 ,可推导出 sum[0,j-1]>0 ,
	// 但是 tmpSum却从j开始，又可以推断出 sum[0,j-1]< 0 , 故矛盾
	return startID

}
