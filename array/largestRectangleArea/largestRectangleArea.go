package largestrectanglearea

/*
	https://leetcode.cn/problems/largest-rectangle-in-histogram/
*/

func largestRectangleArea(heights []int) int {
	var res int
	var leftStack,rightStack []int
	var n = len(heights)
	var leftMin,rightMin = make([]int,n),make([]int,n)
	for i := 0 ; i < n ; i ++ {
		for len(leftStack) > 0 && heights[i] <= heights[leftStack[len(leftStack)-1]] {
			leftStack = leftStack[:len(leftStack)-1]
		}
		if len(leftStack) == 0 {
			leftMin[i] = -1
		} else {
			leftMin[i] = leftStack[len(leftStack)-1]
		}
		leftStack = append(leftStack, i)
	}
	for i := n-1 ; i >= 0 ; i-- {
		for len(rightStack) > 0 && heights[i] <= heights[rightStack[len(rightStack)-1]] {
			rightStack = rightStack[:len(rightStack)-1]
		}
		if len(rightStack) == 0 {
			rightMin[i] = n
		} else {
			rightMin[i] = rightStack[len(rightStack)-1]
		}
		rightStack = append(rightStack, i)
	}
	for i := 0 ; i < n ; i ++ {
		ans := (rightMin[i]-leftMin[i]-1)*heights[i]
		if ans > res {
			res = ans
		}
	}
	return res
}