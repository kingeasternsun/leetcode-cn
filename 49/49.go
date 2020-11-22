package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(group([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

type SortBytes struct {
	s  []byte
	ID int
}

func group(in []string) [][]string {

	if len(in) <= 1 {
		return [][]string{in}
	}

	sortString := make([]SortBytes, 0)

	for i := range in {

		sb := []byte(in[i])
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})

		sortString = append(sortString, SortBytes{s: sb, ID: i})

	}

	sort.Slice(sortString, func(i, j int) bool {
		return string(sortString[i].s) < string(sortString[j].s)
	})

	res := make([][]string, 0)
	// res = append(res, []string{sortString[0]})
	tmp := []string{in[sortString[0].ID]}
	for i := 1; i < len(sortString); i++ {

		if string(sortString[i].s) != string(sortString[i-1].s) {
			res = append(res, tmp)
			tmp = []string{in[sortString[i].ID]}
		} else {

			tmp = append(tmp, in[sortString[i].ID])

		}

	}
	res = append(res, tmp)

	return res

}

func groupAnagrams1(strs []string) [][]string {

	if len(strs) <= 1 {
		return [][]string{strs}
	}

	// sortString := make([]SortBytes, 0)

	sMap := make(map[string][]string, 0)

	for i := range strs {

		sb := []byte(strs[i])
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})

		sMap[string(sb)] = append(sMap[string(sb)], strs[i])

	}

	res := make([][]string, 0)
	for _, v := range sMap {
		res = append(res, v)
	}

	return res

}

//上面的方法中，先利用map进行分组，再返回结果，可以进一步优化，利用map记录归属第res 几个组
func groupAnagrams(strs []string) [][]string {

	if len(strs) <= 1 {
		return [][]string{strs}
	}

	// sortString := make([]SortBytes, 0)

	res := make([][]string, 0)

	sMap := make(map[string]int, 0)

	for i := range strs {

		sb := []byte(strs[i])
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})

		// sMap[(sb)] = append(sMap[(sb)], strs[i])

		idx, ok := sMap[string(sb)]
		if !ok {
			idx = len(res)
			res = append(res, []string{strs[i]})
			sMap[string(sb)] = idx // 记录排序后结果再res中的位置

		} else {
			res[idx] = append(res[idx], strs[i])

		}

	}

	return res

}
