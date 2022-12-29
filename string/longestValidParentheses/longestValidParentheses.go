package longestvalidparentheses

/*
	https://leetcode.cn/problems/longest-valid-parentheses/
*/

func longestValidParentheses(s string) int {
	var res int
	var stack []int
	stack = append(stack, -1)
	for i,v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				res = max(res,i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}
	return res
}

func max(arr ...int) int {
	var min = -1 << 32
	for i := range arr {
		if min < arr[i] {
			min = arr[i]
		}
	}
	return min
}