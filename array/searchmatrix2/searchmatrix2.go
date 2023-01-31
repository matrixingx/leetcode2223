package searchmatrix2

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var n,m = len(matrix),len(matrix[0])
	var startI,startJ = 0,m-1
	for startI < n && startJ >= 0 {
		if matrix[startI][startJ] == target {
			return true
		}
		if matrix[startI][startJ] < target {
			startI++
		} else {
			startJ--
		}
	}
	return false
}