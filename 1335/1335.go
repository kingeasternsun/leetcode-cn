package main

import (
	"math"
)

func minDifficulty(job []int, d int) int {

	jobLen := len(job)

	if len(job) < d {
		return -1
	}

	dp := make([][]int, jobLen+1) // dp[nJob][nDay]表示前nJob个任务分配到nDay天的困难度
	dp[0] = make([]int, d+1)
	for i := 0; i < d+1; i++ {
		dp[0][i] = math.MaxInt32 / 2
	}
	dp[0][0] = 0 //关键

	for nJob := 1; nJob <= jobLen; nJob++ {
		dp[nJob] = make([]int, d+1)
		dp[nJob][0] = math.MaxInt32 / 2

		//因为每天至少要有一个任务，所以nJob个任务不可能分配到nJob+1以上的天中
		for nDay := 1; nDay <= nJob && nDay < d+1; nDay++ {

			md := 0
			dp[nJob][nDay] = math.MaxInt32 / 2

			//这里 从后往前，就可以快速获取中当前天的最大值
			for j := nJob - 1; j >= nDay-1; j-- {
				if job[j] > md {
					md = job[j]
				}

				//比较 前j个任务分配到nDay-1天的困难程度 加上 剩余任务（分配到了nday天）的最大困难程度
				if dp[j][nDay-1]+md < dp[nJob][nDay] {
					dp[nJob][nDay] = dp[j][nDay-1] + md
				}

			}

		}

	}
	// fmt.Println(dp)

	return dp[jobLen][d]
}
