package main

/*
func main() {
	var c [][]int

	maxtrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(maxtrix))

	b := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(b))

	c = [][]int{
		{1},
		// {5, 6, 7, 8},
		// {9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		{1},
		{5},
		// {9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		{1, 2},
		{5, 6},
		// {9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		{1, 2},
		{5, 6},
		{9, 10},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		// {1},
		{5, 6, 7},
		{9, 10, 11},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		{1},
		{5},
		{9},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		{1},
		{5},
		// {9},
	}
	fmt.Println(spiralOrder(c))

	c = [][]int{
		// {1},
		{1, 2, 3},
		{4, 5, -6},
		{-5, 6, 7},
		{9, 10, 11},
		{12, 13, 14},
	}
	fmt.Println(spiralOrder(c))

}

*/
//https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}

	if row == 1 {
		return matrix[0]
	}

	col := len(matrix[0])
	nums := make([]int, row*col)

	if col == 1 {

		for i := 0; i < row; i++ {
			nums[i] = matrix[i][0]
		}
		return nums
	}

	nRoundEven := false
	nRound := row
	if nRound > col {
		nRound = col

	}
	if nRound%2 == 0 {
		nRoundEven = true
	}
	nRound = (nRound + 1) / 2
	id := 0

	// rowEven := false
	// if row%2 == 0 {
	// 	rowEven = true
	// }
	// colEven := false
	// if col%2 == 0 {
	// 	colEven = true
	// }

	for i := 0; i < nRound; i++ {

		// fmt.Println(i, nRound, row, col)

		if (nRoundEven) || (i < nRound-1) {

			for j := i; j < col-i-1; j++ {
				nums[id] = matrix[i][j]
				id++
			}

			// fmt.Println(nums)

			for j := i; j < row-i-1; j++ {
				nums[id] = matrix[j][col-i-1]
				// fmt.Println(nums)

				id++
			}
			// fmt.Println(nums)

			for j := col - i - 1; j > i; j-- {
				nums[id] = matrix[row-i-1][j]
				id++
			}
			// fmt.Println(nums)

			for j := row - i - 1; j > i; j-- {
				nums[id] = matrix[j][i]
				id++
			}
			// fmt.Println(nums)

			continue
		}

		if (nRoundEven == false) && (row <= col) {

			for j := i; j < col-i; j++ {
				nums[id] = matrix[i][j]
				id++
			}
			continue
		}

		// fmt.Println(nums, i, row)

		if (nRoundEven == false) && (row > col) {

			for j := i; j < row-i; j++ {
				nums[id] = matrix[j][col-i-1]

				id++
			}
		}

		// fmt.Println(nums)

	}

	return nums
}
