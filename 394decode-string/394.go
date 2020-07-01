package main

import (
	"strconv"
	"strings"
)

/*
1. 利用栈 ，遇到 ] 就进行扩展
2. 扩展后，记得判断前面的是否是字母字符串，是的话进一步合并
3. 往栈里面添加字母字符串的时候，记得判断前面是否是字母字符串，是的话合并到前面的字符串中
*/
// 0ms 2MB
func decodeString1(s string) string {

	stack := make([]string, 0)
	var preStr string
	for i := range s {

		switch s[i] {
		case '[':
			stack = append(stack, preStr)
			preStr = ""
		case ']':
			if preStr != "" {

				sLen := len(stack)
				//上一个也是字符串，合并 merge  // "3[a2[c]d]"  "accdaccdaccd" 对应第二个]
				if sLen > 0 && (stack[sLen-1][0] < '0' || stack[sLen-1][0] > '9') {
					stack[sLen-1] = stack[sLen-1] + preStr
				} else {
					stack = append(stack, preStr)
				}

				preStr = ""
			}

			stack = extend(stack)

		default:

			if s[i] >= '0' && s[i] <= '9' {
				//前面是字母字符串，就先把前面的字符串保存
				if preStr != "" && (preStr[0] < '0' || preStr[0] > '9') {

					sLen := len(stack)
					//上一个也是字符串，合并 merge  // "3[a2[c]d]"  "accdaccdaccd" 对应第二个]
					if sLen > 0 && (stack[sLen-1][0] < '0' || stack[sLen-1][0] > '9') {
						stack[sLen-1] = stack[sLen-1] + preStr
					} else {
						stack = append(stack, preStr)
					}

					preStr = ""
				}
			}

			preStr = preStr + s[i:i+1]

		}

	}

	if preStr != "" {
		stack = append(stack, preStr)
	}

	var result string
	for i := range stack {
		result = result + stack[i]
	}
	return result

}

func extend(stack []string) []string {
	sLen := len(stack)
	str := stack[sLen-1]
	numStr := stack[sLen-2]
	num, _ := strconv.Atoi(numStr)
	extendStr := strings.Repeat(str, num)

	//前面也是字符串 就合并
	if sLen > 2 && (stack[sLen-3][0] < '0' || stack[sLen-3][0] > '9') {
		extendStr = stack[sLen-3] + extendStr
		stack[sLen-3] = extendStr
		stack = stack[:sLen-2]
	} else {
		stack[sLen-2] = extendStr
		stack = stack[:sLen-1]
	}
	return stack

}

/*优雅的写法，
1. 当前是字母字符，如果栈顶的是字母字符串，追加到前面的字母字符串
2. 当前是数字字母，就保存到preStr中，然后等遇到 [ 的时候 入栈
*/
func decodeString(s string) string {

	stack := make([]string, 0)
	var preStr string
	for i := range s {

		switch s[i] {
		case '[':
			stack = append(stack, preStr)
			preStr = ""
		case ']':

			stack = extend(stack)

		default:

			if s[i] >= '0' && s[i] <= '9' {

				preStr = preStr + s[i:i+1]

			} else {

				sLen := len(stack)
				//上一个也是字母字符串，合并 merge  // "3[a2[c]d]"  "accdaccdaccd" 对应第二个]
				if sLen > 0 && (stack[sLen-1][0] < '0' || stack[sLen-1][0] > '9') {
					stack[sLen-1] = stack[sLen-1] + s[i:i+1]
				} else {
					stack = append(stack, s[i:i+1])
				}
			}

		}

	}

	var result string
	for i := range stack {
		result = result + stack[i]
	}
	return result

}
