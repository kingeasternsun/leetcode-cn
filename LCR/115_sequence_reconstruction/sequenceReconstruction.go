package sequencereconstruction

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	order := make(map[int]map[int]struct{}, 0)
	for _, seq := range sequences {
		for i := 0; i < len(seq)-1; i++ {
			if order[seq[i]] == nil {
				order[seq[i]] = make(map[int]struct{})
			}
			order[seq[i]][seq[i+1]] = struct{}{}
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if _, ok := order[nums[i]][nums[i+1]]; !ok {
			return false
		}
	}
	return true
}
