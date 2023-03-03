package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	if n <= 0 {
		return []int{}
	}
	var res = make([]int,n)
	var stack []int 
	stack = append(stack, 0)
	for i := 1 ; i < n ; i ++ {
		v := stack[len(stack)-1]
		if temperatures[v] >= temperatures[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]]{
				v = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res[v] = i-v
			}
			stack = append(stack, i)
		}
	}
	return res
}