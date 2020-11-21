package main

import "sort"

func main() {

}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	res := make([][]int, 0)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {

			res = append(res, []int{r, c})

		}
	}

	sort.Slice(res, func(i, j int) bool { return manDis(res[i], r0, c0) < manDis(res[j], r0, c0) })

	return res
}

func manDis(v []int, r0 int, c0 int) int {
	return abs(v[0]-r0) + abs(v[1]-c0)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
