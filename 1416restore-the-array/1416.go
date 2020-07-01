package main

//参考 huahualeetcode
func numberOfArrays(s string, k int) int {

	kMod := int(1e9 + 7)
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '0' {
			dp[i] = 0
			continue
		}

		num := 0
		for j := i + 1; j <= len(s); j++ {
			num = num*10 + int(s[j-1]-'0')
			// num, _ = strconv.Atoi(s[i:])
			if num > k {
				break
			}

			dp[i] = (dp[i] + dp[j]) % kMod
		}

	}

	return dp[0]
}
