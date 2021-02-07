/*
 * @Description:344
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 10:01:14
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 10:01:49
 * @FilePath: \344reverseString\344.go
 */

package leetcode

func reverseString(word []byte) {
	if len(word) <= 1 {
		return
	}

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return
}
