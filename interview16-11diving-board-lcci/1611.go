package main

//等差数列 因为只能用这两种板子，所以逐渐增加longer长度的板子的个数就可以
func divingBoard(shorter int, longer int, k int) []int {

	if k == 0 {
		return nil
	}

	if shorter == longer {
		return []int{k * shorter}
	}

	base := k * shorter
	res := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		res[i] = base + (longer-shorter)*i
	}

	return res

}
