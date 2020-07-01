package main

func twoSum(nums []int, target int) []int {
	count := len(nums)
	if count <= 1 {
		return nil
	}

	i := 0
	j := count - 1

	if nums[count-1] < target/2 {
		return nil
	}

	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		}

		if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}

	}

	return nil

}
