package main

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, 0)
	// m2 := make(map[int]struct{}, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		m1[v] = m1[v] + 1
	}

	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			res = append(res, v)
			delete(m1, v)
		}
	}

	return res
}
