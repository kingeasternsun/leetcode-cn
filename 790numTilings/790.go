package main

import "fmt"

func main() {
	fmt.Println(numTilings(4))
	fmt.Println(numTilings(5))
	fmt.Println(numTilings(10)) //1255
}

func numTilings1(N int) int {

	mod := 1000000007

	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	if N == 2 {
		return 2
	}

	dp := make([]int, N+1)
	dp1 := make([]int, N+1) //表示第i层只有一个瓷砖
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*dp1[i-1]) % mod // dp[i-1]:表示横着铺一个X，dp[i-2]表示两个竖着的X， dp1[i-1]表示 两种L
		dp1[i] = (dp[i-2] + dp1[i-1]) % mod            // dp[i-2]:表示一个L 放在整齐的前i-2上， dp[i-1]表示X竖着插在之前的L上
	}

	return dp[N]
}

//优化版本
func numTilings(N int) int {

	mod := 1000000007

	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	if N == 2 {
		return 2
	}

	dp := make([]int, 4) // dp[0]表示 dp[i-1] ;dp[1]表示dp[i];  dp[2] 表示 dp1[i]
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= N; i++ {
		dp[3] = dp[1] //保存之前的dp[1]
		// dp[i] = (dp[i-1] + dp[i-2] + 2*dp1[i-1]) % mod // dp[i-1]:表示横着铺一个X，dp[i-2]表示两个竖着的X， dp1[i-1]表示 两种L
		dp[1] = (dp[1] + dp[0] + 2*dp[2]) % mod
		// dp1[i] = (dp[i-2] + dp1[i-1]) % mod            // dp[i-2]:表示一个L 放在整齐的前i-2上， dp[i-1]表示X竖着插在之前的L上
		dp[2] = (dp[0] + dp[2]) % mod
		dp[0] = dp[3]

	}

	return dp[1]
}
