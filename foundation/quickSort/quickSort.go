package quicksort

func quickSort(i,j int,arr []int) {
	if i >= j {
		return
	}
	p := i
	start,end := i,j
	for start < end {
		for start < end && arr[end] >= arr[p] {
			end--
		}
		for start < end && arr[start] <= arr[p] {
			start++
		}
		arr[start],arr[end] = arr[end],arr[start]
	}
	arr[p],arr[start] = arr[start],arr[p]
	quickSort(i,start-1,arr)
	quickSort(start+1,j,arr)
}