/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-11-14 11:41:01
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-04 09:42:25
 * @FilePath: \354maxEnvelopes\354.go
 */
package leetcode

import "sort"

//dp
/*
1. 首先按照信封的长排序
2. dp[i] 表示第i个信封中最多可以装多少个
3. dp[i] = max( dp[j] +1)  where j < i , and i.width >j.width ,i.height > j.height
*/
func maxEnvelopesdp(envelopes [][]int) int {

	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {

		return envelopes[i][0] < envelopes[j][0]

	})

	// fmt.Println(envelopes)

	dp := make([]int, len(envelopes)) //表示第几个信封里面最多套了多少个
	dp[0] = 1                         //

	maxEvs := 1
	for i, ev := range envelopes {

		if i == 0 {
			continue
		}

		tmpEvs := 0
		for j := i - 1; j >= 0; j-- {
			//如果ev 可以装下 第 j 个信封
			if ev[0] > envelopes[j][0] && ev[1] > envelopes[j][1] && dp[j] > tmpEvs {
				tmpEvs = dp[j]
			}
		}

		dp[i] = tmpEvs + 1
		if dp[i] > maxEvs {
			maxEvs = dp[i]
		}
	}

	return maxEvs

}

// 通过提前对信封按照 某一个方向排序，变为求最大递增子序列的问题
func maxEnvelopes(envelopes [][]int) int {

	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		//如果宽度w相同，那么两个就不能互相套进去，所以这里要把高度h大的放在前面。
		//不然的话，就有可能造成把宽度w相同，h小的装进h大的信封里面
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}

		return envelopes[i][0] < envelopes[j][0]
	})

	//接下来，就根据各个信封的h构成的数组中，查找最大的递增子序列,
	//转为 300 https://leetcode-cn.com/problems/longest-increasing-subsequence/
	hs := []int{}
	for _, v := range envelopes {

		id := sort.SearchInts(hs, v[1])
		if id == len(hs) {
			hs = append(hs, v[1])
		} else {
			hs[id] = v[1]
		}
	}

	return len(hs)

}
