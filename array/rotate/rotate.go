package rotate

func rotate(matrix [][]int)  {
	var n = len(matrix)
	if n <= 0 {
		return
	}
	var m = len(matrix[0])
	for i := 0 ; i < n/2; i++ {
		for j := 0 ; j < m ; j ++ {
			matrix[i][j],matrix[n-i-1][j] = matrix[n-i-1][j],matrix[i][j]
		}
	}
	for i := 0 ; i < n ; i ++ {
		for j := 0 ; j < i ; j ++ {
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
}