package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {

	a := [][]int{

		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	fmt.Println(reconstructQueue(a))
}

//TODO
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	l := list.New()

}

func reconstructQueue_timeout(people [][]int) [][]int {

	cnt := len(people)
	if cnt == 0 {
		return nil
	}

	res := make([][]int, 0)
	hMap := make(map[int]int, 0)
	for _, v := range people {
		hMap[v[0]] = hMap[v[0]] + 1
	}

	help(people, hMap, &res)

	i := 0
	// j:=len(res)-1
	for i < cnt/2 {
		res[i], res[cnt-1-i] = res[cnt-1-i], res[i]
		i++
	}

	return res
}

func help(people [][]int, hMap map[int]int, res *[][]int) bool {

	if len(people) == 0 && len(hMap) == 0 {
		return true
	}

	for i, v := range people {

		//把第i个选出来,放到最后
		people[0], people[i] = people[i], people[0]
		hMap[v[0]] = hMap[v[0]] - 1
		if hMap[v[0]] == 0 {
			delete(hMap, v[0])
		}

		//如果满足条件，把剩下的再进行重排
		if match(v, hMap) {
			*res = append(*res, v)
			if help(people[1:], hMap, res) {
				return true
			}
			*res = (*res)[0 : len(*res)-1]
		}

		//不满足的话 ，再放回去
		people[0], people[i] = people[i], people[0]
		hMap[v[0]] = hMap[v[0]] + 1

	}

	return false
}

func match(v []int, hMap map[int]int) bool {

	preCnt := 0
	for k, cnt := range hMap {
		if k >= v[0] {
			preCnt += cnt
		}

		if preCnt > v[1] {
			return false
		}

	}

	return preCnt == v[1]

}
