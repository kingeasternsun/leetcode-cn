/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-02 10:12:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-03 10:26:57
 * @FilePath: \1004longestOnes\1004.go
 */
package leetcode

/*
 相似题目 424,1493,1004.
 替换后的最长重复字符
 利用二分法加滑窗
*/
func longestOnesBinary(A []int, K int) int {
	if K >= len(A) {
		return len(A)
	}

	start := K
	end := len(A)
	best := -1
	for start <= end {

		mid := (end-start)/2 + start
		if judege(A, mid, K) {
			best, start = mid, mid+1
		} else {
			end = mid - 1
		}

	}
	return best
}

func judege(A []int, winLen, K int) bool {
	numberCnt := [2]int{} //记录0，1的个数
	for i := 0; i < winLen; i++ {
		numberCnt[A[i]]++
	}

	if numberCnt[1]+K >= winLen {
		return true
	}

	for i := winLen; i < len(A); i++ {
		numberCnt[A[i-winLen]]--
		numberCnt[A[i]]++
		if numberCnt[1]+K >= winLen {
			return true
		}
	}

	return false

}

//双指针方法
func longestOnes(A []int, K int) int {
	if K >= len(A) {
		return len(A)
	}
	if K == 0 {
		return findMaxConsecutiveOnes(A)
	}

	numCnt := [2]int{}
	start, end := 0, 0 //左闭，右开区间，包含start,不包含end
	maxLen := 0
	for start <= end && end < len(A) {
		//要加的end是1，可以继续加 或 要加入的是0， 而且目前窗口中0 的数量还小于 K，可以继续加 0
		if A[end] == 1 || numCnt[0] < K {
			numCnt[A[end]]++
			end++
			if end-start > maxLen {
				maxLen = end - start
			}
			continue
		}

		//当前窗口中0的数量等于k,就不能再往里面加0了，就需要把窗口左边缩小
		numCnt[A[start]]--
		start++

	}

	return maxLen

}

func findMaxConsecutiveOnes(nums []int) int {
	maxLen := 0
	curLen := 0
	for _, v := range nums {

		if v == 1 {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 0
		}
	}

	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}
