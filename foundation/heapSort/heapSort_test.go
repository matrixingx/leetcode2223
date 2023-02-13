package heapsort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testArr := []int{1,1,4,5,1,4,1,9,1,9,8,1,0}
	fmt.Println(heapSort(testArr))
}