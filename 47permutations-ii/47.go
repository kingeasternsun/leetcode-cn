package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode-cn.com/problems/permutations-ii/
47. 全排列 II

*/
func main() {

	nums := []int{2, 2, 1, 1}
	// result := make([][]int, 0)
	result := permuteUnique(nums)
	// TestSlice(result)
	fmt.Println(result)

}

//4ms 3.6MB
func permuteUnique(nums []int) [][]int {
	count := len(nums)
	if count == 0 {
		return nil
	}
	result := make([][]int, 0)
	list := make([]int, 0)
	dfs(nums, count, list, &result)
	// fmt.Println(result)
	return result
}

/*
通过map 判断数值是否已经交换过，剩下的跟46. 全排列 https://leetcode-cn.com/problems/permutations/
*/
func dfs(nums []int, count int, tmpArray []int, result *[][]int) {

	swaped := make(map[int]struct{}, 0) //这个位置的数值是否已经交换过

	if count == 0 {
		tmpArray = append(tmpArray[:0:0], tmpArray...)
		*result = append(*result, tmpArray)
		return
	}

	for i := 0; i < count; i++ {

		if _, ok := swaped[nums[i]]; !ok {
			nums[0], nums[i] = nums[i], nums[0] //技巧点在于这里，把要选取的数和第一个数交换，这样 数组后面的数就是剩下的数字了，可以进行递归
			dfs(nums[1:], count-1, append(tmpArray, nums[0]), result)
			nums[0], nums[i] = nums[i], nums[0] // 这里一定要记得恢复

			swaped[nums[i]] = struct{}{}
		}

	}

	return
}

// 4ms 3.7MB
func permuteUnique1(nums []int) [][]int {
	sLen := len(nums)
	if sLen == 0 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, sLen)
	result = dfs1(nums, sLen, 0, list, result, visited)
	// fmt.Println(result)
	return result
}

/*
去重的关键
- 1. 排序，将相同的放在一起
- 2. 相同的数字，后来的能够加入的前提是前面的相同数据已经加入
 result要返回出来
*/
func dfs1(nums []int, sLen int, pos int, list []int, result [][]int, visited []bool) [][]int {

	if pos == sLen {

		// 重点地方
		tmp := make([]int, len(list))
		copy(tmp, list)
		result = append(result, tmp)
		// fmt.Println("get", result)
		// 必须要返回出去
		return result
	}

	for i := 0; i < sLen; i++ {
		if visited[i] {
			continue
		}

		/*
		 i==0 第一个肯定可以放进去
		 nums[i] != nums[i-1] 第一个重复数字可以放进去，不会和已有的重复
		 nums[i] == nums[i-1] && visited[i-1] ;只有前一个重复的放入了，当前的才可以放入，保证重复的数字的相对位置确定，就不会重复
		*/
		if i == 0 || nums[i] != nums[i-1] || visited[i-1] {
			list = append(list, nums[i])
			visited[i] = true
			// fmt.Println("b", list)
			// fmt.Println("b", result)

			result = dfs1(nums, sLen, pos+1, list, result, visited)
			visited[i] = false
			list = list[:pos]
			// fmt.Println("af", result)
			// fmt.Println("af", list)

			// visited[i] = false
		}
	}

	return result
}
