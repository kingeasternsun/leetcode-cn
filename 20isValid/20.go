/*
 * @Description: 20
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:33:49
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:39:34
 * @FilePath: \20isValid\20.go
 */
package leetcode

func isValid(s string) bool {

	if len(s)&1 == 1 {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if s[i] == '}' && stack[len(stack)-1] != '{' {
			return false
		}
		if s[i] == ']' && stack[len(stack)-1] != '[' {
			return false
		}
		if s[i] == ')' && stack[len(stack)-1] != '(' {
			return false
		}
		stack = stack[:len(stack)-1]

	}

	if len(stack) == 0 {
		return true
	}

	return false

}
