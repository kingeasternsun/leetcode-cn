/*
 * @Description: 232
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-05 09:13:53
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-05 09:34:24
 * @FilePath: \232MyQueue\232.go
 */
package leetcode

//先实现stack
type MyStack []int

func (this *MyStack) Push(x int) {
	*this = append(*this, x)
}

func (this *MyStack) Pop() int {
	x := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return x
}
func (this MyStack) Peek() int {
	return this[len(this)-1]
}
func (this MyStack) Empty() bool {
	return len(this) == 0
}

/*
	1. 双栈，一个用来push，一个用来pop
	2. pop的时候优先从OutStatck中Pop，OutStatck为空,就把InStack中的数据导入到OutStack中，再从OutStatck中Pop
*/
type MyQueue struct {
	InStack   MyStack
	OutStatck MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

	this.InStack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.OutStatck.Empty() {
		return this.OutStatck.Pop()
	}

	for !this.InStack.Empty() {
		this.OutStatck.Push(this.InStack.Pop())
	}
	return this.OutStatck.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.OutStatck.Empty() {
		return this.OutStatck.Peek()
	}
	for !this.InStack.Empty() {
		this.OutStatck.Push(this.InStack.Pop())
	}
	return this.OutStatck.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.InStack.Empty() && this.OutStatck.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
