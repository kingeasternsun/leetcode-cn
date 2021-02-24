/*
 * @Description: 1011
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 10:24:31
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 10:53:33
 * @FilePath: \1011shipWithinDays\1011.go
 */
package leetcode

func shipWithinDays(weights []int, D int) int {

	beg := 1 //这个最小值必须是 weights中的最大值，不然这个货物有可能永远都无法上运输带
	end := 0
	best := 0
	for _, v := range weights {
		end += v
		if v > beg {
			beg = v
		}
	}

	for beg <= end {
		mid := (end-beg)/2 + beg
		if match(weights, D, mid) {
			best, end = mid, mid-1
		} else {
			beg = mid + 1
		}
	}

	return best

}

// mid 每日运输多少个
func match(weights []int, D, mid int) bool {

	days := 1
	curWight := 0
	for _, v := range weights {
		if curWight+v > mid { //当天装不下了,放到下一天装
			days++
			curWight = v
			if days > D {
				return false
			}
		} else {
			curWight += v
		}
	}

	return true
}
