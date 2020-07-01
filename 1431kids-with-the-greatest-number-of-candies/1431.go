package main

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := make([]bool, len(candies))
	max := 0
	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := range candies {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
