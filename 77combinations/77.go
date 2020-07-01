package main

func main() {
	combine(5, 4)
}

// https://leetcode-cn.com/problems/combinations/
// dfs
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	if k == 0 || k > n {
		return result
	}

	if k == n {
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = i + 1
		}
		result = append(result, tmp)
		return result
	}

	tmpArray := make([]int, 0)
	dfs(tmpArray, 0, 0, n, k, &result)
	return result

}

// {2}, 1, 2,3,3
func dfs(tmpArray []int, tmpLen int, lastValue int, n, k int, result *[][]int) {

	if tmpLen == k {
		tmpArray = append(tmpArray[:0:0], tmpArray...)
		*result = append(*result, tmpArray)
		return
	}

	// can not make k
	if tmpLen+n-lastValue < k {
		return
	}

	i := lastValue + 1

	for i <= n {
		dfs(append(tmpArray, i), tmpLen+1, i, n, k, result)
		i++
	}

	return
}
