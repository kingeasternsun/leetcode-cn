package main

//https://leetcode-cn.com/problems/distribute-candies-to-people/
func main() {

}

// 1103. 分糖果 II
func distributeCandies(candies int, num_people int) []int {

	if num_people == 0 {
		return nil
	}

	nums := make([]int, num_people)
	// for i := 0; i < candies; i++ {

	// }

	assignNum := 1

	for {
		for i := 0; i < num_people && candies > 0; i++ {

			if candies <= assignNum {
				nums[i] += candies
				return nums
			}

			nums[i] += assignNum
			candies -= assignNum

			assignNum++

		}
	}

	return nums

}
