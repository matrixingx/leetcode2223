package findunsortedsubarray

import (
	"fmt"
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{1,3,2,2,2}))

	fmt.Println(findUnsortedSubarray([]int{1,2,3,3,3}))
	fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,9,15}))
	fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,15}))

}