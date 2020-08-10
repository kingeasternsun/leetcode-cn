package main

import "math"

// https://www.bilibili.com/video/BV1yz4y1D7p3

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func maxNumOfSubstrings(s string) []string {

	count := len(s)
	left := make([]int, 26)
	right := make([]int, 26)

	for i := 0; i < 26; i++ {
		left[i] = math.MaxInt32
	}

	for i := 0; i < 26; i++ {
		right[i] = math.MinInt32
	}

	for i := 0; i < count; i++ {

		left[s[i]-'a'] = min(i, left[s[i]-'a'])
		right[s[i]-'a'] = max(i, right[s[i]-'a'])
	}

	// 关键函数 extend
	//计算从第i个字母开始，满足条件的子串的右边边界,left 表示某个字母第一次出现的位置,right 表示某个字母最后一次出现的位置
	//保证下一次迭代，如果有相交，肯定是上一次迭代的子串
	extend := func(i int) int {

		p := right[s[i]-'a'] //最后出现的位置

		pos := i
		for pos < p {

			if left[s[pos]-'a'] < i {
				return -1
			}

			if right[s[pos]-'a'] > p {
				p = right[s[pos]-'a']
			}

			pos++
		}

		return p

	}

	res := make([]string, 0)
	last := -1
	for i := 0; i < count; i++ {

		// 只计算当前字母第一次出现的位置
		if i != left[s[i]-'a'] {
			continue
		}

		p := extend(i)
		if p == -1 {
			continue
		}
		if i > last {
			res = append(res, s[i:p+1])
		} else {
			res[len(res)-1] = s[i : p+1]
		}
		last = p

	}

	return res

}
