package main

func main() {

	v := [][]int{
		[]int{1, 0},
	}
	canFinish(3, v)
}

/*
https://leetcode-cn.com/problems/course-schedule/submissions/ 课程表



*/
func canFinish(numCourses int, prerequisites [][]int) bool {

	g := make(map[int][]int, 0)   //记录每个课程的后续学习课程
	inNum := make(map[int]int, 0) //记录每个课程入度，也就是如果课程A依赖课程B，那么课程A的入度就要加一

	cLen := len(prerequisites)
	if cLen == 0 {
		return true
	}

	for i := 0; i < cLen; i++ {
		inNum[prerequisites[i][0]] = inNum[prerequisites[i][0]] + 1
		g[prerequisites[i][1]] = append(g[prerequisites[i][1]], prerequisites[i][0])
	}

	if len(inNum) == 0 {
		return true
	}

	q := make([]int, 0)
	qLen := 0
	for i := 0; i < numCourses; i++ {

		if inNum[i] == 0 {
			q = append(q, i)
			qLen++
		}
	}

	if qLen == numCourses {
		return true
	}

	learnNum := 0

	for {
		if qLen == 0 {
			break
		}

		newLen := 0
		for i := 0; i < qLen; i++ {
			cur := q[i]
			learnNum++
			for _, post := range g[cur] {
				inNum[post] = inNum[post] - 1
				if inNum[post] == 0 {
					q = append(q, post)
					newLen++
				}
			}
		}

		q = q[qLen:]
		qLen = newLen

	}

	if learnNum != numCourses {
		return false
	}

	return true

}
