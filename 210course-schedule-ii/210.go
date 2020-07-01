package main

/*
210. 课程表 II https://leetcode-cn.com/problems/course-schedule-ii/
*/
func findOrder(numCourses int, prerequisites [][]int) []int {

	g := make(map[int][]int, 0)   //记录课程学完后后面可以学习哪些课程
	inNum := make(map[int]int, 0) //记录课程的入度，如果课程A需要先学习课程B，那么课程A的入度就加一

	result := make([]int, 0)
	pLen := len(prerequisites)
	if pLen == 0 {
		for i := 0; i < numCourses; i++ {
			result = append(result, i)
		}
		return result
	}

	for i := 0; i < pLen; i++ {
		g[prerequisites[i][1]] = append(g[prerequisites[i][1]], prerequisites[i][0])
		inNum[prerequisites[i][0]] = inNum[prerequisites[i][0]] + 1
	}

	q := make([]int, 0)
	qLen := 0
	for i := 0; i < numCourses; i++ {
		if inNum[i] == 0 { //入度为0 也就是不需要先修别的课程,可以提前先学习
			q = append(q, i)
			qLen++
		}
	}

	resultLen := 0
	for qLen > 0 {

		newLen := 0

		result = append(result, q...) //里面都是可以学习的课程
		resultLen = resultLen + qLen

		for i := 0; i < qLen; i++ { //学习这个课程

			//对每一个 q[i] 的后面的课程
			for _, post := range g[q[i]] {
				inNum[post]--         //因为前面的一门课程已经学习完了，入度就会少一
				if inNum[post] == 0 { //如果先修课程都已经学完了，就可以准备学习这个课程
					q = append(q, post)
					newLen++
				}
			}

		}

		//这一轮课程学习完后
		q = q[qLen:]
		qLen = newLen
	}
	if resultLen == numCourses {
		return result
	}

	return nil

}
