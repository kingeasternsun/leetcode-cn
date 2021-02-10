/*
 * @Author: kingeasternsun
 * @Date: 2021-02-08 11:05:29
 * @LastEditTime: 2021-02-08 15:06:03
 * @LastEditors: Please set LastEditors
 * @Description: leetcode
 * @FilePath: /978maxTurbulenceSize/978.go
 */
package leetcode

//其实就是寻找最大波浪长度
func maxTurbulenceSize(arr []int) int {
	if len(arr) <= 1 {
		return 1
	}

	// 0表示 arr[i-1]==arr[i] 相同，
	// 1表示 arr[i-1]<arr[i],
	// 2表示 arr[i-1]>arr[i]
	preStatus, curStatus := 0, 0
	curLen, maxLen := 0, 0
	for i := 1; i < len(arr); i++ {
		preStatus = curStatus

		switch {
		case arr[i] == arr[i-1]:
			curStatus = 0
		case arr[i-1] < arr[i]:
			curStatus = 1
		case arr[i-1] > arr[i]:
			curStatus = 2
		}

		// 拐点
		if preStatus+curStatus == 3 {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			if curStatus == 0 {
				curLen = 1
			} else {
				curLen = 2
			}
		}

	}

	if curLen > maxLen {
		maxLen = curLen
	}

	return maxLen

}

func cmp(i, j int) int {
	if i == j {
		return 0
	}
	if i > j {
		return 1
	}
	return -1
}

func maxTurbulenceSize2(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	start := 0
	res := 1

	for i := 0; i < len(arr)-1; i++ {
		if cmp(arr[i], arr[i+1]) == 0 {
			res = max(res, i-start+1)
			start = i + 1
		} else if i > 0 && cmp(arr[i], arr[i+1])*cmp(arr[i-1], arr[i]) != -1 {
			start = i
		} else {
			res = max(res, i-start+2)
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
