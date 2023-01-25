package findkthlargest

import "container/heap"

type hp []int

func (this hp) Len() int {
	return len(this)
}

func (this hp) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this hp) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this *hp) Push(v interface{}) {
	*this = append(*this, v.(int))
}

func (this *hp) Pop() interface{} {
	h := *this
	v := h[h.Len()-1]
	*this = h[:h.Len()-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	var h hp
	for i := range nums {
		h = append(h, nums[i])
	}
	heap.Init(&h)
	var res int
	for i := 0 ; i < k ; i++ {
		res = heap.Pop(&h).(int)
	}
	return res
}

func findKthLargest2(nums []int,k int) int {
	var(
		quickSort func (i,j int)
		find bool
		res int
		n = len(nums)
	)
	quickSort = func(i, j int) {
		if i > j || find {
			return
		}
		point := nums[i]
		start,end := i,j
		for start < end {
			for start < end && nums[end] >= point {
				end --
			}
			for start < end && nums[start] <= point {
				start ++
			}
			nums[start],nums[end] = nums[end],nums[start]
		}
		nums[i],nums[start] = nums[start],nums[i]
		if start == n-k {
			res = nums[start]
			find = true
			return
		} else if start < n-k {
			quickSort(start+1,j)
		} else {
			quickSort(i,start-1)
		}
	}
	quickSort(0,n-1)
	return res
}