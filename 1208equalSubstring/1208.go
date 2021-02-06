/*
 * @Description:leetcode: 1208
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 09:12:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 09:29:05
 * @FilePath: \557reverseWordse:\github\leetcode-cn\1208equalSubstring\1208.go
 */
package leetcode

//纯滑窗
func equalSubstring(s string, t string, maxCost int) int {

	tmpCost, maxLen := 0, 0
	beg, end := 0, 0 //左闭，右开区间
	//注意这里 条件是 beg<= end，beg也可以等于end,意味着从新开始新的区间. 例如s和t中相应字符的差值 都大于 maxCost 的情况
	for beg <= end && end < len(s) {

		if tmpCost <= maxCost { //当前窗口可以满足预算
			if end-beg > maxLen {
				maxLen = end - beg
			}
			//然后追加 end
			delTa := abs(int(s[end]) - int(t[end]))
			tmpCost += delTa
			end++
			continue
		}

		//去掉beg
		delTa := abs(int(s[beg]) - int(t[beg]))
		tmpCost -= delTa
		beg++

	}

	//这里不能漏掉，要再次判断下
	if tmpCost <= maxCost && (end-beg) > maxLen {
		maxLen = end - beg
	}
	return maxLen

}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//二分搜索加滑窗
func equalSubstringBinary(s string, t string, maxCost int) int {

	left, right, best := 0, len(s), 0
	for left <= right {
		mid := (left-right)/2 + right
		if judge(s, t, maxCost, mid) {
			best, left = mid, mid+1
		} else {
			right = mid - 1
		}
	}

	return best

}

func judge(s string, t string, maxCost int, winLen int) bool {

	winSum := 0
	for i := 0; i < winLen; i++ {
		winSum += abs(int(s[i]) - int(t[i]))
	}

	if winSum <= maxCost {
		return true
	}

	for i := winLen; i < len(s); i++ {
		winSum -= abs(int(s[i-winLen]) - int(t[i-winLen]))
		winSum += abs(int(s[i]) - int(t[i]))
		if winSum <= maxCost {
			return true
		}

	}
	return false
}
