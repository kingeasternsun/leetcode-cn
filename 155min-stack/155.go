package main

/*
statck 先入后出，也就是右边添加，从右边剔除
MonoArray 单调递减
*/
type MinStack struct {
	Array     []int
	MonoArray []int //
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

	this.Array = append(this.Array, x)
	if len(this.MonoArray) == 0 || x <= this.MonoArray[len(this.MonoArray)-1] {
		this.MonoArray = append(this.MonoArray, x)
	}

}

func (this *MinStack) Pop() {
	if len(this.Array) == 0 {
		return
	}

	x := this.Array[len(this.Array)-1]
	this.Array = this.Array[:len(this.Array)-1]
	if x <= this.MonoArray[len(this.MonoArray)-1] {
		this.MonoArray = this.MonoArray[:len(this.MonoArray)-1]
	}

}

func (this *MinStack) Top() int {
	return this.Array[len(this.Array)-1]
}

func (this *MinStack) GetMin() int {
	return this.MonoArray[len(this.MonoArray)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
