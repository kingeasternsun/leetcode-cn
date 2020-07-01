package main

//https://leetcode-cn.com/problems/maximal-square/
/*
1. 计算每一行中每个元素往上的最长连续竖线长度
2. 然后计算可以构成的最大正方形，代码中只需要记录正方形边长就可以
*/
func maximalSquare(matrix [][]byte) int {

	rows := len(matrix)
	if rows == 0 {
		return 0
	}

	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	maxV := 0

	//以当前位置的连续竖线的高度
	hMaxtrix := make([][]int, 2)
	hMaxtrix[0] = make([]int, cols)
	hMaxtrix[1] = make([]int, cols)
	for i := 0; i < rows; i++ {
		hMaxtrix[i&1] = make([]int, cols)
		for j := range hMaxtrix[i&1] {
			if matrix[i][j] == '0' {
				hMaxtrix[i&1][j] = 0
				continue
			}

			//考虑上一层 consider the pre row
			hMaxtrix[i&1][j] = hMaxtrix[(i+1)&1][j] + 1
		}

		tmpV := help2(hMaxtrix[i&1], rows-i, cols)
		if tmpV > maxV {
			maxV = tmpV
		}
	}

	return maxV

}

// 0ms 5.2MB
func help(heighs []int, rows, cols int) int {

	maxH := 0 //最大正方形的边长

	dp := make([]int, cols) //以当前竖线作为右边的最大正方形的边长
	for i := range heighs {

		if heighs[i] == 0 {
			dp[i] = 0
			continue
		}

		//竖线不为0，
		//第一个
		if i == 0 {
			dp[i] = 1
			continue
		}

		//比左边所能构成的正方形边还要短或一样，当前竖线的高度决定了正方形的边长
		// 5 5 3 3 2 对应的dp 1 2 3 3 2
		if heighs[i] <= dp[i-1] {
			dp[i] = heighs[i]
			continue
		}

		canAddOne := true
		//这里重要  dp[i]比dp[i-1]最多只会多1
		//看dp[i-1]对应的正方形的上方有没有多一行
		for j := i - 1; j >= i-dp[i-1]; j-- {
			if heighs[j] == dp[i-1] { //这里高度刚好才到 dp[i-1]
				canAddOne = false
				break
			}
		}

		if canAddOne {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}

	}

	for i := range dp {
		if dp[i] > maxH {
			maxH = dp[i]
		}
	}

	return maxH * maxH
}

// 0ms 4.3MB
func help2(heighs []int, rows, cols int) int {

	maxH := 0 //最大正方形的边长

	preDP := 0 //以当前竖线作为右边的最大正方形的边长
	for i := range heighs {

		if preDP > maxH {
			maxH = preDP
		}

		//如果竖线为0
		if heighs[i] == 0 {
			preDP = 0
			continue
		}

		//竖线不为0，
		//第一个
		if i == 0 {
			preDP = 1
			continue
		}

		//比左边所能构成的正方形边还要短或一样，当前竖线的高度决定了正方形的边长
		// 5 5 3 3 2 对应的dp 1 2 3 3 2
		if heighs[i] <= preDP {
			preDP = heighs[i]
			continue
		}

		canAddOne := true
		//这里重要  dp[i]比dp[i-1]最多只会多1
		//看dp[i-1]对应的正方形的上方有没有多一行
		for j := i - 1; j >= i-preDP; j-- {
			if heighs[j] == preDP { //这里高度刚好才到 dp[i-1] ，就无法再加一了
				canAddOne = false
				break
			}
		}

		if canAddOne {
			preDP = preDP + 1
		}

	}

	if preDP > maxH {
		maxH = preDP
	}

	return maxH * maxH
}
