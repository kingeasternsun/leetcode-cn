package merge

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pre := intervals[0]
	ret := make([][]int, 0)
	for _, inter := range intervals {
		if inter[0] <= pre[1] {
			if inter[1] > pre[1] {
				pre[1] = inter[1]
			}
		} else {
			ret = append(ret, pre)
			pre = inter
		}
	}

	return append(ret, pre)
}
