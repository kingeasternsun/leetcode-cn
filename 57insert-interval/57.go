package main

import "sort"

//https://leetcode-cn.com/problems/insert-interval/

// 8ms 4.7MB use extra memory
func insert(intervals [][]int, newInterval []int) [][]int {

	count := len(intervals)
	if count == 0 {
		return [][]int{newInterval}
	}

	// put in the end
	if newInterval[0] >= intervals[count-1][0] {
		intervals = append(intervals, newInterval)
		return merge0(intervals)

	}

	newIntervals := make([][]int, 0)
	newCount := 0
	i := 0
	hasAddNewInterval := false
	for i < count {

		tmp := intervals[i]

		if (!hasAddNewInterval) && newInterval[0] < intervals[i][0] {
			tmp = newInterval
			i--
			hasAddNewInterval = true
		}

		if newCount == 0 {
			newIntervals = append(newIntervals, tmp)
			i++
			newCount++
			continue
		}

		if tmp[0] > newIntervals[newCount-1][1] {
			newIntervals = append(newIntervals, tmp)
			newCount++
			i++
			continue
		}

		// join interval
		// only large : update the last interval int new result
		if tmp[1] > newIntervals[newCount-1][1] { //!!!! attention
			newIntervals[newCount-1][1] = tmp[1]
		}

		i++

	}

	return newIntervals

}

func merge0(intervals [][]int) [][]int {

	count := len(intervals)
	if count <= 1 {
		return intervals
	}

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

// use the code 56//https://leetcode-cn.com/problems/merge-intervals/ 16ms 4.7MB
func insert1(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)

}

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
