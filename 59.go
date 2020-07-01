package main

// func main() {

// 	fmt.Println(generateMatrix(1))
// }
func generateMatrix(n int) [][]int {

	if n <= 0 {
		return nil
	}

	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	// sLen := n * n
	// for i := 1; i < sLen+1; i++ {

	// }

	num := 1

	nRound := n / 2
	if nRound*2 != n {
		nRound = nRound + 1
	}
	for nr := 0; nr < nRound; nr++ {
		tmpLen := n - 2*nr - 1
		if tmpLen == 0 { //奇数，圈只有一个
			result[nr][nr] = num
			continue
		}

		if tmpLen < 0 {
			break
		}

		// fmt.Println(n, nr, tmpLen)
		for i := 0; i < tmpLen; i++ {
			result[nr][nr+i] = num
			num++
		}

		for i := 0; i < tmpLen; i++ {
			result[nr+i][n-nr-1] = num
			num++
		}

		for i := tmpLen; i > 0; i-- {
			result[n-nr-1][nr+i] = num
			num++
		}

		for i := tmpLen; i > 0; i-- {
			result[nr+i][nr] = num
			num++
		}
	}
	return result
}

// 1, 2,  3, 4
// 12,13,14, 5
// 11,16,15, 6
// 10, 9, 8, 7
