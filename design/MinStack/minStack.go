package minstack

type MinStack struct {
	stack []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		v := this.GetMin()
		if v < val {
			this.minStack = append(this.minStack, v)
		} else {
			this.minStack = append(this.minStack, val)
		}
	}
}


func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}