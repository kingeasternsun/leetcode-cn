package main

/*

https://leetcode-cn.com/problems/jump-game-iv/
1345. 跳跃游戏 IV
BFS
- 左右位置，以及具有相同数值的节点都是相邻节点

*/
func minJumps(arr []int) int {

	inLen := len(arr)
	if inLen == 0 || inLen == 1 {
		return 0
	}

	m := make(map[int][]int, 0)
	inQueue := make([]int, inLen) //节点是否已经加入到队列
	for i := range arr {
		m[arr[i]] = append(m[arr[i]], i)
	}

	step := 0
	q := make([]int, 0)

	//下面两个初始化很重要
	q = append(q, 0) //从0开始 广度优先遍历
	inQueue[0] = 1   //要标记0 已经加入到队列中了

	for len(q) > 0 {

		//当前层有多少个节点
		qLen := len(q)
		//把当前层的节点都取出来，然后将这些节点的相邻节点加入队列
		for qLen > 0 {

			// 取出头部第一个点
			curID := q[0]

			if curID >= inLen-1 {
				return step
			}
			q = q[1:]
			//左边的点没有访问过,就把左边的点入队列，然后更新标记
			if curID > 0 && inQueue[curID-1] == 0 {
				inQueue[curID-1]++
				q = append(q, curID-1)
			}

			//左边的点没有访问过,就把右边的点入队列，然后更新标记
			if curID+1 < inLen && inQueue[curID+1] == 0 {
				inQueue[curID+1]++
				q = append(q, curID+1)
			}

			// 具有相同数值的节点,也是相邻节点
			tmpList := m[arr[curID]]
			for i := range tmpList {
				if inQueue[tmpList[i]] == 0 {
					inQueue[tmpList[i]]++
					q = append(q, tmpList[i])
				}
			}

			// 相同数值的都加入到queue中了，所以可以删除了
			delete(m, arr[curID])

			qLen--
		}

		//这一层访问过了
		step++
	}

	// 这道题的信息，肯定会到达最后一个节点，所以这个地方是不会访问到的
	return step - 1
}
