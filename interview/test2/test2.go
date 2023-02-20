package test2

import (
	"container/heap"
	"fmt"
)

type numsIndex struct {
	Val   int
	Index int
}

type hp []numsIndex

func (this hp) Less(i, j int) bool {
	return this[i].Val > this[j].Val || (this[i].Val == this[j].Val && this[i].Index < this[j].Index)
}

func (this hp) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this hp) Len() int {
	return len(this)
}

func (this *hp) Push(v interface{}) {
	*this = append(*this, v.(numsIndex))
}

func (this *hp) Pop() interface{} {
	h := *this
	v := h[h.Len()-1]
	*this = h[:h.Len()-1]
	return v
}

func main() {
	var arr = []int{3, 1, 2, 3, 3, 3, 8, 3, 3}
	var n = len(arr)
	var h hp
	for i := 0; i < n; i++ {
		h = append(h, numsIndex{
			Val:   arr[i],
			Index: i,
		})
	}
	heap.Init(&h)
	var k = 3
	for i := 0 ; i < k ; i ++ {
		v := heap.Pop(&h)
		fmt.Println(v)
	}
}
