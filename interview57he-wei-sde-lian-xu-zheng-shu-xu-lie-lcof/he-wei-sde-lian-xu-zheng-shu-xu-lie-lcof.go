package main

import (
	"log"
)

// (a+b)n/2 = target

func main() {

	log.Println(findContinuousSequence(15))
}

func findContinuousSequence(target int) [][]int {

	if target <= 0 {
		return nil
	}

	var result [][]int
	var ar []int
	var arLen int
	var sum int

	// i := 1
	j := 1
	for j <= target/2+2 {
		if sum < target {
			ar = append(ar, j)
			sum += j
			arLen++

			j++
		} else if sum == target {
			tmp := make([]int, arLen)
			copy(tmp, ar)
			result = append(result, tmp)
			//去掉最左
			sum -= ar[0]
			ar = ar[1:]
			arLen--

		} else {
			//去掉最左
			sum -= ar[0]
			ar = ar[1:]
			arLen--
		}

	}

	return result

}
func findContinuousSequence1(target int) [][]int {

	if target <= 0 {
		return nil
	}

	var result [][]int
	// var ar []int

	// dfs(target, 0, 0, ar, 0, &result)

	// 因为要求至少两个，所以数组的第一个数值，最大不会超过target/2+1
	for i := 1; i <= target/2; i++ {
		ar := getSequence(target, i)
		if len(ar) >= 2 {
			result = append(result, ar)
		}
	}

	return result

}

func getSequence(target, start int) []int {
	var ar []int
	var sum int
	for i := start; i <= target-start; i++ {

		// log.Println(start, i, sum, target)

		ar = append(ar, i)
		sum += i

		if sum > target {
			return nil
		}
		if sum == target {
			return ar
		}
	}

	return nil
}

// target 目标数值，start 表示接下来要插入的数值， ar表示已经插入的,sum表示ar中的和
func dfs(target int, start int, sum int, ar []int, arLen int, result *[][]int) {

	// 肯定 无法满足 返回
	// if start > target {
	// 	return
	// }

	// log.Println(ar)
	if sum > target {
		// log.Println(sum)
		return
	}

	if sum == target {
		newAr := make([]int, arLen)
		// log.Println(ar)
		copy(newAr, ar)
		// *result = append(*result, newAr)
		return
	}

	for i := start + 1; i < target; i++ {
		// log.Println(ar)
		ar = append(ar, i)
		arLen++
		sum += i
		dfs(target, i, sum, ar, arLen, result)
		// log.Println(ar)
		ar = ar[:arLen-1]
		arLen--
		sum -= i
	}

	return

}
