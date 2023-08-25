package findorder

// 8ms 6.44mb
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 每个课程的入度
	inCnt := make([]int, numCourses)
	// 记录每个课程后面可以学习的课程
	outMap := make(map[int][]int, 0)

	for _, pre := range prerequisites {
		// pre[1] -> pre[0]
		inCnt[pre[0]] += 1
		outMap[pre[1]] = append(outMap[pre[1]], pre[0])
	}

	root := -1
	ret := []int{}
	for {

		if root < 0 {
			for c, cnt := range inCnt {
				if cnt == 0 {
					root = c
				}
			}
		}

		if root < 0 {
			return []int{}
		}

		ret = append(ret, root)
		inCnt[root] = -1
		if len(ret) == numCourses {
			return ret
		}

		// root -> out
		nextRoot := -1
		for _, out := range outMap[root] {
			inCnt[out] -= 1
			if inCnt[out] == 0 {
				nextRoot = out
			}
		}

		root = nextRoot

	}
}
