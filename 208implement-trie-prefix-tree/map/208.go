package map

type Trie struct {
	isEnd bool
	child map[byte]*Trie //84ms  12.8MB
}

/** Initialize your data structure here. */
func Constructor() Trie {

	return Trie{child: make(map[byte]*Trie, 0)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	curNode := this
	wordLen := len(word)

	for i := 0; i < wordLen; i++ {

		nextNode, ok := curNode.child[word[i]]
		if !ok {
			nextNode = &Trie{child: make(map[byte]*Trie, 0)}
			curNode.child[word[i]] = nextNode
		}

		curNode = nextNode

	}

	curNode.isEnd = true

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	curNode := this
	wordLen := len(word)

	for i := 0; i < wordLen; i++ {

		nextNode, ok := curNode.child[word[i]]
		if !ok {
			return false
		}

		curNode = nextNode

	}

	return curNode.isEnd

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	curNode := this
	wordLen := len(prefix)

	for i := 0; i < wordLen; i++ {

		nextNode, ok := curNode.child[prefix[i]]
		if !ok {
			return false
		}

		curNode = nextNode

	}

	return true

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
