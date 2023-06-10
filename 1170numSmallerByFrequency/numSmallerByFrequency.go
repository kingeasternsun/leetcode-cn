package numsmallerbyfrequency

import (
	"sort"
)

func numSmallerByFrequency(queries []string, words []string) []int {

	wordRet := make([]int, len(words))
	for i, w := range words {
		wordRet[i] = frequencyOfMinByte(w)
	}

	sort.Ints(wordRet)
	queryRet := make([]int, len(queries))
	for i, q := range queries {
		qRet := frequencyOfMinByte(q)
		queryRet[i] = len(wordRet) - sort.Search(len(wordRet), func(i int) bool { return wordRet[i] > qRet })
	}

	return queryRet
}

func frequencyOfMinByte(word string) int {
	ret := 0
	minB := byte(255)
	for i := 0; i < len(word); i++ {
		if word[i] < minB {
			minB = word[i]
			ret = 1
		} else if word[i] == minB {
			ret += 1
		}
	}
	return ret
}
