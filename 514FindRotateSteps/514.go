package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(findRotateStepsBFS("godding", "gd") == 4)
	fmt.Println(findRotateStepsBFS("d", "dddd") == 4)
	fmt.Println(findRotateStepsBFS("abcde", "ade") == 6) // 1 2 4;1 3 2
	fmt.Println(findRotateStepsBFS("ababcab", "acbaacba") == 17)
	fmt.Println(findRotateStepsBFS("caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx") == 137)
	// "ababcab"
	// "acbaacba"

	// 	"caotmcaataijjxi"
	// "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"
	fmt.Println(findRotateSteps("godding", "gd") == 4)
	fmt.Println(findRotateSteps("d", "dddd") == 4)
	fmt.Println(findRotateSteps("abcde", "ade") == 6) // 1 2 4;1 3 2
	fmt.Println(findRotateSteps("ababcab", "acbaacba") == 17)
	fmt.Println(findRotateSteps("caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx") == 137)
}

/*
 dp[i][j] 表示完成key中第ｉ个字符时候，ring中第j个字符在12点方向　所需要的步骤.只跟 dp[i-1][*] 相关
*/
func findRotateSteps(ring string, key string) int {

	if len(key) == 0 {
		return 0
	}

	dp := make([][]int, len(key))
	for i := range dp {
		dp[i] = make([]int, len(ring))
	}

	for i := 0; i < len(ring); i++ {

		if ring[i] == key[0] {
			dp[0][i] = min(i, len(ring)-i) + 1
		} else {
			dp[0][i] = math.MaxInt32
		}

	}

	for i := 1; i < len(key); i++ {

		for j := 0; j < len(ring); j++ {

			if key[i] == ring[j] {

				minVal := math.MaxInt32
				for k := 0; k < len(ring); k++ {

					if key[i-1] != ring[k] {
						continue
					}

					step := min(abs(j-k), len(ring)-abs(j-k)) + 1
					if dp[i-1][k]+step < minVal {
						minVal = dp[i-1][k] + step
					}

				}

				dp[i][j] = minVal

			} else {
				dp[i][j] = math.MaxInt32
			}

		}
	}

	minStep := math.MaxInt32
	for j := 0; j < len(ring); j++ {
		if dp[len(key)-1][j] < minStep {
			minStep = dp[len(key)-1][j]
		}
	}

	return minStep

}

// func findMinStep(startID int, ring string, c byte) int {

// 	if ring[startID] == c {
// 		return 1
// 	}
// 	minStep := 100

// 	for i := 0; i < len(ring); i++ {
// 		if ring[i] == c {

// 			step := i - startID
// 			if step < 0 {
// 				step = -step
// 			}

// 			minStep = min(minStep, min(step, len(ring)-step))
// 		}
// 	}

// 	return minStep + 1
// }

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return -i
}

type Node struct {
	i     int //ring现在在ｉ
	j     int
	steps int
}

func findRotateStepsBFS(ring string, key string) int {
	if len(key) == 0 {
		return 0
	}

	if len(ring) == 1 { //d  dddd
		return len(key)
	}

	dp := make([][]int, len(key)) //必须缓存判断剪纸，不然时间超时
	for i := range dp {
		dp[i] = make([]int, len(ring))
		for j := 0; j < len(ring); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	ringBytes := []byte(ring)

	cur := make([]Node, 0)
	cur = append(cur, Node{i: 0, j: -1, steps: 0})

	minSteps := math.MaxInt32
	for len(cur) > 0 {

		newCur := make([]Node, 0)
		for _, c := range cur {

			//已经完成ｋｅｙ
			if c.j == len(key)-1 {
				if c.steps < minSteps {
					minSteps = c.steps
				}
				continue
			}

			//当前ｒｉｎｇ上的字符和key上的下一个字符相同，直接点击就可以
			if ringBytes[c.i] == key[c.j+1] {
				c.j++
				c.steps++
				newCur = append(newCur, c)
				continue
			}

			candis := findNearNode(c.i, ringBytes, key[c.j+1])
			for _, d := range candis { // c.i 的　两边跟　　key[c.j+1] 相同字符所在位置
				step := abs(d - c.i)
				if len(ring)-step < step {
					step = len(ring) - step
				}

				if c.steps+step+1 < dp[c.j+1][d] {
					newCur = append(newCur, Node{i: d, j: c.j + 1, steps: c.steps + step + 1})
					dp[c.j+1][d] = c.steps + step + 1
				}

			}

		}

		cur = newCur
	}

	return minSteps
}

//查找　ｒｉｎｇ　中　startID 等于　c 的字符的下标
func findNearNode(startID int, ring []byte, c byte) []int {
	var res []int
	i := (startID + 1) % len(ring)
	for i != startID {
		if ring[i] == c {
			res = append(res, i)
			// break
		}
		i++
		if i == len(ring) {
			i = 0
		}
	}

	// if startID == 0 {
	// 	startID = len(ring) - 1
	// } else {
	// 	startID = i - 1
	// }

	// for i != startID {
	// 	if ring[i] == c && i == res[0] {
	// 		res = append(res, i)
	// 		break
	// 	}
	// 	if i == 0 {
	// 		i = len(ring) - 1
	// 	} else {
	// 		i--
	// 	}
	// }

	return res

}
