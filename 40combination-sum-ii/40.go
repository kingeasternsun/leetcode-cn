package main

import "sort"

//https://leetcode-cn.com/problems/combination-sum-ii/ dfs
func combinationSum2(candidates []int, target int) [][]int {
	tmpArray := make([]int, 0)
	selected := make([]bool, len(candidates))
	result := make([][]int, 0)
	sort.Ints(candidates)
	dfs1(tmpArray, selected, 0, -1, candidates, target, &result)
	return result
}

// selected[i] = true 关键地方 ，标记第几个是否选择了， 重复的只有在前面的被选择了当前才可以选择
// 4ms 2.6MB
func dfs1(curArray []int, selected []bool, curSum int, endIndex int, candidates []int, target int, result *[][]int) {

	if curSum == target {
		curArray = append(curArray[:0:0], curArray...)
		*result = append(*result, curArray)
		return
	}

	i := endIndex + 1
	for i < len(candidates) {

		if curSum+candidates[i] > target {
			return
		}

		//  for example  [1,1,2,3]  ,target = 6
		//  the combination of the first `1` must contail the 2nd `1`
		if i == 0 || candidates[i] != candidates[i-1] || selected[i-1] {
			selected[i] = true
			dfs1(append(curArray, candidates[i]), selected, curSum+candidates[i], i, candidates, target, result)
		}
		selected[i] = false
		i++
	}

	return
}

/*
 selected[i] = true 关键地方 ，标记第几个是否选择了， 重复的只有在前面的被选择了当前才可以选择
 8ms 4.8MB
*/
func dfs(curArray []int, selected []bool, curSum int, endIndex int, candidates []int, target int, result *[][]int) {

	if curSum == target {
		curArray = append(curArray[:0:0], curArray...)
		*result = append(*result, curArray)
		return
	}

	if curSum > target {
		return
	}

	i := endIndex + 1
	if i >= len(candidates) {
		return
	}

	//选择添加当前数值 只有非重复 或者 有重复 但是重复的数值已经加过了才可以加
	if i == 0 || candidates[i] != candidates[i-1] || selected[i-1] {
		selected[i] = true
		dfs(append(curArray, candidates[i]), selected, curSum+candidates[i], i, candidates, target, result)
	}

	// 不选择当前选项
	selected[i] = false // 记得这里要标注不选
	dfs(curArray, selected, curSum, i, candidates, target, result)

	return
}
