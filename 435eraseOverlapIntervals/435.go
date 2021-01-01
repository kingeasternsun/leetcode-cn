package leetcode

import "sort"

/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:

可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-overlapping-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 为了选择最大数量的不重叠区间，那么第一个区间肯定是右边界最靠左边的，选择第一个区间后，剩下不重叠的区间里面再选择右边节最左的 ,
// 类似题目 https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
func eraseOverlapIntervals(intervals [][]int) int {

	if len(intervals) <= 1 {
		return 0
	}

	//所以每次选择与上一个区间不冲突的剩余区间里面，右边区间最靠左边的
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	preRight := intervals[0][1] //上一个选择的区间的右边边界

	for _, v := range intervals[1:] {

		if v[0] < preRight {
			//有重叠，就要移除
			res++
			continue
		}

		//遇到不重叠的，加入
		preRight = v[1]

	}

	return res
}
