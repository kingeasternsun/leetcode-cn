package main

import "fmt"

//https://leetcode-cn.com/problems/rotting-oranges/
func main() {

	grid := [][]int{
		// {2, 1, 1},
		// {1, 1, 0},
		// {0, 1, 1},
		{0, 2},
	}

	fmt.Println(orangesRotting(grid))

}

type Point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	dir := []int{0, 1, 0, -1, 0}
	rLen := len(grid)
	if rLen == 0 {
		return 0
	}
	cLen := len(grid[0])
	if cLen == 0 {
		return 0
	}

	step := 0
	queue := make([]Point, 0)
	qLen := 0
	orangeNum := 0 //正常的橘子个数
	visted := make([]int, rLen*cLen)
	var startPoint Point

	// 这个比较关键，第一次要把所有的已经坏掉的放入到队列中
	var findStart = func() {
		for i := 0; i < rLen; i++ {
			for j := 0; j < cLen; j++ {
				if grid[i][j] == 2 {
					queue = append(queue, Point{i, j})
					qLen++
					// vistedLen++
					visted[i*cLen+j] = 1 // 这里要记得标记访问过
				} else if grid[i][j] == 1 {
					orangeNum++
				}
			}
		}

		return
	}
	findStart()

	// fmt.Println(queue)

	// 这里比较逗， 如果没有新鲜的橘子，就返回0
	if orangeNum == 0 {
		return 0
	}

	//没有坏掉的
	if qLen == 0 {
		return -1
	}

	for {

		// 这里关键，因为这个题目的重点在于要把正常的橘子全部感染，所以这里要加个判断，不过会多判断一层
		if orangeNum == 0 {
			return step
		}

		if qLen == 0 {
			return -1
		}

		newQlen := 0

		//当前层中的每一个点
		for i := 0; i < qLen; i++ {

			// pop front
			startPoint = queue[i]
			// queue = queue[1:]
			// qLen--

			// 四个方向的点
			for j := 0; j < 4; j++ {
				tmpPoint := Point{
					startPoint.x + dir[j],
					startPoint.y + dir[j+1],
				}
				if tmpPoint.x < 0 || tmpPoint.y < 0 || tmpPoint.x >= rLen || tmpPoint.y >= cLen {
					continue
				}

				// 已经访问过
				if visted[tmpPoint.x*cLen+tmpPoint.y] == 1 {
					continue
				}

				// 没有橘子跳过
				if grid[tmpPoint.x][tmpPoint.y] == 0 {
					continue
				}

				// 感染放进去
				queue = append(queue, tmpPoint)
				newQlen++
				visted[tmpPoint.x*cLen+tmpPoint.y] = 1 // 这里要记得标记访问过

				if grid[tmpPoint.x][tmpPoint.y] == 1 {
					orangeNum-- //正常的橘子少了一个
				}

			}

		}

		queue = queue[qLen:]
		qLen = newQlen
		step++

	}

	// if orangeNum == 0 {
	// 	return step
	// }

	return -1

}
