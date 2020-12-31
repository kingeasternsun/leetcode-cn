package main

import "fmt"

func main() {

	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
	fmt.Println(countRangeSum([]int{1, 3, 2, 3, 1, 2, -1, -2}, -2, 2))
	fmt.Println(countRangeSum([]int{1, 3, 2, 3}, -2, 2))
}

// [1, 3, 2, 3, 1,2,-1,-2] 预期10
// -2
// 2

func countRangeSum(nums []int, lower int, upper int) int {

	sums := make([]int, len(nums))
	// tmpSum := 0
	for i, v := range nums {
		if i == 0 {
			sums[i] = v
		} else {
			sums[i] = sums[i-1] + v
		}
	}

	return help(sums, lower, upper)
}

// -2 3 2
func help(preSums []int, lower int, upper int) int {

	cnt := len(preSums)
	if cnt == 0 {
		return 0
	}

	if cnt == 1 {

		if lower <= preSums[0] && preSums[0] <= upper {
			return 1
		}

		return 0
	}

	newSums := make([]int, cnt)
	copy(newSums, preSums)
	left := newSums[0 : cnt/2]
	leftNum := help(left, lower, upper)

	right := newSums[cnt/2:]
	rightNum := help(right, lower, upper)

	num := leftNum + rightNum
	j := 0
	k := 0
	for _, v := range left {

		// right[j]- v >= lower ,找到 第一个  right[j] >= lower + v的
		for j < len(right) && right[j] < lower+v {
			j++
		}

		if k < j {
			k = j
		}
		// right[k] - v < = upper ,找到第一个 right[k] > upper + v 的
		for k < len(right) && right[k] <= upper+v {
			k++
		}

		num += (k - j)

	}

	//排序
	i := 0
	j = 0
	k = 0

	//i和j不能同时到达2右边边界，但是有一个到达右边边界是可以的
	for i < len(left) || j < len(right) {

		if i == len(left) {
			preSums[k] = right[j]
			j++

		} else if j == len(right) {
			preSums[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			preSums[k] = left[i]
			i++
		} else {
			preSums[k] = right[j]
			j++
		}

		k++
	}

	return num

}
