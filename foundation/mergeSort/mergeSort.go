package mergesort

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr)/2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	res := merge(left,right)
	return res
}

func merge(left,right []int) []int {
	var n,m = len(left),len(right)
	var res []int
	var i,j = 0,0
	for i < n && j < m {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for i < n {
		res = append(res, left[i])
		i++
	}
	for j < m {
		res = append(res, right[j])
		j++
	}
	return res
}