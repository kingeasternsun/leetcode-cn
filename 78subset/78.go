package main

import "fmt"

func main() {

	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{}))
	fmt.Println(subsets([]int{1}))
}

func subsets(nums []int) [][]int {

	tmpRes := make([]int, 0)
	res := make([][]int, 0)
	bfs(nums, 0, tmpRes, &res)
	return res

}

func bfs(nums []int, id int, tmpRes []int, res *[][]int) {

	if id == len(nums) {
		tmp := make([]int, len(tmpRes))
		copy(tmp, tmpRes)
		*res = append(*res, tmp)
		return
	}
	bfs(nums, id+1, tmpRes, res)

	bfs(nums, id+1, append(tmpRes, nums[id]), res)

}
