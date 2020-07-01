// 118. 杨辉三角 https://leetcode-cn.com/problems/pascals-triangle/
package main

func generate(numRows int) [][]int {

	if numRows == 0 {
		return nil
	}

	result := make([][]int, 1)
	result[0] = []int{1}

	if numRows == 1 {
		return result
	}

	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1

		j := 1
		for j < i {
			tmp[j] = result[i-1][j-1] + result[i-1][j]
			j++
		}

		result = append(result, tmp)

	}

	return result

}
