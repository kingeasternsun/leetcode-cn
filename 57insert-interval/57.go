package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/insert-interval/

// 8ms 4.7MB use extra memory
func insert0(intervals [][]int, newInterval []int) [][]int {

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
	return merge1(intervals)

}

func merge1(intervals [][]int) [][]int {

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

//4ms 4.8MB
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	//先处理边界情况
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	// resID := 0
	i := 0

	//插入 开始位置在 在newInterval开始位置前面或相同的 都放入res中
	for i < len(intervals) {
		// 找到了插入的位置
		if newInterval[0] < intervals[i][0] {
			res = append(res, newInterval)
			break
		}

		//到这里 肯定 newInterval[0] >= intervals[i][0]
		res = append(res, intervals[i])
		i++

	}

	//newInterval 排在最后
	if i == len(intervals) {
		res = append(res, newInterval)
	}
	// }

	//如果res 最后两个区间有重叠 ,最后一个合入到前一个
	if len(res) > 1 && res[len(res)-1][0] <= res[len(res)-2][1] {

		if res[len(res)-1][1] > res[len(res)-2][1] {
			res[len(res)-2][1] = res[len(res)-1][1]
		}
		res = res[:len(res)-1]
	}

	resLen := len(res)
	// 走到这里  intervals[i] 的开始位置 肯定 也肯定大于 res 最后一个区间的开始位置
	for i < len(intervals) {
		// 剩余的跟 res 的最后一个不重叠，都在res的后面
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i:]...)
			break
		}

		//这里  res[reslen-1][0] < intervas[i][0] <= res[reslen-1][1]
		if intervals[i][1] > res[resLen-1][1] {
			res[resLen-1][1] = intervals[i][1]
		}
		i++

	}

	return res
}

func main() {

	a := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
		// []int{1,2},

	}
	b := []int{4, 8}

	fmt.Println(insert(a, b))

	a1 := [][]int{
		[]int{1, 5},
		// []int{2, 7},
		// []int{1,2},

	}
	b1 := []int{2, 7}
	fmt.Println(insert(a1, b1))

}
