package nextpermutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var arr = []int{1, 2, 3, 8, 5, 7, 6, 4}
	nextPermutation(arr)
	fmt.Println(arr)
	var arr2 = []int{1, 1, 5}
	nextPermutation(arr2)
	fmt.Println(arr2)
	var arr3 = []int{1, 2, 3}
	nextPermutation(arr3)
	fmt.Println(arr3)
	var arr4 = []int{3,2,1}
	nextPermutation(arr4)
	fmt.Println(arr4)
}
