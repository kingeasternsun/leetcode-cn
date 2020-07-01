package main

/*
func main() {

	// a := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	a := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(a)
	fmt.Println(a)
}
*/
func rotate(matrix [][]int) {
	mLen := len(matrix)
	// tmp := make([]int, mLen)
	var tmp int

	if mLen == 0 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}

	nRound := mLen / 2

	for i := 0; i < nRound; i++ {
		// copy(tmp, matrix[mLen-1-i])

		for j := i; j < mLen-i-1; j++ {

			tmp = matrix[mLen-1-i][j]
			matrix[mLen-1-i][j] = matrix[mLen-j-1][mLen-1-i]
			matrix[mLen-j-1][mLen-1-i] = matrix[i][mLen-j-1]
			matrix[i][mLen-j-1] = matrix[j][i]
			matrix[j][i] = tmp

		}

	}
}
