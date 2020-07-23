package main

import "sort"

func main() {

	a := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}

	maxEnvelopes(a)
}

//dp
/*
1. 首先按照信封的长排序
2. dp[i] 表示第i个信封中最多可以装多少个
3. dp[i] = max( dp[j] +1)  where j < i , and i.width >j.width ,i.height > j.height
*/
func maxEnvelopes(envelopes [][]int) int {

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
