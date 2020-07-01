package main

import (
	"math"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	if horizontalCuts[len(horizontalCuts)-1] != h {
		horizontalCuts = append(horizontalCuts, h)
	}

	if verticalCuts[len(verticalCuts)-1] != w {
		verticalCuts = append(verticalCuts, w)
	}
	max := math.MinInt64

	for i := range verticalCuts {

		// if i == 0 {
		// 	continue
		// }

		for j := range horizontalCuts {

			var tmp int

			if i == 0 && j == 0 {
				tmp = (horizontalCuts[j]) * (verticalCuts[i])
			} else if i == 0 {
				tmp = (horizontalCuts[j] - horizontalCuts[j-1]) * (verticalCuts[i])
			} else if j == 0 {
				tmp = (horizontalCuts[j]) * (verticalCuts[i] - verticalCuts[i-1])
			} else {
				tmp = (horizontalCuts[j] - horizontalCuts[j-1]) * (verticalCuts[i] - verticalCuts[i-1])
			}

			if tmp > max {
				max = tmp
			}

		}

	}

	return int(max % 1000000007)

}
