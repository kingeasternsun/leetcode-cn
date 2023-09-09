package canfinish

// 12ms 6.1MB
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 计算每个课程的下一个课程
	inE := make(map[int]int, 0)
	zeroIn := make([]int, 0)
	nextCourses := make([][]int, numCourses)
	// v[1]->v[0]
	for _, v := range prerequisites {
		nextCourses[v[1]] = append(nextCourses[v[1]], v[0])
		inE[v[0]]++
	}

	for c := 0; c < numCourses; c++ {
		if inE[c] == 0 {
			zeroIn = append(zeroIn, c)
		}
	}

	for len(zeroIn) > 0 {
		cur := zeroIn[0]
		zeroIn = zeroIn[1:]
		for _, next := range nextCourses[cur] {
			inE[next] -= 1
			if inE[next] == 0 {
				zeroIn = append(zeroIn, next)
				delete(inE, next)
			}
		}

	}

	return len(inE) <= 0
}
