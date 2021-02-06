/*
 * @Description:二分加滑窗解决
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-02 09:10:08
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-03 10:12:40
 * @FilePath: \424characterReplacement\424.go
 */
package leetcode

/*
424. 替换后的最长重复字符
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过 104。
1. 二分法，每次定义最长的重复字符长度 winLen
2. 通过滑窗判断是否可以满足 ; 计算滑窗内各个字符的个数，然后得到频率最大的字符个数 cnt, 判断 cnt+k>= winLen ，就返回true
*/
func characterReplacement(s string, k int) int {
	if len(s) == 0 || k >= len(s)-1 {
		return len(s)
	}

	sb := []byte(s)
	start := k
	best := -1
	end := len(s)
	for start <= end {
		mid := (end-start)/2 + start
		if judge(sb, mid, k) {
			best, start = mid, mid+1
		} else {
			end = mid - 1
		}
	}

	return best

}

func judge(s []byte, winEnd, k int) bool {

	charCnt := make([]int, 26)
	for i := 0; i < winEnd; i++ {
		charCnt[s[i]-'A']++
	}

	if windowCanfill(charCnt, winEnd, k) {
		return true
	}

	for i := winEnd; i < len(s); i++ {
		charCnt[s[i]-'A']++
		charCnt[s[i-winEnd]-'A']--
		if windowCanfill(charCnt, winEnd, k) {
			return true
		}

	}

	return false

}

/*
1. 求窗口内频率最高的字符 c
2. 把窗口内c之外的字符个数如果小于等于k，那么就可以把其他的字符都换为c
*/
func windowCanfill(charCnt []int, winEnd, k int) bool {

	maxCnt := 0
	for _, v := range charCnt {
		if v > maxCnt {
			maxCnt = v
		}
	}
	if maxCnt+k >= winEnd {
		return true
	}
	return false
}
