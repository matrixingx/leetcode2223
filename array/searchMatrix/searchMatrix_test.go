package searchmatrix

import (
	"fmt"
	"testing"
)

func TestSearchmatrix(t *testing.T) {
	var matrix = [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	fmt.Println(searchMatrix(matrix,3))
	fmt.Println(searchMatrix(matrix,13))
}