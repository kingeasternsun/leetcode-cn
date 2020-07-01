package main

/*
dp 记录以i为结尾的子数组个数
超时
*/
func subarraysDivByK1(A []int, K int) int {

	if K < 2 {
		return 0
	}

	dp := make([]int, len(A))

	for i := 0; i < len(A); i++ {

		curSum := 0
		exist := false
		j := i
		for ; j >= 0; j-- {
			curSum += A[j]
			if curSum%K == 0 {
				exist = true
				break
			}
		}

		if exist {

			if j > 0 {
				dp[i] = dp[j-1] + 1
			} else {
				dp[i] = +1
			}

		}

	}

	num := 0
	for i := range dp {
		num += dp[i]
	}

	return num
}

/* 计算 前i个的和，然后 %K 作为key，i作为value 放入到对应的map中
   相同key里面的pos之间的差值就肯定能整除K
*/
func subarraysDivByK(A []int, K int) int {

	if K < 2 {
		return 0
	}

	posMap := make(map[int][]int, 0)
	curSum := 0

	for i := range A {
		curSum += A[i]
		key := curSum % K
		if key < 0 {
			key = key + K
		}
		posMap[key] = append(posMap[key], i)
	}

	num := 0
	for key, pos := range posMap {

		num = num + len(pos)*(len(pos)-1)/2

		if key == 0 {
			num += len(pos)
		}

	}
	return num
}
