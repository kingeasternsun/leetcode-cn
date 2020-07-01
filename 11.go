package main

// func main() {
// 	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
// 	fmt.Println(maxArea(a))
// }

func maxArea(height []int) int {

	j := len(height)
	if j == 0 {
		return 0
	}

	j = j - 1

	i := 0

	var maxAr, tmpArea int

	for {
		if i >= j {
			return maxAr
		}

		if height[i] > height[j] {
			tmpArea = height[j] * (j - i)
			j = j - 1
		} else {
			tmpArea = height[i] * (j - i)
			i = i + 1

		}
		// fmt.Println(tmpArea)
		if tmpArea > maxAr {
			maxAr = tmpArea
		}

	}

}
