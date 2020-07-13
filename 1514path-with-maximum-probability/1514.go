package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	Node  int
	Value float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {

	g := make([][]Edge, n)
	for i := range edges {
		g[edges[i][0]] = append(g[edges[i][0]], Edge{Node: edges[i][1], Value: -math.Log2(succProb[i])})
		g[edges[i][1]] = append(g[edges[i][1]], Edge{Node: edges[i][0], Value: -math.Log2(succProb[i])})
	}

	visted := make([]bool, n)
	maxP := make([]float64, n)
	for i := range maxP {
		maxP[i] = math.MaxFloat64
	}

	h := &Heap{}
	heap.Init(h)

	heap.Push(h, Edge{Node: start, Value: 0.0})
	for h.Len() > 0 {
		top := heap.Pop(h).(Edge)
		visted[top.Node] = true
		if top.Node == end {
			return math.Exp2(-top.Value)
		}

		for _, edge := range g[top.Node] {
			if visted[edge.Node] || top.Value+edge.Value > maxP[edge.Node] {
				continue
			}

			maxP[edge.Node] = top.Value + edge.Value
			heap.Push(h, Edge{Node: edge.Node, Value: maxP[edge.Node]})
		}

	}

	return 0
}

type Heap []Edge

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].Value < (*h)[j].Value
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *Heap) Remove(idx int) interface{} {
	h.Swap(idx, h.Len()-1)
	return h.Pop()
}
