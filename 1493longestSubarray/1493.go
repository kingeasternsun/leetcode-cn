/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-02 11:50:42
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-03 10:23:31
 * @FilePath: \1493longestSubarray\1493.go
 */
package leetcode

/*
 相似题目 424,1493,1004.
二分法
https://github.com/kingeasternsun/leetcode-cn
*/
func longestSubarrayBinary(nums []int) int {

	start := 0
	end := len(nums)
	best := -1
	for start <= end {

		mid := (end-start)/2 + start
		if judege(nums, mid) {
			best, start = mid, mid+1
		} else {
			end = mid - 1
		}

	}
	if best > 0 {
		return best - 1
	}
	return best

}

func judege(nums []int, winLen int) bool {
	numCnt := [2]int{}
	for i := 0; i < winLen; i++ {
		numCnt[nums[i]]++
	}

	if numCnt[0] <= 1 {
		return true
	}

	for i := winLen; i < len(nums); i++ {
		numCnt[nums[i-winLen]]--
		numCnt[nums[i]]++
		if numCnt[0] <= 1 {
			return true
		}

	}

	return false
}

func longestSubarray(nums []int) int {

	maxLen := 0
	numCnt := [2]int{}
	beg, end := 0, 0 //左闭右开区间
	for beg <= end && end < len(nums) {
		//如果当前end上是1 或 [beg,end)里面没有0 ，就可以把 当前end加入窗口
		if nums[end] == 1 || numCnt[0] == 0 {
			numCnt[nums[end]]++
			end++
			if end-beg-1 > maxLen {
				maxLen = end - beg - 1
			}
			continue
		}
		//end上是0 ,而且 [beg,end)里面已经有一个0了，就需要去掉左边的beg
		numCnt[nums[beg]]--
		beg++

	}

	return maxLen
}
