package numfactoredbinarytrees

import "sort"

// 8ms 3.48mb
func numFactoredBinaryTrees(arr []int) int {

	sort.Ints(arr)
	arr_map := make(map[int]int, 0)
	for _, v := range arr {
		arr_map[v] = 1
	}

	for i, v := range arr {
		for _, a := range arr[0:i] {
			b := v / a
			if b < a {
				break
			}
			if _, ok := arr_map[b]; ok && a*b == v {
				arr_map[v] += arr_map[a] * arr_map[b]
				if a != b {
					arr_map[v] += arr_map[a] * arr_map[b]
				}
			}
		}
	}

	sum := 0
	for _, v := range arr_map {
		sum += v
	}
	return sum % 1000000007
}
