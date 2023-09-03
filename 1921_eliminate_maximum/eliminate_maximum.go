package eliminatemaximum

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	arriveTime := make([]int, len(dist))
	for i := range dist {
		arriveTime[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(arriveTime)
	ret := 0
	// i 表示，在时间 i 要击败的第 i 个怪兽
	for i, t := range arriveTime {
		if i < t {
			ret++
		} else {
			break
		}

	}
	return ret

}
