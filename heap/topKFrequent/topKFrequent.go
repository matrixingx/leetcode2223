package topkfrequent

import "container/heap"

type numsCount struct {
	Val   int
	Count int
}

type hp []numsCount

func (this hp) Less(i, j int) bool {
	return this[i].Count > this[j].Count
}

func (this hp) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this hp) Len() int {
	return len(this)
}

func (this *hp) Push(v interface{}) {
	*this = append(*this, v.(numsCount))
}

func (this *hp) Pop() interface{} {
	h := *this
	v := h[h.Len()-1]
	*this = h[:h.Len()-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	var mp = make(map[int]int)
	var res []int
	for i := range nums {
		if v, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = 1
		} else {
			mp[nums[i]] = v + 1
		}
	}
	var h hp
	for key, value := range mp {
		h = append(h, numsCount{
			Val:   key,
			Count: value,
		})
	}
	heap.Init(&h)
	for i := 0 ; i < k ; i++ {
		v := heap.Pop(&h).(numsCount)
		res = append(res, v.Val)
	}
	return res
}