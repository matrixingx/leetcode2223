package lengthoflis

func lengthOfLIS(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var stack []int
	var binarySearch func(i,j int,target int,arr []int) int
	binarySearch = func(i, j int,target int, arr []int) int {
		for i <= j {
			mid := i + (j-i)/2
			if nums[arr[mid]] < target {
				i = mid+1
			} else {
				j = mid-1
			}
		}
		return i
	}
	for i := 0 ; i < n ; i ++ {
		if len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			replaceIndex := binarySearch(0,len(stack)-1,nums[i],stack)
			if replaceIndex == len(stack) {
				stack[replaceIndex-1] = i
			} else {
				stack[replaceIndex] = i
			}
		}
	}
	return len(stack)
}