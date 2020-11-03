package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, 0)
	for _, v := range nums1 {
		m1[v] = m1[v] + 1
	}

	res := make([]int, 0)
	for _, k := range nums2 {
		if v, ok := m1[k]; ok {
			res = append(res, k)
			if v == 1 {
				delete(m1, k)
			} else {
				m1[k] = v - 1
			}
		}
	}

	return res
}
