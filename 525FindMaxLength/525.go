package main

import "fmt"

func main() {

	fmt.Println(findMaxLength([]int{0, 1}) == 2)
	fmt.Println(findMaxLength([]int{0, 1, 0}) == 2)
	fmt.Println(findMaxLength([]int{0, 1, 1, 0, 0}) == 4)
	fmt.Println(findMaxLength([]int{1, 1, 1, 1, 1, 1, 1, 1}) == 0)
	fmt.Println(findMaxLength([]int{1, 1, 1, 1, 1, 1, 1}) == 0)
}

//先把里面的0变为-1，变为子区间的和为0 ，也就是两个子序列的和相同
func findMaxLength(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	sum := 0
	begMap := make(map[int]int, 0)
	endMap := make(map[int]int, 0)
	begMap[0] = -1

	for i, v := range nums {
		sum += v
		if _, ok := begMap[sum]; !ok {
			begMap[sum] = i
		}

		endMap[sum] = i
	}

	maxV := 0
	for k, v := range begMap {

		if endv, ok := endMap[k]; ok {
			if endv == v {
				continue
			}

			if endv-v > maxV {
				maxV = endv - v
			}
		}

	}

	return maxV
}
