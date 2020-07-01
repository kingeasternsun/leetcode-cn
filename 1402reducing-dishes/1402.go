package main

import "sort"

/*
贪心算法 0ms 2.3MB
*/
func maxSatisfaction(satisfaction []int) int {

	count := len(satisfaction)
	if count == 0 {
		return 0
	}
	sort.Ints(satisfaction)
	singleSum := 0 //各个菜不考虑时间的单个满意度的当前和
	sf := 0        //考虑时间的满意度
	for i := count - 1; i >= 0; i-- {
		singleSum = satisfaction[i] + singleSum
		if singleSum <= 0 {
			break
		}

		//各个菜集体往后延时一个时间 ，那么满意度就会增加  singleSum
		sf += singleSum
	}

	return sf

}
