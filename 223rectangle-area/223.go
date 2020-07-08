package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {

	joinArea := computeJoinArea(A, B, C, D, E, F, G, H)

	return abs(A-C)*abs(B-D) + abs(E-G)*abs(F-H) - joinArea

}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

//计算两个区间的相交长度
func join(A, C, E, G int) int {
	if A > C {
		A, C = C, A
	}

	if E > G {
		E, G = G, E
	}

	if A > E {
		A, E = E, A
		C, G = G, C
	}

	// [A 		C]
	// [ E G]
	// [ E 			G]
	// [ 				E 	G]

	if E >= C {
		return 0
	}

	// E <= A
	return min(C, G) - E
}

//计算相交的区间面积
func computeJoinArea(A int, B int, C int, D int, E int, F int, G int, H int) int {

	xj := join(A, C, E, G)
	if xj == 0 {
		return 0
	}
	yj := join(B, D, F, H)
	if yj == 0 {
		return 0
	}

	return xj * yj
}
