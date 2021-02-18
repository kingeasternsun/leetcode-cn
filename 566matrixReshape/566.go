/*
 * @Description:566
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-17 21:57:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-17 22:08:42
 * @FilePath: /566matrixReshape/566.go
 */
package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	rows := len(nums)
	cols := len(nums[0])
	if rows == r && cols == c {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for rowID := 0; rowID < rows; rowID++ {
		for colID := 0; colID < cols; colID++ {
			id := rowID*cols + colID
			res[id/c][id%c] = nums[rowID][colID]
		}
	}

	return res
}
