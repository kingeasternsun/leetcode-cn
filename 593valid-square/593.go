package main

import "sort"

/*
4边相等
对角线是边长的平方的2倍
*/
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	data := [][]int{p1, p2, p3, p4}

	sort.Slice(data, func(i, j int) bool {
		if data[i][0] < data[j][0] {
			return true
		}
		if data[i][0] > data[j][0] {
			return false
		}

		return data[i][1] < data[j][1]
	})

	lenth := sqrLen(data[0], data[1])
	if lenth == 0 {
		return false
	}
	if lenth != sqrLen(data[0], data[2]) {
		return false
	}

	if lenth != sqrLen(data[3], data[1]) {
		return false
	}

	if lenth != sqrLen(data[3], data[2]) {
		return false
	}

	if 2*lenth != sqrLen(data[1], data[2]) {
		return false
	}

	return true
}

func sqrLen(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
