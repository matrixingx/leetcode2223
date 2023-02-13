package binarysearch

func binarysearch2(arr []int,target int) int {
	start,end := 0,len(arr)
	for start < end {
		mid := start + (end - start)/2
		if arr[mid] < target {
			start = mid+1
		} else {
			end = mid
		}
	}
	return start
}