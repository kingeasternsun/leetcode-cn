package leetcode

type Graph struct {
	Parent []int
	Height []int //以当前点为root节点的树的高度,不用height就会超时
}

func (g *Graph) Init(n int) {

	g.Parent = make([]int, n)
	g.Height = make([]int, n)
	for i := range g.Parent {
		g.Parent[i] = i
	}
	return
}

func (g *Graph) GetRoot(i int) int {
	for i != g.Parent[i] {
		i = g.Parent[i]
	}

	return i
}

func (g *Graph) Union(i, j int) bool {
	i = g.GetRoot(i)
	j = g.GetRoot(j)
	if i == j {
		return false
	}

	hi := g.Height[i]
	hj := g.Height[j]
	if hi > hj {
		g.Parent[j] = i
	} else if hi < hj {
		g.Parent[i] = j
	} else {
		g.Parent[j] = i
		g.Height[i]++
	}

	return true
}

//是否全连接
func (g *Graph) IfLinkAll() bool {
	pre := g.GetRoot(0)
	for i := 1; i < len(g.Parent); i++ {
		if g.GetRoot(i) != pre {
			return false
		}
	}
	return true
}

/*
参考
https://mp.weixin.qq.com/s/Ug324o1wPf4YoaLrH30khQ
1. 优先通过两个窦可以访问的边进行连接
*/
func maxNumEdgesToRemove(n int, edges [][]int) int {

	if n <= 1 {
		return 0
	}

	// gA, gB := Graph{}, Graph{}
	// gA.Init(n)
	// gB.Init(n)

	gs := make([]Graph, 2)
	for i := range gs {
		gs[i].Init(n)
	}

	delNum := 0
	//公共边放在前面
	// sort.Slice(edges, func(i, j int) bool {
	// 	return edges[i][0] > edges[j][0]
	// })

	for _, edge := range edges {

		if edge[0] == 3 {

			ua := gs[0].Union(edge[1]-1, edge[2]-1)
			ub := gs[1].Union(edge[1]-1, edge[2]-1)
			if ua || ub { //必须的，不能删除
				continue
			}

			delNum++
		}

	}

	for _, edge := range edges {

		if edge[0] == 3 {
			continue
		}
		u := gs[edge[0]-1].Union(edge[1]-1, edge[2]-1)
		if !u {
			delNum++
		}

	}

	for _, g := range gs {
		if !g.IfLinkAll() {
			return -1
		}
	}

	return delNum

}
