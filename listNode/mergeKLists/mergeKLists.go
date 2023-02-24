package mergeklists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListHeap []*ListNode

func (this ListHeap) Len() int {
	return len(this)
}

func (this ListHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ListHeap) Less(i, j int) bool {
	return this[i].Val < this[j].Val
}

func (this *ListHeap) Push(v interface{}) {
	*this = append(*this, v.(*ListNode))
}

func (this *ListHeap) Pop() interface{} {
	l := *this
	v := l[l.Len()-1]
	*this = l[:l.Len()-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	var dummyNode = &ListNode{}
	var head = dummyNode
	var hp ListHeap
	for i := range lists {
		if lists[i] != nil {
			hp = append(hp, lists[i])
		}
	}
	heap.Init(&hp)
	for hp.Len() > 0 {
		node := heap.Pop(&hp).(*ListNode)
		if node.Next != nil {
			heap.Push(&hp,node.Next)
		}
		head.Next = node
		head = head.Next
	}
	return dummyNode.Next
}