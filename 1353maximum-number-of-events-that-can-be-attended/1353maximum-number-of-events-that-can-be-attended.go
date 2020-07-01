package main

import (
	"fmt"
	"sort"
)

func main() {

	events := [][]int{
		// []int{1, 2},
		// []int{2, 3},
		// []int{3, 4},
		[]int{1, 100000},
	}

	fmt.Println(maxEvents(events))

}

/*
这个题目注意的是，每个会议不是要全程参加的，也就是一个会议只需要参加一天就算参加
- 所以最贪心的就是按照会议的结束时间排序后，每个会议只参加可以参加的最早的第一天
*/

//300ms 16.6MB
func maxEvents1(events [][]int) int {

	eLen := len(events)
	if eLen <= 1 {
		return eLen
	}

	minDay := events[0][0]
	sort.Slice(events, func(i int, j int) bool {
		if events[i][0] < minDay {
			minDay = events[i][0]
		}
		if events[j][0] < minDay {
			minDay = events[j][0]
		}
		return events[i][1] < events[j][1]
	})
	// fmt.Println(minDay)
	maxDay := events[eLen-1][1]

	days := make([]bool, maxDay-minDay+1)
	maxEvent := 0
	for i := 0; i < eLen; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if !days[j-minDay] {
				days[j-minDay] = true
				maxEvent++
				break
			}
		}
	}

	return maxEvent

}

//216ms 17.6MB
func maxEvents2(events [][]int) int {

	eLen := len(events)
	if eLen <= 1 {
		return eLen
	}

	// minDay := events[0][0]
	sort.Slice(events, func(i int, j int) bool {
		// if events[i][0] < minDay {
		// 	minDay = events[i][0]
		// }
		// if events[j][0] < minDay {
		// 	minDay = events[j][0]
		// }
		return events[i][1] < events[j][1]
	})
	// fmt.Println(minDay)
	maxDay := events[eLen-1][1]

	days := make([]bool, maxDay+1)
	maxEvent := 0
	for i := 0; i < eLen; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if !days[j] { //每个会议只占用一天
				days[j] = true
				maxEvent++
				break
			}
		}
	}

	return maxEvent

}
