package main

import "sort"

func main() {

}

// [2,3,1,3,2,4,6,7,9,2,19,18]
// [2,1,4,3,9,6]

//[2,2,2,1,4,3,3,9,6,7,18,19]
func relativeSortArray(arr1 []int, arr2 []int) []int {

	numMap := make(map[int]int, 0)
	for i, v := range arr2 {
		numMap[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		idi, oki := numMap[arr1[i]]
		idj, okj := numMap[arr1[j]]
		if (!oki) && (!okj) {
			return arr1[i] <= arr1[j]
		}

		if !oki {
			return false

		}

		if !okj {
			return true
		}

		return idi < idj
	})
	return arr1
}
