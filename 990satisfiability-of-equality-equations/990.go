package main

/*
 分组问题
 1. 先根据==将不同的字母分组
 2. 然后判断！=两边的字母是否在同一个组中
*/

// func max(i, j int)int{
// 	if i >j{
// 		return i
// 	}

// 	return j
// }

func findRoot(i int, group []int) int {
	for group[i] != i {
		i = group[i]
	}

	return i
}

func union(i, j int, height, group []int) {
	pi := findRoot(i, group)
	pj := findRoot(j, group)
	if pi == pj {
		return
	}

	if height[pi] > height[pj] {
		group[pj] = pi
		return
	}

	if height[pj] > height[pi] {
		group[pi] = pj
		return
	}

	group[pi] = pj
	height[pj] = height[pj] + 1
	return
}

func equationsPossible(equations []string) bool {

	if len(equations) == 0 {
		return false
	}

	group := make([]int, 26)
	height := make([]int, 26)
	for i := range group {
		group[i] = i
	}

	for _, eq := range equations {

		if eq[1] == '=' {
			union(int(eq[0]-'a'), int(eq[3]-'a'), height, group)
		}
	}

	for _, eq := range equations {

		if eq[1] == '!' {
			pi := findRoot(int(eq[0]-'a'), group)
			pj := findRoot(int(eq[3]-'a'), group)
			if pi == pj {
				return false
			}
		}
	}

	return true
}
