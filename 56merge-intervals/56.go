package main

import "sort"

//https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {

	count := len(intervals)
	if count <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// has first interval
	newCount := 1

	for i := 1; i < count; i++ {
		if intervals[i][0] > intervals[newCount-1][1] {

			if newCount != i {
				copy(intervals[newCount], intervals[i])
			}
			newCount++
			continue
		}

		// join interval:  update the last interval int new result
		if intervals[i][1] > intervals[newCount-1][1] {
			intervals[newCount-1][1] = intervals[i][1]
		}
	}

	return intervals[:newCount]

}
