package main

func setZeroes(matrix [][]int) {

	row := len(matrix)
	if row == 0 {
		return
	}

	col := len(matrix[0])
	if col == 0 {
		return
	}

	//标记第0行 第0列中是否包含0
	var rowZero, colZero bool

	//将第0行，第0列作为map，标记对应的列或行是否包含0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					rowZero = true
				}

				if j == 0 {
					colZero = true
				}

				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
			// }
		}
	}

	if rowZero {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}

	if colZero {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}

}
