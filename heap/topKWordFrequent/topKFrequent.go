package topkwordfrequent

import "container/heap"

type wordCount struct {
	Word  string
	Count int
}

type wordHeap []wordCount

func (this wordHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this wordHeap) Less(i, j int) bool {
	return this[i].Count > this[j].Count || (this[i].Count == this[j].Count && this[i].Word < this[j].Word)
}

func (this wordHeap) Len() int {
	return len(this)
}

func (this *wordHeap) Push(v interface{}) {
	*this = append(*this, v.(wordCount))
}

func (this *wordHeap) Pop() interface{} {
	h := *this
	v := h[h.Len()-1]
	*this = h[:h.Len()-1]
	return v
}

func topKFrequent(words []string, k int) []string {
	var mp = make(map[string]int)
	for i := range words {
		mp[words[i]]++
	}
	var hp wordHeap
	for k, v := range mp {
		hp = append(hp, wordCount{
			Word:  k,
			Count: v,
		})
	}
	heap.Init(&hp)
	var res []string
	for i := 0 ; i < k ; i++ {
		v := heap.Pop(&hp).(wordCount)
		res = append(res, v.Word)
	}
	return res
}