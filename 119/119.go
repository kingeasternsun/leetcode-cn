//119. 杨辉三角 II https://leetcode-cn.com/problems/pascals-triangle-ii/

package main

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	result := make([][]int, 2)
	result[0] = make([]int, rowIndex+1)
	result[1] = make([]int, rowIndex+1)

	for i := 1; i <= rowIndex; i++ {
		tmp := result[i&1]
		tmp[0] = 1
		tmp[i] = 1

		j := 1
		for j < i {
			tmp[j] = result[(i-1)&1][j-1] + result[(i-1)&1][j]
			j++
		}

	}

	return result[rowIndex&1]

}
