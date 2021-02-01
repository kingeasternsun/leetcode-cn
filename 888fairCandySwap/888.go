package leetcode

import "sort"

/*
最后等价转换为求两个有序数组中的相同数字
*/
func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	sort.Ints(A)
	sort.Ints(B)

	//双指针遍历 A，B，判断是否存在 a=b+target
	target := (sumA - sumB) / 2
	i := 0
	j := 0
	for i < len(A) && j < len(B) {

		if A[i] == B[j]+target {
			return []int{A[i], B[j]}
		}

		if A[i] < B[j]+target {
			i++
		} else {
			j++
		}

	}

	// }

	return []int{0, 0}
}

func fairCandySwapHash(A []int, B []int) []int {

	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	aMap := make(map[int]struct{}, 0)
	for _, v := range A {
		aMap[v] = struct{}{}
	}

	delta := (sumA - sumB) / 2
	for _, v := range B {
		if _, ok := aMap[v+delta]; ok {
			return []int{v + delta, v}
		}
	}

	return nil

}
