package lockingtree

type LockingTree struct {
	parent       []int
	lockNodeUser []int
	children     [][]int
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	lockNodeUser := make([]int, n)
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		lockNodeUser[i] = -1
		p := parent[i]
		if p != -1 {
			children[p] = append(children[p], i)
		}
	}
	return LockingTree{parent, lockNodeUser, children}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.lockNodeUser[num] == -1 {
		this.lockNodeUser[num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.lockNodeUser[num] == user {
		this.lockNodeUser[num] = -1
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {

	if this.lockNodeUser[num] == -1 && !this.hasLockedAncestor(num) && this.hasLockedDescendant(num) {

		this.unlockDescendant(num, user)
		this.lockNodeUser[num] = user
		return true
	}
	return false
}

func (this *LockingTree) hasLockedAncestor(num int) bool {
	par := this.parent[num]
	for par != -1 {
		if this.lockNodeUser[par] != -1 {
			return true
		}
		par = this.parent[par]
	}
	return false
}

func (this *LockingTree) hasLockedDescendant(num int) bool {

	if this.lockNodeUser[num] != -1 {
		return true
	}
	for _, child := range this.children[num] {
		if this.hasLockedDescendant(child) {
			return true
		}
	}
	return false

}

func (this *LockingTree) unlockDescendant(num, user int) {

	this.lockNodeUser[num] = -1
	for _, child := range this.children[num] {
		this.unlockDescendant(child, user)
	}
	return

}
