package main

import (
	"sort"
)

// func main() {

// 	var nums = []int{-1, 2, 1, -4}
// 	// sort.Ints(nums)
// 	fmt.Println(threeSumClosest(nums, 1))

// }

// 外侧循环后，内层双指针
func threeSumClosest(nums []int, target int) int {

	slen := len(nums)
	if slen == 0 {
		return 0
	}

	sort.Ints(nums)
	// fmt.Println(nums)

	var minAbs = 0x7fffffff

	var beg, end, result int

	for i := range nums {

		beg = 0
		end = slen - 1

		for {

			if beg >= end {
				break
			}

			if beg == i {
				beg = beg + 1
				continue
			}

			if end == i {
				end = end - 1
				continue
			}

			tmpS := nums[beg] + nums[end] + nums[i]

			// fmt.Printf("%v %v %v %v\n", i, beg, end, tmpS)

			tmpsAbs := tmpS - target
			if tmpsAbs == 0 {
				return target
			}

			if tmpsAbs < 0 {
				tmpsAbs = -tmpsAbs
				beg = beg + 1
			} else {
				end = end - 1
			}

			if tmpsAbs < minAbs {
				minAbs = tmpsAbs
				result = tmpS
				// fmt.Printf("%v %v %v %v\n", i, beg, end, result)
			}

		}
	}

	return result
}
