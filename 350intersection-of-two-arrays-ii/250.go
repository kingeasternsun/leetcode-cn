package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums1 {
		m[v] = m[v] + 1
	}

	var res []int
	for _, v := range nums2 {
		if cnt, ok := m[v]; ok {
			res = append(res, v)
			cnt = cnt - 1
			if cnt == 0 {
				delete(m, v)
			} else {
				m[v] = cnt
			}
		}
	}

	return res

}
