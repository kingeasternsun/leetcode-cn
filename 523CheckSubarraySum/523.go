package main

import "fmt"

func main() {

	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6) == true)
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6) == true)
	fmt.Println(checkSubarraySum([]int{0, 0}, 0) == true)
	fmt.Println(checkSubarraySum([]int{0, 1, 0}, 0) == false)
}

func checkSubarraySum(nums []int, k int) bool {

	cnt := len(nums)
	if cnt <= 1 {
		return false
	}
	for i := 0; i < cnt; i++ {
		sum := 0
		for j := i; j < cnt; j++ {
			sum += nums[j]
			if j-i > 0 {

				if k == 0 && sum == 0 {
					return true
				}

				if k != 0 && sum%k == 0 {
					return true

				}
			}
		}
	}

	return false
}
