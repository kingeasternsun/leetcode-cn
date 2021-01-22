package leetcode

import "sort"

// https://leetcode-cn.com/problems/maximal-rectangle/

// https://leetcode-cn.com/problems/maximal-square/

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

func largestSubmatrix1(matrix [][]int) int {
	/*
		1.对每一行，计算当前位置向下的连续列长，计算最大的子矩阵
		2. 列长排序后，就容易以当前列长为边的最大子矩阵

	*/

	rows := len(matrix)
	if rows == 0 {
		return 0
	}

	res := 0
	cols := len(matrix[0])

	preHis := make([]int, cols) //记录上一行，每个点往下的连续列长

	for rowID := 0; rowID < rows; rowID++ {
		curHis := make([]int, 0)
		for colID := 0; colID < cols; colID++ {
			//计算 matrix[rowID][colID]开始向下的连续1的长度
			if preHis[colID] > 1 { //如果上面已经计算好 连续长度了，就减去一就可以了
				preHis[colID]--
				curHis = append(curHis, preHis[colID])
				continue
			} else if matrix[rowID][colID] == 0 {
				preHis[colID] = 0 //注意这里要清0
				continue

			}

			curLen := 1
			//往下找1
			for j := rowID + 1; j < rows; j++ {
				if matrix[j][colID] == 0 {
					break
				}
				curLen++
			}

			preHis[colID] = curLen
			curHis = append(curHis, preHis[colID])

		}

		//对curHis排序,然后计算以当前列长为边的最大子矩阵
		sort.Ints(curHis)
		// tmpMax := 0
		for i, h := range curHis {
			res = max(res, h*(len(curHis)-i))
		}

	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

//对每一行，计算当前位置向上的连续列长，计算最大的子矩阵 ,比上面的方法更加优雅些
func largestSubmatrix(matrix [][]int) int {
	/*
		1. 对每一行，计算当前位置向上的连续列长，计算最大的子矩阵
		2. 列长排序后，就容易以当前列长为边的最大子矩阵

	*/

	rows := len(matrix)
	if rows == 0 {
		return 0
	}

	res := 0
	cols := len(matrix[0])

	preHis := make([]int, cols) //记录上一行，每个点往上的连续列长

	for rowID := 0; rowID < rows; rowID++ {
		curHis := make([]int, 0)
		for colID := 0; colID < cols; colID++ {
			//计算 matrix[rowID][colID]开始向上的连续1的长度
			if matrix[rowID][colID] == 0 {
				preHis[colID] = 0 //注意这里要清0
			} else {
				preHis[colID]++
			}

			curHis = append(curHis, preHis[colID])

		}

		//对curHis排序,然后计算以当前列长为边的最大子矩阵
		sort.Ints(curHis)
		// tmpMax := 0
		for i, h := range curHis {
			res = max(res, h*(len(curHis)-i))
		}

	}

	return res
}
