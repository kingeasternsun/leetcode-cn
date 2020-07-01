package main

import "sort"

//https://leetcode-cn.com/problems/combination-sum/
//dfs 12ms 5.8MB
func combinationSum0(candidates []int, target int) [][]int {

	tmpArray := make([]int, 0)
	result := make([][]int, 0)
	dfs0(0, tmpArray, 0, candidates, target, &result)
	return result
}

func dfs0(curSum int, tmpArray []int, endIndex int, candidates []int, target int, result *[][]int) {

	if curSum == target {
		tmpArray = append(tmpArray[:0:0], tmpArray...)
		*result = append(*result, tmpArray)
		return
	}

	if curSum > target {
		return
	}

	i := endIndex
	for i < len(candidates) {
		dfs0(curSum+candidates[i], append(tmpArray, candidates[i]), i, candidates, target, result)
		i++
	}

	return
}

//dfs 4ms 3.1MB
func combinationSum(candidates []int, target int) [][]int {

	tmpArray := make([]int, 0)
	result := make([][]int, 0)

	sort.Ints(candidates)

	dfs(0, tmpArray, 0, candidates, target, &result)
	return result
}

func dfs(curSum int, tmpArray []int, endIndex int, candidates []int, target int, result *[][]int) {

	if curSum == target {
		tmpArray = append(tmpArray[:0:0], tmpArray...)
		*result = append(*result, tmpArray)
		return
	}

	// if curSum > target {
	// 	return
	// }

	i := endIndex
	for i < len(candidates) {

		if curSum+candidates[i] > target {
			return
		}

		dfs(curSum+candidates[i], append(tmpArray, candidates[i]), i, candidates, target, result)
		i++
	}

	return
}
