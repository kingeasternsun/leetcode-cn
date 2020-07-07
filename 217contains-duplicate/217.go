package main

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{}, 0)
	for _, n := range nums {

		if _, ok := m[n]; ok {
			return true
		}

		m[n] = struct{}{}

	}

	return false
}
