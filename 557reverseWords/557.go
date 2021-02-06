/*
 * @Description: 557
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 16:14:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 09:19:06
 * @FilePath: \557reverseWords\557.go
 */
package leetcode

func reverseWords(s string) string {

	in := []byte(s)
	start, end := 0, 0 //左开右闭
	for end <= len(in) {
		//遇到空格或边界
		if end == len(in) || in[end] == ' ' {
			reverse(in[start:end])
			end = end + 1
			start = end
		} else {
			end++
		}

	}

	return string(in)
}

func reverse(word []byte) {
	if len(word) <= 1 {
		return
	}
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return
}
