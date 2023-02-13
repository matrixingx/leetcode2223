package uthuglynumber

import "container/heap"

type myHeap []int

func (this myHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this myHeap) Len() int {
	return len(this)
}

func (this myHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this *myHeap) Push(v interface{}) {
	*this = append(*this, v.(int))
}

func (this *myHeap) Pop() interface{} {
	arr := *this
	res := arr[this.Len()-1]
	*this = arr[:this.Len()-1]
	return res
}

func nthUglyNumber(n int) int {
	var mp = make(map[int]bool)
	var h myHeap
	h = append(h, 1)
	mp[1] = true
	heap.Init(&h)
	for i := 1 ; i <= n ; i ++ {
		v := heap.Pop(&h).(int)
		if i == n {
			return v
		}
		if _,ok := mp[v*2];!ok {
			mp[v*2] = true
			heap.Push(&h,v*2)
		}
		if _,ok := mp[v*3];!ok {
			mp[v*3] = true
			heap.Push(&h,v*3)
		}
		if _,ok := mp[v*5];!ok {
			mp[v*5] = true
			heap.Push(&h,v*5)
		}
	}
	return -1
}