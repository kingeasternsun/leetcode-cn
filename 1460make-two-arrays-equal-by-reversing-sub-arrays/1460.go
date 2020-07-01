package main

func canBeEqual(target []int, arr []int) bool {

	m := make(map[int]int, 0)
	for i := range target {
		m[target[i]] = m[target[i]] + 1
	}

	for i := range arr {

		c := m[arr[i]]
		if c == 0 {
			return false
		}

		m[arr[i]] = m[arr[i]] - 1
	}

	return true

}
