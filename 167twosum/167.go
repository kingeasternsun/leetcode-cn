package main

func twoSum(numbers []int, target int) []int {

	if len(numbers) < 2 {
		return nil
	}

	beg := 0
	end := len(numbers) - 1

	for beg < end {
		if numbers[beg]+numbers[end] == target {
			return []int{beg + 1, end + 1}
		}

		if numbers[beg]+numbers[end] > target {
			end--
		} else {
			beg++
		}
	}

	return nil
}
