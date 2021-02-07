/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 10:22:18
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 10:50:01
 * @FilePath: \13romanToInt\13.go
 */
package leetcode

/*

思路：从右往左遍历，当前的字符大于等于上一个字符，则累加； 否则累减

代码
*/
func romanToInt(s string) int {
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10。

	// 当右边的字符比左边的大， 说明是特殊组合
	pre, r := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := _map[s[i]]
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}
	return r
}
