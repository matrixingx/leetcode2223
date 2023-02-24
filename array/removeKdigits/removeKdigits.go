package removekdigits

import "strings"

func removeKdigits(num string, k int) string {
	var n = len(num)
	if n <= k {
		return "0"
	}
	var stack []byte
	var removeCount = 0
	for i := 0; i < n; i++ {
		dig := num[i]
		for len(stack) > 0 && removeCount < k && dig < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			removeCount++
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-(k-removeCount)]
	res := strings.TrimLeft(string(stack),"0")
	if res == "" {
		return "0"
	}
	return res
}