package main

/*
378. 有序矩阵中第K小的元素 https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/\
 1.2,3,4,3.2.1
 斜着一行一行的去看，后面一行比前面的一行要大

 log(n)的算法 需要好好再看看
 https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/85170/O(n)-from-paper.-Yes-O(rows).

 https://www.cnblogs.com/grandyang/p/5727892.html

*/

func kthSmallest(matrix [][]int, k int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	if k == 0 {
		return matrix[0][0]
	}

	if k >= rows*rows {
		return matrix[rows-1][rows-1]
	}

	begValue := matrix[0][0]
	endValue := matrix[rows-1][rows-1]

	for begValue < endValue {

		midValue := (begValue + endValue) / 2
		cnt := count(matrix, midValue, rows, cols)
		// if cnt == k {
		// 	return begValue
		// }
		if cnt < k {
			begValue = midValue + 1
		} else { //cnt >=k 那就说明 mivValue 有可能包含 k 个，所以 endValue = midValue
			endValue = midValue
		}

	}

	return begValue

}

//查询小于等于 midValue的个数
func count(matrix [][]int, midValue, rows, cols int) int {
	//从左下角开始
	i := rows - 1
	j := 0

	cnt := 0 //记录当前矩阵中有多少个小于等于 midValue
	for i >= 0 && j < cols {
		if matrix[i][j] <= midValue { //当前值小于等于MidValue，那么当前值所在列中当前值上面的数值肯定都小于midValue， 然后向右移动
			cnt = cnt + i + 1

			//然后向右移动
			j++
		} else { //当前值大于 MidValue 的数字，向上移动,变小
			i--
		}
	}
	return cnt
}
