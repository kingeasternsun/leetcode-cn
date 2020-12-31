package main

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	i := 0
	j := 0

	cnt := 0

	//g是胃口  j是饼干大小
	for i < len(g) && j < len(s) {

		if s[j] >= g[i] {
			j++
			i++
			cnt++
		} else {

			j++ //加大饼干大小
		}

	}

	return cnt
}
