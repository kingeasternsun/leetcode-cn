package main

func containsNearbyDuplicate(nums []int, k int) bool {

	posMap := make(map[int]int, 0)
	for i, n := range nums {

		if prePos, ok := posMap[n]; ok {

			if i-prePos <= k {
				return true
			}

		}

		posMap[n] = i

	}

	return false

}
