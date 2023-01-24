package generatematrix

func generateMatrix(n int) [][]int {
	var res [][]int = make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	var (
		left,right,up,down = 0,n-1,0,n-1
		index = 1
	)
	for left <= right && up <= down {
		for i := left ; i <= right ; i++ {
			res[up][i] = index
			index++
		}
		for i := up+1 ; i <= down ; i++ {
			res[i][right] = index
			index++
		}
		if left < right && up < down {
			for i := right-1 ; i > left ; i-- {
				res[down][i] = index
				index++
			}
			for i := down ; i > up ; i-- {
				res[i][left] = index
				index++
			}
		}
		left++
		right--
		up++
		down--
	}
	return res
}