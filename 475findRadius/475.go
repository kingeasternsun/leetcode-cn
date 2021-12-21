package leetcode

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {

	sort.Ints(heaters)
	ans := 0
	for _, h := range houses {
		tmp := math.MaxInt32
		i := sort.SearchInts(heaters, h)

		// consider the next heater
		if i < len(heaters) && heaters[i]-h < tmp {
			tmp = heaters[i] - h
		}

		// consider the pre heater
		if i > 0 && h-heaters[i-1] < tmp {
			tmp = h - heaters[i-1]
		}

		if tmp > ans {
			ans = tmp
		}

	}

	return ans
}
