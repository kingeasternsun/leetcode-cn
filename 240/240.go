package leetcode

// 240. 搜索二维矩阵 II
import (
	"fmt"
)

//从右上角开始查询，target更大，往下走，target更小，往左走
func searchMatrix(matrix [][]int, target int) bool {

	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	x := 0
	y := cols - 1
	for y >= 0 && x < rows {

		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target {
			y-- //左移动
		} else {
			x++ //下走
		}
	}

	return false

}

type Point struct {
	row, col int
}

func searchSlice(a []int, target int) bool {

	if len(a) == 0 {
		return false
	}

	var beg, mid, end int

	end = len(a) - 1
	for {

		if beg > end {
			return false
		}

		mid = (beg + end) / 2
		if a[mid] == target {
			return true
		}

		if target > a[mid] {
			beg = mid + 1

			if beg >= len(a) {
				return false
			}

		} else {
			end = mid - 1

			if end < 0 {
				return false
			}
		}
	}

}
func search(matrix [][]int, beg, end Point, target int) bool {

	fmt.Printf("%+v %+v\n", beg, end)

	mLen := len(matrix)
	mcLen := len(matrix[0])

	if beg.col > end.col {
		return false
	}
	if beg.row > end.row {
		return false
	}

	if beg.row >= mLen || beg.row < 0 {
		return false
	}

	if end.row >= mLen || end.row < 0 {
		return false
	}

	if beg.col >= mcLen || beg.col < 0 {
		return false
	}

	if end.col >= mcLen || end.col < 0 {
		return false
	}

	if matrix[beg.row][beg.col] == target {
		return true
	}

	mid := Point{
		row: (beg.row + end.row) / 2,
		col: (beg.col + end.col) / 2,
	}

	midValue := matrix[mid.row][mid.col]
	if midValue == target {
		return true
	}
	fmt.Printf("%v %v \n", mid, midValue)

	// 如果中间的数值比目标小
	if midValue > target {
		//第二象限
		if search(matrix, Point{row: mid.row + 1, col: mid.col + 1}, end, target) {
			return true
		}

		//第一象限
		if search(matrix, Point{row: beg.row, col: mid.col + 1}, Point{row: mid.row, col: end.col}, target) {
			return true
		}

		//第三象限
		if search(matrix, Point{row: mid.row + 1, col: beg.col}, Point{row: end.row, col: mid.col}, target) {
			return true
		}

	}

	//第一象限
	if search(matrix, Point{row: beg.row, col: mid.col}, Point{row: mid.row - 1, col: end.col}, target) {
		return true
	}
	//第四象限
	if search(matrix, beg, Point{row: mid.row - 1, col: mid.col - 1}, target) {
		return true
	}

	//第三象限
	if search(matrix, Point{row: mid.row, col: beg.col}, Point{row: end.row, col: mid.col - 1}, target) {
		return true
	}

	return false
}
