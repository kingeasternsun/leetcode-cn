/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-01-24 16:17:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-27 14:32:24
 * @FilePath: /395longestSubstring/longestSubstring.go
 */
package main

// 二分加滑动窗口
func longestSubstring(s string, k int) int {
	if k <= 1 {
		return len(s)
	}

	beg, end := k, len(s)
	best := 0

	for beg <= end {
		mid := (end-beg)/2 + beg
		if match([]byte(s), k, mid) {
			best, beg = mid, mid+1
		} else {
			end = mid - 1
		}

	}

	return best
}

func match(s []byte, k int, mid int) bool {

	cnt := [26]int{}
	maxCnt := 0
	for i := 0; i < mid; i++ {
		cnt[(s[i]-'a')]++
		if cnt[(s[i]-'a')] > maxCnt {
			maxCnt = cnt[(s[i] - 'a')]
		}
	}

	if maxCnt < k {
		return false
	}
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && cnt[i] < k {
			return false
		}
	}

	for i := k; i < len(s); i++ {
		cnt[(s[i]-'a')]++
		cnt[(s[i-k]-'a')]--
		if cnt[(s[i-k]-'a')] < 0 && cnt[(s[i-k]-'a')] < k {
			return false
		}

	}

	return true
}
