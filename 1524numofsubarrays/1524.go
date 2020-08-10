package main

// 参考 https://www.bilibili.com/video/BV1ez4y1D7kP
func numOfSubarrays(arr []int) int {

	ans, odd, even := 0, 0, 0

	for i := range arr {
		//当前是偶数，那么以当前数字结尾且和为奇数的子串不变，那么以当前数字结尾且和为偶数的子串比之前加1（只有当前数字）
		if (arr[i] & 1) == 0 {
			even = even + 1
		} else { //当前奇数，那么以当前数字结尾且和为奇数的子串个数 等于 上一次迭代中和为偶数的子串个数一样 + 1（只有当前数字）
			// 以当前数字结尾且和为偶数的子串个数跟上一次迭代中和为奇数的子串个数一样
			odd, even = even+1, odd
		}

		ans = ans + odd
	}

	return ans % (1000000000 + 7)
}
