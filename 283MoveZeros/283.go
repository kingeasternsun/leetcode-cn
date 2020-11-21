package main

func main() {
	a := []int{0}
	moveZeroes(a)
}

//4ms 3.7MB
func moveZeroes(nums []int) {

	cnt := len(nums)
	if cnt == 0 {
		return
	}

	i := 0 //i之前的都是大于0的,不包含i
	j := 0

	for j = 0; j < cnt; j++ {

		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}

	for j := i; j < cnt; j++ {
		nums[j] = 0
	}
	return

}
