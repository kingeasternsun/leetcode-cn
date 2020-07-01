package array

type Trie struct {
	isEnd bool
	child [26]*Trie //72ms 19.6MB
}

/** Initialize your data structure here. */
func Constructor() Trie {

	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	curNode := this
	wordLen := len(word)

	for i := 0; i < wordLen; i++ {

		nextNode := curNode.child[word[i]-'a']
		if nextNode == nil {
			nextNode = &Trie{}
			curNode.child[word[i]-'a'] = nextNode
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

		nextNode := curNode.child[word[i]-'a']
		if nextNode == nil {
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

		nextNode := curNode.child[prefix[i]-'a']
		if nextNode == nil {
			return false
		}

		curNode = nextNode

	}

	return true

}
