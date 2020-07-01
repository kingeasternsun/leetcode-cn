package main

func maxScoreSightseeingPair(A []int) int {

	preIMax := 0
	max := 0

	for j := range A {
		if A[j]-j+preIMax > max {
			max = A[j] - j + preIMax
		}

		if A[j]+j > preIMax {
			preIMax = A[j] + j
		}
	}

	return max
}
