package main

/*
func main() {

	for i := 1; i <= 6; i++ {
		fmt.Println(getPermutation(3, i))

	}

	for i := 1; i <= 9; i++ {
		fmt.Println(getPermutation(4, i))

	}

	// fmt.Println(getPermutation(3, 3))

}
*/
func getPermutation(n int, k int) string {

	if n > 9 || n <= 0 {
		return ""
	}

	s := []byte("123456789")
	result := make([]byte, n)

	// n! 的数值
	var factorial = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880} //阶乘0-10
	if k > factorial[n] {
		return ""
	}

	//这个 减一 处理很关键
	k = k - 1

	//前1 ~ i-1 个字符一样，不同的i字符，后面都有(n-i)! 个方式

	//从高到底
	for i := n - 1; i >= 0; i-- {
		// fmt.Println(k, factorial[i])
		id := k / factorial[i] //第i个位置上的数字，
		k = k - factorial[i]*id
		// fmt.Println(id, k, factorial[i], n-i-1, i)
		result[n-i-1] = s[id]
		s = append(s[:id], s[id+1:]...)
		// fmt.Println(string(s))
	}

	return string(result)

}
