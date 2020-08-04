package main

func findMagicIndex(nums []int) int {

	count := len(nums)
	if count == 0 {
		return -1
	}

	for i, n := range nums {
		if i == n {
			return i
		}
	}

	return -1

}
