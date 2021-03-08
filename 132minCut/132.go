/*
 * @Description: 132
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-08 09:42:18
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-08 11:14:20
 * @FilePath: \132minCut\132.go
 */
package leetcode

/*
1. 时间换空间，计算任意两个位置是否是回文
2. 然后使用dp计算最小数量
*/
func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}

	isPaliLen := make([][]bool, len(s)) //保存以当前位置为右边节的最长回文串长度
	for i := 0; i < len(s); i++ {
		isPaliLen[i] = make([]bool, len(s))
	}

	var judgePaliFromMid = func(beg, end int, cnt int) {
		for beg >= 0 && end < cnt && s[beg] == s[end] {
			isPaliLen[beg][end] = true
			beg--
			end++
		}
	}

	for i := 0; i < len(s); i++ {
		judgePaliFromMid(i, i, len(s))
		judgePaliFromMid(i, i+1, len(s))
	}

	//以第i个位置为结尾的子串的最小分割次数
	dp := make([]int, len(s))
	for end := 1; end < len(s); end++ {
		dp[end] = len(s)
		for beg := 0; beg <= end; beg++ {
			if isPaliLen[beg][end] {
				//可以直接break
				if beg == 0 {
					dp[end] = 0
					break
				}
				if dp[beg-1]+1 < dp[end] {
					dp[end] = dp[beg-1] + 1
				}
			}
		}
	}

	return dp[len(s)-1]

}
