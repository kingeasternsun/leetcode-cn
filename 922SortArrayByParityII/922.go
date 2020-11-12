package main

func main() {

}
func sortArrayByParityII(A []int) []int {

	cnt := len(A)
	if cnt < 2 {
		return nil
	}

	res := make([]int, cnt)
	i := 0
	j := 1

	for _, v := range A {
		if v&1 == 0 {
			res[i] = v
			i += 2
		} else {
			res[j] = v
			j += 2
		}

	}

	return res

}
