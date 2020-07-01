package main

func main() {

}

/*
46. 全排列 https://leetcode-cn.com/problems/permutations/
0ms 2.7MB
*/

func permute(nums []int) [][]int {
	count := len(nums)
	if count == 0 {
		return nil
	}

	var tmpArray []int
	result := make([][]int, 0)
	dfs(nums, count, tmpArray, &result)
	return result
}

func dfs(nums []int, count int, tmpArray []int, result *[][]int) {

	if count == 0 {
		tmpArray = append(tmpArray[:0:0], tmpArray...)
		*result = append(*result, tmpArray)
		return
	}

	for i := 0; i < count; i++ {
		nums[0], nums[i] = nums[i], nums[0] //技巧点在于这里，把要选取的数和第一个数交换，这样 数组后面的数就是剩下的数字了，可以进行递归
		dfs(nums[1:], count-1, append(tmpArray, nums[0]), result)
		nums[0], nums[i] = nums[i], nums[0] // 这里一定要记得恢复
	}

	return
}
