package main

import (
	"fmt"
	"sort"
)

func main() {

	a := [][]int{
		[]int{10, 16},
		[]int{2, 8},
		[]int{1, 6},
		[]int{7, 12},
	}
	fmt.Println(findMinArrowShots(a))
	b := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
		[]int{7, 8},
	}
	fmt.Println(findMinArrowShots(b))

	c := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{4, 5},
	}
	fmt.Println(findMinArrowShots(c))

	d := [][]int{
		[]int{-2147483646, -2147483645},
		[]int{2147483646, 2147483647},
	}
	fmt.Println(findMinArrowShots(d))

}

func findMinArrowShots(points [][]int) int {

	cnt := len(points)
	if cnt == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[i][0]
		}

		return points[i][1] < points[j][1]

	})

	mark := points[0][1]
	res := 1
	for _, p := range points {

		//没有重叠了，就要新发射一个
		if p[0] > mark {

			res++
			mark = p[1]

		}

	}

	return res
}
