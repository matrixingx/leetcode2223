package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testArr := []int{1,1,4,5,1,4,1,9,1,9,8,1,0}
	n := len(testArr)
	quickSort(0,n-1,testArr)
	fmt.Println(testArr)
}