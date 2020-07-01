package main

// https://leetcode-cn.com/problems/search-a-2d-matrix/submissions/

func biSearch(arr []int, target int) bool {

	beg := 0
	end := len(arr) - 1

	mid := 0
	for {
		if beg > end {
			return false
		}

		// if beg == end {
		// 	if arr[beg] == target {
		// 		return true
		// 	}
		// }

		mid = (beg + end) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			end = mid - 1
		} else {
			beg = mid + 1
		}

	}

	// return false
}

func searchMatrix1(matrix [][]int, target int) bool {

	row := len(matrix)
	if row == 0 {
		return false
	}

	col := len(matrix[0])
	if col == 0 {
		return false
	}

	beg := 0
	end := row - 1
	var mid int

	for {

		if beg > end || end < 0 || beg >= row {
			return false
		}

		mid = (beg + end) / 2

		if matrix[mid][0] <= target && target <= matrix[mid][col-1] {
			return biSearch(matrix[mid], target)
		}

		if matrix[mid][0] > target {
			end = mid - 1
		} else {
			beg = mid + 1
		}

	}

	return false
}

// 二维数组虚拟为一维数组
func searchMatrix(matrix [][]int, target int) bool {

	row := len(matrix)
	if row == 0 {
		return false
	}

	col := len(matrix[0])
	if col == 0 {
		return false
	}

	beg := 0
	end := row*col - 1

	mid := 0
	for {
		if beg > end {
			return false
		}

		// if beg == end {
		// 	if arr[beg] == target {
		// 		return true
		// 	}
		// }

		mid = (beg + end) / 2
		midV := matrix[mid/col][mid%col]
		if midV == target {
			return true
		}
		if midV > target {
			end = mid - 1
		} else {
			beg = mid + 1
		}

	}

	return false
}
