/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-12 21:24:14
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-12 22:00:36
 * @FilePath: /179largestNumber/179.go
 */
package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	numStr := make([]string, len(nums))
	allZero := true

	for i, n := range nums {
		numStr[i] = strconv.Itoa(n)
		if n > 0 {
			allZero = false
		}
	}

	if allZero {
		return "0"
	}

	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] > numStr[j]+numStr[i]
	})
	res := strings.Builder{}
	for _, str := range numStr {
		res.WriteString(str)
	}

	return res.String()
}
