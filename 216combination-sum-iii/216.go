package main

func main() {
	combinationSum3(2, 18)
}

func combinationSum3(k int, n int) [][]int {

	var res = [][]int{}
	var vec = []int{}
	combinate(k, n, 0, vec, &res)
	return res
}

//  选取k个大于 curMax 的数，组成 n  // 0ms 2MB
func combinate(k int, n int, curMax int, vec []int, res *[][]int) {

	// 这里肯定大于 n  ，可以提前返回
	if k*curMax >= n {
		return
	}

	if k == 1 {

		if n > curMax && n <= 9 {
			vec = append(vec, n)
			vec = append(vec[:0:0], vec...)
			*res = append(*res, vec)
		}

		return

	}

	//如果平均大于9，那么肯定至少有一个 大于 9
	if n/k >= 9 {
		return
	}

	// end := 9
	// if n/k < end {
	// 	end = n / k
	// }

	for i := curMax + 1; i <= n/k; i++ {
		combinate(k-1, n-i, i, append(vec, i), res)
	}

}
