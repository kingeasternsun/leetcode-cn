package main

func main() {

	a := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}

	maxEnvelopes(a)
}

//dp
/*
func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {

		return envelopes[i][0] < envelopes[i][1]

	})

	// fmt.Println(envelopes)

	dp := make([]int, len(envelopes)) //表示第几个信封里面最多套了多少个
	dp[0] = 1                         //

	for i, ev := range envelopes {

		for j := i - 1; j >= 0; j-- {

		}
	}

	return 0

}
*/
