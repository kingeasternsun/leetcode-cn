package main

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {

	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}

	c := len(this.queue)
	v := this.queue[c-1]
	this.queue = this.queue[:c-1]
	return v

}

/** Get the top element. */
func (this *MyStack) Top() int {
	c := len(this.queue)
	v := this.queue[c-1]
	return v
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
