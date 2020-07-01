package main

// dp 问题，类似 在固定时间内参加更多活动的题型
func mincostTickets(days []int, costs []int) int {

	dp := make([]int, 366) //第i天时候最少花多少钱

	for _, dayNum := range days {
		dp[dayNum] = -1 //标注旅游天
	}

	for dayNum := range dp {
		if dayNum == 0 {
			continue
		}

		//当前不旅游
		if dp[dayNum] == 0 {
			dp[dayNum] = dp[dayNum-1]
			continue
		}

		//当前是一天的票
		dp[dayNum] = dp[dayNum-1] + costs[0]

		//当前是7天票的最后一天
		if dayNum >= 7 {
			if dp[dayNum] > dp[dayNum-7]+costs[1] {
				dp[dayNum] = dp[dayNum-7] + costs[1]
			}
		} else {
			// 这里重要 ，不满足7天的也可以买7天的票
			if dp[dayNum] > costs[1] {
				dp[dayNum] = costs[1]
			}
		}

		//当前是30天票的最后一天
		if dayNum >= 30 {
			if dp[dayNum] > dp[dayNum-30]+costs[2] {
				dp[dayNum] = dp[dayNum-30] + costs[2]
			}
		} else {
			// 这里重要 ，不满足30天的也可以买30天的票
			if dp[dayNum] > costs[2] {
				dp[dayNum] = costs[2]
			}
		}

	}

	return dp[days[len(days)-1]]
}

func mincostTickets1(days []int, costs []int) int {

	dp := make([]int, 366) //第i天花多少钱

	for i, dayNum := range days {

		//当前是一天的票
		dp[dayNum] = dp[dayNum-1] + costs[0]

		//当前是7天票的最后一天
		if dayNum >= 7 {
			if dp[dayNum] > dp[dayNum-7]+costs[1] {
				dp[dayNum] = dp[dayNum-7] + costs[1]
			}
		} else {
			// 这里重要 ，不满足7天的也可以买7天的票
			if dp[dayNum] > costs[1] {
				dp[dayNum] = costs[1]
			}
		}

		//当前是30天票的最后一天
		if dayNum >= 30 {
			if dp[dayNum] > dp[dayNum-30]+costs[2] {
				dp[dayNum] = dp[dayNum-30] + costs[2]
			}
		} else {
			// 这里重要 ，不满足30天的也可以买30天的票
			if dp[dayNum] > costs[2] {
				dp[dayNum] = costs[2]
			}
		}

		// j := dayNum + 1
		if i+1 < len(days) {
			for j := dayNum + 1; j < days[i+1]; j++ {
				dp[j] = dp[dayNum]
			}
		}

	}

	return dp[days[len(days)-1]]
}
