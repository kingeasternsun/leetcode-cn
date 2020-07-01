package main

import "math"

// https://leetcode-cn.com/problems/01-matrix/
// bfs
type Point struct {
	x, y int
}

// 72ms 8MB
func updateMatrix(matrix [][]int) [][]int {

	rows := len(matrix)
	if rows == 0 {
		return nil
	}

	cols := len(matrix[0])
	if cols == 0 {
		return nil
	}

	cur := make([]Point, 0)
	curLen := 0
	visted := make([]bool, rows*cols)

	result := make([][]int, rows)
	for row := 0; row < rows; row++ {
		result[row] = make([]int, cols)

		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				result[row][col] = 0

				cur = append(cur, Point{row, col})
				curLen++
				visted[row*cols+col] = true

			} else {
				result[row][col] = math.MaxInt32
			}
		}
	}

	if curLen == rows*cols {
		return result
	}

	dir := []int{0, 1, 0, -1, 0}
	for curLen > 0 {

		newPoints := make([]Point, 0)
		addMap := make(map[int]struct{}, 0)
		for i := 0; i < curLen; i++ {

			// try the relate point
			for d := 0; d < 4; d++ {
				newPoint := Point{
					x: cur[i].x + dir[d],
					y: cur[i].y + dir[d+1],
				}

				if newPoint.x < 0 || newPoint.x >= rows || newPoint.y < 0 || newPoint.y >= cols {
					continue
				}

				// not visted, only this layer finish, we can set the visted flag
				if !visted[newPoint.x*cols+newPoint.y] && result[cur[i].x][cur[i].y]+1 < result[newPoint.x][newPoint.y] {

					//update distance
					result[newPoint.x][newPoint.y] = result[cur[i].x][cur[i].y] + 1
					// in each layer of bfs, the newPoint can by be vitsted many times,so remove dup add
					if _, ok := addMap[newPoint.x*cols+newPoint.y]; !ok {
						newPoints = append(newPoints, newPoint)
					}

				}
			}

		}

		// this layer has finish ,so set the visted
		for i := range newPoints {
			visted[newPoints[i].x*cols+newPoints[i].y] = true
		}

		cur = newPoints
		curLen = len(newPoints)

	}

	return result
}
