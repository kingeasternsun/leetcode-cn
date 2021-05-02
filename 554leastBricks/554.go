/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-02 20:29:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-02 20:38:51
 * @FilePath: /554leastBricks/554.go
 */
package leetcode

func leastBricks(wall [][]int) int {
	lenMap := make(map[int]int)
	for _, row := range wall {

		tmpLen := 0
		for i, brick := range row {
			if i == len(row)-1 {
				continue
			}
			tmpLen += brick
			lenMap[tmpLen] = lenMap[tmpLen] + 1
		}

	}

	//看哪一种长度的砖缝最多
	maxCnt := 0
	for _, v := range lenMap {
		if v > maxCnt {
			maxCnt = v
		}
	}

	return len(wall) - maxCnt

}
