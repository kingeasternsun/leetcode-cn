package main

import "math"

func main() {

}

/*

利用BFS
			 不能只考虑向下 或向右
			[
			[0,1,0,0,0,1,0,0],
			[0,1,0,1,0,1,0,1],
			[0,0,0,1,0,0,1,0]]
			1

*/

type Point struct {
	x, y int
}

func shortestPath(grid [][]int, k int) int {

	dirs := []int{0, 1, 0, -1, 0}
	rLen := len(grid)
	if rLen == 0 {
		return 0
	}
	cLen := len(grid[0])

	ems := make([]int, rLen*cLen) //记录到达每个位置上的所需拆除的数目
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			ems[i*cLen+j] = math.MaxInt32
		}
	}

	ems[0] = 0

	stack := make([]Point, 0)

	stack = append(stack, Point{0, 0})
	stLen := 1
	// maxE := 0
	steps := 0

	for {

		if stLen == 0 {
			return -1
		}

		newStackLen := 0

		for j := 0; j < stLen; j++ { // 当前层的每一个节点

			cur := stack[j]
			if cur.x == rLen-1 && cur.y == cLen-1 {
				return steps
			}

			// stack = stack[1 : stLen-1]
			// stLen--

			//到达当前位置需要拆除的数目
			curE := ems[cur.x*cLen+cur.y]

			/* 不能只考虑向下 或向右
			[
			[0,1,0,0,0,1,0,0],
			[0,1,0,1,0,1,0,1],
			[0,0,0,1,0,0,1,0]]
			1
			*/
			for i := 0; i < 4; i++ { //当前节点的每一个方向的相邻点
				tmp := Point{
					cur.x + dirs[i],
					cur.y + dirs[i+1],
				}

				if tmp.x < 0 || tmp.x >= rLen || tmp.y < 0 || tmp.y >= cLen {
					continue
				}

				//如果通过当前点到达相邻点要拆除更多的墙或跟之前没啥区别，就不用把这个邻居再放回去了
				// 或者到达需要拆除的墙个数超过了k
				if curE+grid[tmp.x][tmp.y] >= ems[tmp.x*cLen+tmp.y] || curE+grid[tmp.x][tmp.y] > k {
					continue
				}

				ems[tmp.x*cLen+tmp.y] = curE + grid[tmp.x][tmp.y]
				//把这个邻居再放回去
				stack = append(stack, tmp)
				// stLen++
				newStackLen++

			}
		}

		steps++

		stack = stack[stLen:]
		stLen = newStackLen

	}
}
