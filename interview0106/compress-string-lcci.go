/*
 * @Author: your name
 * @Date: 2020-03-16 08:55:56
 * @LastEditTime: 2020-03-16 09:10:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\interview0106\compress-string-lcci.go
 */
package main

import "strconv"

func main() {

}

func compressString(S string) string {

	// str := []byte("0123456789")
	count := len(S)
	if count <= 2 {
		return S
	}
	pre := S[0]
	preCount := 1
	newS := make([]byte, 0)
	newSLen := 0
	for i := 1; i < count; i++ {
		if S[i] != pre {
			newS = append(newS, pre)
			newS = append(newS, []byte(strconv.Itoa(preCount))...)
			pre = S[i]
			preCount = 1
			newSLen = newSLen + 2
			if newSLen >= count {
				return S
			}
		} else {
			preCount++
		}
	}
	newS = append(newS, pre)
	newS = append(newS, []byte(strconv.Itoa(preCount))...)
	newSLen = newSLen + 2
	if newSLen >= count {
		return S
	}

	return string(newS)

}
