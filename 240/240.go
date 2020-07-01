package main

// 240. 搜索二维矩阵 II
import "fmt"

// 另外一个思路 从右上角开始查询，target更大，往下走，target更小，往左走
/*
func main() {
	// fmt.Println(searchSlice([]int{1}, 1))
	// fmt.Println(searchSlice([]int{1}, 2))
	// fmt.Println(searchSlice([]int{1, 2}, 1))
	// fmt.Println(searchSlice([]int{1, 2}, 3))
	// fmt.Println(searchSlice([]int{1, 2}, 0))

	var m [][]int
	m=[][]int{
		{1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println(searchMatrix(m,5))

}
*/
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

func searchMatrix(matrix [][]int, target int) bool {

	// if len(matrix) == 0 {
	// 	return false
	// }

	// return search(matrix,Point{0,0},Point{row:len(matrix)-1,col:len(matrix[0])-1},target)
	return searchMatrix2(matrix, target)

}

func searchMatrix2(matrix [][]int, target int) bool {
	maxRow := len(matrix)
	if maxRow == 0 {
		return false
	}
	maxCol := len(matrix[0])

	if maxCol == 0 {
		return false
	}

	row := 0
	col := maxCol - 1

	for {

		fmt.Printf("%v %v\n", row, col)
		if matrix[row][col] == target {
			return true
		}

		if target > matrix[row][col] {
			if row == maxRow-1 {
				return false
			}

			row = row + 1
			continue
		}

		if col == 0 {
			return false
		}

		col = col - 1

	}
}

type Point struct {
	row, col int
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
