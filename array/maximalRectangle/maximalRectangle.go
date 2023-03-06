package maximalrectangle

func maximalRectangle(matrix [][]byte) int {
	var res int
	var n = len(matrix)
	if n == 0 {
		return res
	}
	var m = len(matrix[0])
	var levelMaxLength = make([][]int,n)
	for i := range levelMaxLength {
		levelMaxLength[i] = make([]int, m)
	}
	for j := 0 ; j < m ; j++ {
		for i := 0 ; i < n ; i ++ {
			if i == 0 {
				levelMaxLength[i][j] = int(matrix[i][j]-'0')
			} else {
				if matrix[i][j] == '1' {
					levelMaxLength[i][j] = levelMaxLength[i-1][j]+1
				} else {
					levelMaxLength[i][j] = 0
				}
			}
		}
	}
	for i := range levelMaxLength {
		ans := largestRectangleArea(levelMaxLength[i])
		if ans > res {
			res = ans
		}
	}
	return res
}

func largestRectangleArea(arr []int) int {
	var n = len(arr)
	var res int
	if n == 0 {
		return 0
	}
	var leftStack,rightStack []int
	var leftMin,rightMin = make([]int,n),make([]int,n)
	for i := 0 ; i < n ; i++ {
		for len(leftStack) > 0 && arr[i] <= arr[leftStack[len(leftStack)-1]] {
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
		for len(rightStack) > 0 && arr[i] <= arr[rightStack[len(rightStack)-1]] {
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
		ans := (rightMin[i]-leftMin[i]-1)*arr[i]
		if res < ans {
			res = ans
		}
	}
	return res
}