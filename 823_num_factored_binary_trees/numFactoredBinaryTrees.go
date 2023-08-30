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

// LC:1654 16ms 6.82mb
func minimumJumps(forbidden []int, a int, b int, x int) int {
	type Item struct {
		Pos int
		Dir bool
	}
	forbinnedMap := make(map[int]struct{}, 0)
	max_f := 0
	for _, v := range forbidden {
		forbinnedMap[v] = struct{}{}
		max_f = max(v, max_f)
	}

	visitedMap := make(map[Item]struct{}, 0)

	// 初始化
	queue := []Item{{0, true}}
	visitedMap[Item{0, true}] = struct{}{}

	step := 0
	for len(queue) > 0 {
		step += 1
		tmpLen := len(queue)
		for i := 0; i < tmpLen; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Pos == x {
				return step - 1
			}

			// 往右边跳
			if (tmp.Pos + a) <= (max_f + a + b + x) {
				_, forbid := forbinnedMap[tmp.Pos+a]
				_, visited := visitedMap[Item{tmp.Pos + a, true}]
				if (!forbid) && (!visited) {
					queue = append(queue, Item{tmp.Pos + a, true})
					visitedMap[Item{tmp.Pos + a, true}] = struct{}{}
				}
			}

			// 左边跳, 只有上一步是往右跳，才可以左跳
			if tmp.Dir && tmp.Pos-b >= 0 {
				_, forbid := forbinnedMap[tmp.Pos-b]
				_, visited := visitedMap[Item{tmp.Pos - b, false}]
				if (!forbid) && (!visited) {
					queue = append(queue, Item{tmp.Pos - b, false})
					visitedMap[Item{tmp.Pos - b, false}] = struct{}{}
				}
			}

		}

	}

	return -1

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
