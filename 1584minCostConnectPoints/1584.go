package leetcode

import "sort"

type Edge struct {
	x, y   int
	Length int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func edgeLen(pi, pj []int) int {
	return abs(pi[0]-pj[0]) + abs(pi[1]-pj[1])
}

func getRoot(i int, par []int) int {

	for i != par[i] {
		i = par[i]
	}
	return i
}

func union(i, j int, par []int) bool {

	i = getRoot(i, par)
	j = getRoot(j, par)
	if i == j {
		return false
	}

	par[i] = j
	return true
}

//最小代价生成树
func minCostConnectPoints(points [][]int) int {
	edges := make([]Edge, 0)
	for i, pi := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{x: i, y: j, Length: edgeLen(pi, points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Length < edges[j].Length
	})

	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}

	res := 0
	edgeNum := 0 //加入树的边的个数
	for _, edge := range edges {

		if union(edge.x, edge.y, parent) {
			res += edge.Length
			edgeNum++
			if edgeNum == len(points)-1 {
				return res
			}
		}

	}

	return res
}
