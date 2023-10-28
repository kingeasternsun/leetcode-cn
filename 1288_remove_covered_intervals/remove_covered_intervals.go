package removecoveredintervals

import "sort"

// 12ms
func removeCoveredIntervalsa(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		// NOTE: [[3,10],[4,10],[5,11]] ,we need to put [4,10] before [3,10]
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	ans := len(intervals)
	for i := 0; i < len(intervals); i++ {

		cover := 0
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] >= intervals[j][0] {
				cover = 1
				break
			}
		}
		ans -= cover

	}
	return ans

}

// 8ms
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		// NOTE: [[3,4],[3,10],[5,11]] ,we need to put [3,10]] before [3,4]
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	//  如果 i<j, 那么 interval[i].left <= interval[j].left
	//  所以如果 interval[j] 的左边有 interval[i] 的右边界 大于 interval[j]的右边界，那么当前 j 就是被覆盖了
	ans := len(intervals)
	preMax := -1
	for i := 0; i < len(intervals); i++ {

		if i == 0 {
			preMax = intervals[0][1]
			continue
		}

		// 左边有区间的右边界超过了当前区间的右边界
		if preMax >= intervals[i][1] {
			ans--
		}

		if intervals[i][1] > preMax {
			preMax = intervals[i][1]
		}

	}
	return ans

}
