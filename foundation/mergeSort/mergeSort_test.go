package mergesort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testArr := []int{1,1,4,5,1,4,1,9,1,9,8,1,0}
	fmt.Println(mergeSort(testArr))
}