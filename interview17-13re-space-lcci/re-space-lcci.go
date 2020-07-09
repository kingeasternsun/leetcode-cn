package main

import "math"

type Trie struct {
	child  map[uint8]*Trie
	MaxLen int
	isEnd  bool
}

func (t *Trie) Add(word string) {

	cur := t
	for i := 0; i < len(word); i++ {
		id := uint8(word[i] - uint8('a'))
		child, ok := cur.child[id]
		if ok {
			cur = child
		} else {
			child = &Trie{child: make(map[uint8]*Trie, 0)}
			cur.child[id] = child
			cur = child
		}
	}

	if len(word) > t.MaxLen {
		t.MaxLen = len(word)
	}

	cur.isEnd = true

}
func (t *Trie) AddReverse(word string) {

	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		id := uint8(word[i] - uint8('a'))
		child, ok := cur.child[id]
		if ok {
			cur = child
		} else {
			child = &Trie{child: make(map[uint8]*Trie, 0)}
			cur.child[id] = child
			cur = child
		}
	}

	if len(word) > t.MaxLen {
		t.MaxLen = len(word)
	}

	cur.isEnd = true

}
func (t *Trie) Find(word []byte) *Trie {

	cur := t
	for i := 0; i < len(word); i++ {
		id := uint8(word[i] - uint8('a'))
		child, ok := cur.child[id]
		if ok {
			cur = child
		} else {
			return nil
		}
	}

	return cur

}

func (t *Trie) FindReverse(word []byte) *Trie {

	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		id := uint8(word[i]) - uint8('a')
		child, ok := cur.child[id]
		if ok {
			cur = child
		} else {
			return nil
		}
	}

	return cur

}

type Solution struct {
	Trie *Trie

	Sentence []byte
	Count    int
	Cache    map[int]int
}

func (solution *Solution) RespaceHelp(beg int) int {

	res, ok := solution.Cache[beg]
	if ok {
		return res
	}

	if beg == solution.Count {
		return 0
	}

	res = math.MaxInt32
	for i := beg + 1; i <= min(beg+solution.Trie.MaxLen, solution.Count); i++ {

		tmp := i - beg
		node := solution.Trie.Find(solution.Sentence[beg:i])

		if node != nil && node.isEnd {
			tmp = 0
		}

		tmp = tmp + solution.RespaceHelp(i)
		if tmp < res {
			res = tmp
		}

		//如果 solution.Sentence[beg:i] 在trie中找不到节点，后面肯定也找不到节点
		if node == nil {
			break
		}

	}

	solution.Cache[beg] = res
	return res

}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

//534ms 28MB
func respace1(dictionary []string, sentence string) int {

	trie := &Trie{child: make(map[uint8]*Trie, 0)}
	for i := range dictionary {
		trie.Add(dictionary[i])
	}

	solution := Solution{
		Trie:     trie,
		Sentence: []byte(sentence),
		Count:    len(sentence),
		Cache:    make(map[int]int, 0),
	}

	return solution.RespaceHelp(0)
}

//172ms 33MB
func respace(dictionary []string, sentence string) int {

	trie := &Trie{child: make(map[uint8]*Trie, 0)}
	for i := range dictionary {
		trie.AddReverse(dictionary[i])
	}

	word := []byte(sentence)
	//表示 前i中的最小
	dp := make([]int, len(word)+1)
	// minRes := math.MaxInt32

	for i := 1; i <= len(word); i++ {

		res := math.MaxInt32
		cur := trie
		// [0...j...i] -> [0,j] (j,i]
		for j := i - 1; j >= 0; j-- {

			node, ok := cur.child[uint8(word[j])-uint8('a')]
			if ok == false {
				if i-j+dp[j] < res {
					res = i - j + dp[j]
				}
			} else if node.isEnd {
				if dp[j] < res {
					res = dp[j]
				}
			} else {
				if i-j+dp[j] < res {
					res = i - j + dp[j]
				}
			}

			cur = node

			if ok == false {
				break
			}

		}

		dp[i] = res

	}

	return dp[len(sentence)]

}
