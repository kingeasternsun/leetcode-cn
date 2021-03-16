/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-16 16:10:57
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 18:52:43
 * @FilePath: /336palindromePairs/336.go
 */
package leetcode

/*
1. 提前处理，记录每个单词的各个前缀 ，各个后缀是不是回文
2. 对于每个单词，如果某个前缀是回文，就判断这个单词除去前缀后剩下的部分是否和其他单词对称
3. 对于每个单词，如果某个后缀是回文，就判断这个单词除去后缀后剩下的部分是否和其他单词对称
*/
func palindromePairs(words []string) [][]int {

	//判断某个单词的某个区间是否是回文
	judgePali := func(i, beg, end int) bool {
		for beg < end {
			if words[i][beg] != words[i][end] {
				return false
			}
			beg++
			end--
		}

		return true
	}

	//预处理，提前记录每个单词的逆序的map
	reverseMap := make(map[string]int, 0)
	preMap := make(map[int][]int, 0)  //记录每个单词哪些前缀是回文
	suffMap := make(map[int][]int, 0) //记录每个单词哪些后缀是回文

	for i, w := range words {
		reverseMap[reverse(w)] = i
		for j := 0; j < len(w); j++ {
			if judgePali(i, 0, j) {
				preMap[i] = append(preMap[i], j)
			}
			if judgePali(i, j, len(w)-1) {
				suffMap[i] = append(suffMap[i], j)
			}
		}
	}

	res := make([][]int, 0)
	for i, w := range words {

		if j, ok := reverseMap[w]; ok {
			if i > j {
				res = append(res, []int{i, j})
				res = append(res, []int{j, i})
			}
		}

		//对于回文的每个前缀
		for _, endID := range preMap[i] {
			//判断剩下的部分是否存在 反向的单词
			if j, ok := reverseMap[(w[endID+1:])]; ok {
				res = append(res, []int{j, i})
			}
		}

		//对于回文的每个后缀
		for _, endID := range suffMap[i] {
			//判断剩下的部分是否存在 反向的单词
			if j, ok := reverseMap[(w[:endID])]; ok {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func reverse(s string) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		res[len(s)-i-1] = s[i]
	}
	return string(res)
}
