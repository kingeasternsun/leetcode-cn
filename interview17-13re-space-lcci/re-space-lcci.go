package main

import "math"

func main() {

}

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

func (t *Trie) Exist(word []byte) bool {

	cur := t
	for i := 0; i < len(word); i++ {
		id := uint8(word[i] - uint8('a'))
		child, ok := cur.child[id]
		if ok {
			cur = child
		} else {
			return false
		}
	}

	return cur.isEnd

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
		// for i := beg + 1; i <= min(beg+solution.Trie.MaxLen, solution.Count); i++ { //优化 从最长的开始匹配，如果在Trie中存在 就可以直接返回了
		tmp := i - beg
		if solution.Trie.Exist(solution.Sentence[beg:i]) {
			tmp = 0
		}

		tmp = tmp + solution.RespaceHelp(i)
		if tmp < res {
			res = tmp
		}

	}

	solution.Cache[beg] = res
	return res

}

func respace(dictionary []string, sentence string) int {

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

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
