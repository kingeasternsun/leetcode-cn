package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("word")
	param_2 := obj.Search("word")
	fmt.Println(param_2)
}

type Trie struct {
	isEnd bool
	child map[byte]*Trie
}

type WordDictionary struct {
	root *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {

	trie := &Trie{child: make(map[byte]*Trie, 0)}

	return WordDictionary{root: trie}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {

	wLen := len(word)
	cur := this.root
	for i := 0; i < wLen; i++ {

		next, ok := cur.child[word[i]]
		if ok {
			cur = next
		} else {
			next = &Trie{child: make(map[byte]*Trie, 0)}
			cur.child[word[i]] = next
		}

		cur = next
	}

	cur.isEnd = true

}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	return this.root.Search(word)
}

func (this *Trie) Search(word string) bool {

	wLen := len(word)
	cur := this
	for i := 0; i < wLen; i++ {

		next, ok := cur.child[word[i]]
		if ok {
			cur = next
			continue
		}

		if 'a' <= word[i] && word[i] <= 'z' {
			return false
		}

		//如果是 .
		for _, c := range cur.child {

			if c.Search(word[i+1:]) {
				return true
			}

		}

		return false

	}

	return cur.isEnd

}
