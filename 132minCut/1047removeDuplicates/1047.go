/*
 * @Description: 1047
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-09 09:13:18
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-09 10:03:16
 * @FilePath: \1047removeDuplicates\1047.go
 */
package letcode

// 利用栈保存新生成的字符串，然后消消乐
func removeDuplicates(S string) string {

	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {

		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}

	return string(stack)

}
