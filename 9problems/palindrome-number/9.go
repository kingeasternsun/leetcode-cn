package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	a := make([]int, 0)
	for x > 0 {
		a = append(a, x%10)
		x = x / 10
	}

	count := len(a)

	for i := 0; i < count/2; i++ {
		if a[i] != a[count-i-1] {
			return false
		}
	}

	return true

}
