package trap

/*
	https://leetcode.cn/problems/trapping-rain-water/
*/

func trap(height []int) int {
	var n = len(height)
	var maxLeft = make([]int,n)
	var maxRight = make([]int,n)
	if n == 0 || n == 1 {
		return 0
	}
	maxLeft[0] = height[0]
	maxRight[n-1] = height[n-1]
	for i := 1;i <n ; i++ {
		maxLeft[i] = max(maxLeft[i-1],height[i]) // 假设右边有一个很高的板子，记录左边最高的水位线
	}
	for j := n-2 ; j >= 0 ; j-- {
		maxRight[j] = max(maxRight[j+1],height[j])// 假设左边有一个很高的板子，记录右边最高的水位线
	}
	var res = make([]int,n)
	for i := 0 ; i < n ; i ++ {
		// 实际上两边都妹有板子，所以取最低值
		if maxLeft[i] < maxRight[i] {
			res[i] = maxLeft[i] - height[i]
		} else {
			res[i] = maxRight[i] - height[i]
		}
	}
	var ans = 0
	for i := range res {
		ans += res[i]
	}
	return ans
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i,j int) int {
	if i < j {
		return i
	}
	return j
}