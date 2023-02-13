package heapsort

type Heap struct {
	arr []int
	size int
}

func adjustHeap(h Heap,parentIndex int) {
	var (
		maxNode = parentIndex
		leftNode = parentIndex*2+1
		rightNode = parentIndex*2+2
	)
	if leftNode < h.size && h.arr[leftNode] > h.arr[maxNode] {
		maxNode = leftNode
	}
	if rightNode < h.size && h.arr[rightNode] > h.arr[maxNode] {
		maxNode = rightNode
	}
	if parentIndex != maxNode {
		h.arr[parentIndex],h.arr[maxNode] = h.arr[maxNode],h.arr[parentIndex]
		adjustHeap(h,maxNode)
	}
}

func createHeap(arr []int) Heap {
	n := len(arr)
	h := Heap {
		arr: arr,
		size: n,
	}
	for i := n/2 ; i >= 0 ; i-- {
		adjustHeap(h,i)
	}
	return h
}

func heapSort(arr []int) []int {
	h := createHeap(arr)
	for h.size > 0 {
		h.arr[0],h.arr[h.size-1] = h.arr[h.size-1],h.arr[0]
		h.size--
		adjustHeap(h,0)	
	}
	return h.arr
}