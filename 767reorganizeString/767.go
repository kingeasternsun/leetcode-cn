package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
	fmt.Println(reorganizeString("vvvlo"))
	fmt.Println(reorganizeString("abbabbaaab"))
	fmt.Println(reorganizeString("ogccckcwmbmxtsbmozli")) //"cocgcickmlmsmtbwbxoz"

	//"ogccckcwmbmxtsbmozli"
}

var byteCnt [26]int

type MapHeap struct{ sort.IntSlice }

func (h MapHeap) Less(i, j int) bool { return byteCnt[h.IntSlice[i]] > byteCnt[h.IntSlice[j]] }

func (h *MapHeap) Push(v interface{}) {

	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *MapHeap) Pop() interface{} {
	// v := h.IntSlice[len(h.IntSlice)-1]
	// h.IntSlice = h.IntSlice[0 : len(h.IntSlice)-1]
	// return v

	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *MapHeap) pop() int {
	return heap.Pop(h).(int)
}
func (h *MapHeap) push(v int) { heap.Push(h, v) }

func reorganizeString(S string) string {

	cnt := len(S)
	byteCnt = [26]int{}
	maxCnt := 0
	for i := 0; i < cnt; i++ {
		byteCnt[S[i]-'a']++

		if byteCnt[S[i]-'a'] > maxCnt {
			maxCnt = byteCnt[S[i]-'a']
		}

	}

	if maxCnt > (cnt+1)/2 {
		return ""
	}

	var hp MapHeap
	var res []byte
	for i := 0; i < 26; i++ {
		if byteCnt[i] > 0 {
			hp.IntSlice = append(hp.IntSlice, i)
		}
	}

	heap.Init(&hp)

	for len(hp.IntSlice) > 1 {

		i, j := hp.pop(), hp.pop()
		res = append(res, byte('a'+i))
		res = append(res, byte('a'+j))

		if byteCnt[i]--; byteCnt[i] > 0 {
			hp.push(i)
		}

		if byteCnt[j]--; byteCnt[j] > 0 {
			hp.push(j)
		}

	}

	if len(hp.IntSlice) == 1 {

		// if byteCnt[hp.IntSlice[0]] > 1 {
		// 	return ""
		// }

		res = append(res, byte('a'+hp.IntSlice[0]))
	}

	return string(res)

}
