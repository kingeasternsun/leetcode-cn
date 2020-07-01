package main

import (
	"strconv"
)

// d^n
func sqrtN(d, n int) int {

	r := 1
	for n != 0 {
		if n&1 == 1 {
			r = r * d
		}
		d *= d
		n /= 2
	}

	return r
}

//search first num thant less than n
func biSearch(d []int, n int) (c int) {
	beg := 0
	end := len(d) - 1

	for beg < end {

		mid := (beg + end) / 2
		if d[mid] < n {

			if mid+1 < len(d) && d[mid+1] >= n {
				return mid + 1
			}

			beg = mid + 1
		} else {
			end = mid - 1
		}

	}

	if beg > end {
		return 0
	}

	// if beg == end-1 {
	// 	return beg + 1
	// }

	if d[beg] < n {
		return beg + 1
	}

	return 0

}

// https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/
func atMostNGivenDigitSet(D []string, N int) int {
	if len(D) == 0 {
		return 0
	}
	d := make([]int, len(D))
	//判断D是否可以完全构成N，因为题目要求 小于 或 等于
	dMap := make(map[int]struct{}, 0)
	// canNot := false

	for i := range D {
		d[i], _ = strconv.Atoi(D[i])
		dMap[d[i]] = struct{}{}
	}

	// sort.Ints(d)
	//compute N length, and each number
	var nums []int
	for N > 0 {
		nums = append(nums, N%10)

		// if _, ok := dMap[N%10]; !ok {
		// 	canNot = true
		// }

		N = N / 10
	}

	if len(nums) > 1 {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
	if len(nums) == 0 {
		return 0
	}

	c := biSearch(d, nums[0])
	result := 0

	// 位数比N的位数小，肯定就小，计算D中字符最多组成 len(nums)-1 位数字的组合个数
	// 位数相同，最高位比nums[0]小的有c个，然后乘以 D中字符组成 len(nums)-1 位数字的组合个数；注意里面只能组成 len(nums)-1 位
	q := len(d)
	if q == 1 {
		result = (len(nums) - 1) + c*1
	} else {
		result = q*(1-sqrtN(q, len(nums)-1))/(1-q) + c*sqrtN(q, len(nums)-1)
	}

	//如果最高位在D中
	if _, ok := dMap[nums[0]]; ok {
		// tmpR := 1

		for i, tmpN := range nums {
			if i == 0 {
				continue
			}

			if tmpN == 0 {
				break
			}

			// search number of less than tmpN
			c := biSearch(d, tmpN)
			if c > 0 {
				// 位数相同，i之前的都相同，最高位比nums[i]小的有c个，然后乘以 D中字符组成 len(nums)-i-1 位数字的组合个数；注意里面只能组成 len(nums)-i-1 位
				result = result + c*sqrtN(q, len(nums)-i-1)
			}

			// tmpN large than any of d, so we can stop here
			//位数相同，i之前的都相同 ，如果i位比N的i位还要大，得到的数字肯定大，可以退出了
			if _, ok := dMap[tmpN]; !ok {

				break
			}

			if i == len(nums)-1 {
				result++
			}

		}

		if len(nums) == 1 {
			result++
		}

	}

	return result
}
