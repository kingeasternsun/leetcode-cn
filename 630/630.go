package main

import "sort"

//[[5,15],[3,19],[6,7],[2,10],[5,16],[8,14],[10,11],[2,19]]
func main() {

	courses := [][]int{
		[]int{100, 200},
		[]int{200, 1300},
		[]int{1000, 1250},
		[]int{2000, 3200},
	}
	scheduleCourse(courses)
}

/*
https://leetcode-cn.com/problems/course-schedule-iii/

*/

func scheduleCourse(courses [][]int) int {

	cLen := len(courses)
	if cLen == 0 {
		return 0
	}

	if cLen == 1 {
		if courses[0][0] > courses[0][1] {
			return 0
		}

		return 1
	}

	sort.SliceStable(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	var preMaxBeg int //dp 是 递增的，类似 111122223333这样的 ，preMaxBeg所示当前最大值的左边界

	dp := make([]int, courses[len(courses)-1][1]+1)
	for i, cr := range courses {

		if cr[1] == 0 {
			continue
		}

		//填充:对于上一门课程，我们已经得到了到时间preEnd的可以进行的最大课程数目
		preEnd := 0
		if i > 0 {
			preEnd = courses[i-1][1]
			if cr[1] > preEnd {
				for j := preEnd + 1; j <= cr[1]; j++ {
					dp[j] = dp[preEnd]
				}
			}
		}

		beg := cr[1] - cr[0] + 1 //如果加入当前课程，当前课程的最晚开始时间是pre
		if beg <= 0 {
			continue
		}

		if beg <= preMaxBeg {
			// 如果这个课程在最晚时间beg开始学习 然后进行比较
			if dp[beg]+1 > dp[cr[1]] {
				dp[cr[1]] = dp[beg] + 1
			}

			preMaxBeg = cr[1]

			// continue
		} else {
			// 课程可以从 preMaxBeg+1 开始，或者 一直到 从 beg 开始
			// 如果这个课程从 preMaxBeg+1 开始学习,那么结束这个课程的时间是 preMaxBeg + cr[0]
			// 如果这个课程从 beg 开始学习， 那么这个课程的结束时间是 cr[1]
			for j := preMaxBeg + cr[0]; j <= cr[1]; j++ {
				dp[j] = dp[preMaxBeg] + 1
			}

			preMaxBeg = preMaxBeg + cr[0]
		}

	}

	return dp[courses[cLen-1][1]]
}
