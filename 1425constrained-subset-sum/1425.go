package main

//https://leetcode-cn.com/problems/constrained-subset-sum/
// 92ms 9.5MB
func constrainedSubsetSum(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	dp := make([]int, len(nums))
	q := make([]int, 0) // to compute  dp[i-k]...dp[i] 的最大值
	for i := range nums {

		if i == 0 {
			dp[0] = nums[0]
			// q = append(q, 0)  //这里不能加，因为在 i 为1的时候也会加入
			continue
		}

		//虽然 q[0]对应的dp最大，但是超过了k的窗口 只能放弃
		if len(q) > 0 && i-q[0] > k {
			q = q[1:]
		}

		// add [i-1] to q
		qID := len(q) - 1
		q = append(q, i-1)
		for qID >= 0 {
			if dp[i-1] > dp[q[qID]] {
				qID--
			} else {
				break
			}
		}
		q[qID+1] = i - 1
		q = q[:qID+2]

		// max is dp[q[0]]

		dp[i] = max(dp[q[0]], 0) + nums[i]
		if dp[i] > maxSum {
			maxSum = dp[i]
		}

	}

	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
