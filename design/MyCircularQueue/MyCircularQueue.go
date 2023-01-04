package mycircularqueue

/*
	https://leetcode.cn/problems/design-circular-queue/
*/

type MyCircularQueue struct {
	Head int
	Tail int
	Size int
	MaxSize int
	Queue []int
}

func Constructor(k int) MyCircularQueue {
	myCircularQueue := MyCircularQueue{
		Queue: make([]int, k+1),
		Size: k,
		MaxSize: k+1,
	}
	return myCircularQueue
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	next := (this.Tail+1)%this.MaxSize
	this.Queue[next] = value
	this.Tail = next
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head+1)%this.MaxSize
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[(this.Head+1)%this.MaxSize]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Tail]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.Tail == this.Head
}


func (this *MyCircularQueue) IsFull() bool {
	return this.Head == (this.Tail+1)%this.MaxSize
}
